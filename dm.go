// 本方言包基于gorm v1.24.2开发，需要配合达梦数据库驱动使用
package dameng

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"
	"strings"

	_ "github.com/zjy2414/dm8-gorm/dm8" // 引入dm数据库驱动包
	"gorm.io/gorm"                      // 引入gorm v2包
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/migrator"
	"gorm.io/gorm/schema"
)

// DriverName 数据库驱动、连接字符串协议名称
const DriverName = "dm"

type Config struct {
	DriverName        string
	DSN               string
	Conn              gorm.ConnPool
	DefaultStringSize uint
	GormMode          int //dm兼容 = 0; mysql兼容 = 1
}

type Dialector struct {
	*Config
}

var (
	// CreateClauses create clauses
	CreateClauses = []string{"INSERT", "VALUES", "ON CONFLICT"}
	// QueryClauses query clauses
	QueryClauses = []string{}
	// UpdateClauses update clauses
	UpdateClauses = []string{"UPDATE", "SET", "WHERE", "LIMIT"} // 不支持 "ORDER BY"
	// DeleteClauses delete clauses
	DeleteClauses = []string{"DELETE", "FROM", "WHERE", "LIMIT"} // 不支持 "ORDER BY"

	defaultDatetimePrecision = 3
)

func Open(dsn string) gorm.Dialector {
	return &Dialector{Config: &Config{DSN: dsn}}
}

func New(config Config) gorm.Dialector {
	return &Dialector{Config: &config}
}

// BuildUrl 构建达梦数据库连接字符串
//
//   - 如： dm://user:password@host:port?schema=SYSDBA[&propName2=propValue2]…
//   - 若要指定用户登录后的当前模式，请在 options 中设置 schema，缺省为用户的默认模式，如 SYSDBA
//   - 密码等属性值不做 URL 转义（达梦驱动按原始字符解析连接串，特殊字符无需转义）
//   - 参考链接： https://eco.dameng.com/document/dm/zh-cn/pm/go-rogramming-guide.html#11.3%20%E8%BF%9E%E6%8E%A5%E4%B8%B2%E5%B1%9E%E6%80%A7%E8%AF%B4%E6%98%8E
func BuildUrl(user, password, host string, port int, urlOptions ...map[string]string) string {
	var builder strings.Builder
	builder.WriteString("dm://")
	builder.WriteString(user)
	builder.WriteString(":")
	builder.WriteString(password)
	builder.WriteString("@")
	builder.WriteString(net.JoinHostPort(host, strconv.Itoa(port)))

	var options []string
	for _, o := range urlOptions {
		for key, value := range o {
			if value == "" {
				options = append(options, key)
			} else {
				options = append(options, key+"="+value)
			}
		}
	}
	if len(options) > 0 {
		builder.WriteString("?")
		builder.WriteString(strings.Join(options, "&"))
	}
	return builder.String()
}

func (d Dialector) Name() string {
	return DriverName
}

func (d Dialector) Initialize(db *gorm.DB) (err error) {
	if d.DriverName == "" {
		d.DriverName = DriverName
	}

	if d.Conn != nil {
		db.ConnPool = d.Conn
	} else {
		db.ConnPool, err = sql.Open(d.DriverName, d.DSN)
		if err != nil {
			return
		}
	}

	// register callbacks
	callbackConfig := &callbacks.Config{
		CreateClauses: CreateClauses,
		QueryClauses:  QueryClauses,
		UpdateClauses: UpdateClauses,
		DeleteClauses: DeleteClauses,
	}
	callbacks.RegisterDefaultCallbacks(db, callbackConfig)
	db.Callback().Create().Replace("gorm:create", Create)

	return
}

func (d Dialector) DefaultValueOf(*schema.Field) clause.Expression {
	// return clause.Expr{SQL: "DEFAULT VALUES"}
	// 和gorm v1不一样，gorm v1是 INSERT INTO XXXX DEFAULT VALUES;
	// gorm v2是 INSERT INTO XXXX(C1, C2) VALUES(XX, NULL);
	return clause.Expr{SQL: "NULL"}
}

func (d Dialector) Migrator(db *gorm.DB) gorm.Migrator {
	return Migrator{
		Migrator: migrator.Migrator{
			Config: migrator.Config{
				DB:                          db,
				Dialector:                   d,
				CreateIndexAfterCreateTable: true,
			},
		},
		Dialector: d,
	}
}

func (d Dialector) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	writer.WriteByte('?')
}

func (d Dialector) QuoteTo(writer clause.Writer, str string) {
	var (
		underQuoted, selfQuoted bool
		continuousBacktick      int8
		shiftDelimiter          int8
	)

	for _, v := range []byte(str) {
		switch v {
		case '"':
			continuousBacktick++
			if continuousBacktick == 2 {
				writer.WriteString(`""`)
				continuousBacktick = 0
			}
		case '.':
			if continuousBacktick > 0 || !selfQuoted {
				shiftDelimiter = 0
				underQuoted = false
				continuousBacktick = 0
				writer.WriteByte('"')
			}
			writer.WriteByte(v)
			continue
		default:
			if shiftDelimiter-continuousBacktick <= 0 && !underQuoted {
				writer.WriteByte('"')
				underQuoted = true
				if selfQuoted = continuousBacktick > 0; selfQuoted {
					continuousBacktick -= 1
				}
			}

			for ; continuousBacktick > 0; continuousBacktick -= 1 {
				writer.WriteString(`""`)
			}

			writer.WriteByte(v)
		}
		shiftDelimiter++
	}

	if continuousBacktick > 0 && !selfQuoted {
		writer.WriteString(`""`)
	}
	writer.WriteByte('"')
}

func (d Dialector) Explain(sql string, vars ...interface{}) string {
	return logger.ExplainSQL(sql, nil, `'`, vars...)
}

func (d Dialector) DataTypeOf(field *schema.Field) string {
	switch field.DataType {
	case schema.Bool:
		return "BIT"
	case schema.Int, schema.Uint:
		return d.getSchemaIntAndUnitType(field)
	case schema.Float:
		return d.getSchemaFloatType(field)
	case schema.String:
		return d.getSchemaStringType(field)
	case schema.Time:
		return d.getSchemaTimeType(field)
	case schema.Bytes:
		return d.getSchemaBytesType(field)
	default:
		return d.getSchemaCustomType(field)
		// what oracle do:
		//notNull, _ := field.TagSettings["NOT NULL"]
		//unique, _ := field.TagSettings["UNIQUE"]
		//additionalType := fmt.Sprintf("%s %s", notNull, unique)
		//if value, ok := field.TagSettings["DEFAULT"]; ok {
		//	additionalType = fmt.Sprintf("%s %s %s%s", "DEFAULT", value, additionalType, func() string {
		//		if value, ok := field.TagSettings["COMMENT"]; ok {
		//			return " COMMENT " + value
		//		}
		//		return ""
		//	}())
		//}
		//sqlType = fmt.Sprintf("%v %v", sqlType, additionalType)
	}
}

func (d Dialector) getSchemaIntAndUnitType(field *schema.Field) string {
	constraint := func(sqlType string) string {
		//if field.NotNull {
		//	sqlType += " NOT NULL"
		//}
		if field.AutoIncrement {
			sqlType += " IDENTITY(1,1)"
		}
		return sqlType
	}

	switch {
	case field.Size <= 8:
		return constraint("TINYINT")
	case field.Size <= 16:
		return constraint("SMALLINT")
	case field.Size <= 32:
		return constraint("INT")
	default:
		return constraint("BIGINT")
	}
}

func (d Dialector) getSchemaFloatType(field *schema.Field) string {
	if field.Precision > 0 {
		return fmt.Sprintf("DECIMAL(%d, %d)", field.Precision, field.Scale)
	}

	return "DOUBLE"
}

func (d Dialector) getSchemaStringType(field *schema.Field) string {
	size := field.Size

	if size == 0 {
		if d.DefaultStringSize > 0 {
			size = int(d.DefaultStringSize)
		} else {
			hasIndex := field.TagSettings["INDEX"] != "" || field.TagSettings["UNIQUE"] != ""
			// TEXT, GEOMETRY or JSON column can't have a default value
			if field.PrimaryKey || field.HasDefaultValue || hasIndex {
				size = 255 // mysql:191, dm not support utf8mb4
			}
		}
	}

	if size <= 0 || size >= 32768 {
		return "CLOB"
	} else {
		return fmt.Sprintf("VARCHAR(%d)", size)
	}
}

func (d Dialector) getSchemaTimeType(field *schema.Field) string {
	if field.Precision > 0 {
		return fmt.Sprintf("TIMESTAMP(%d)", field.Precision)
	}
	return "TIMESTAMP"
	//if field.NotNull || field.PrimaryKey {
	//	sqlType += " NOT NULL"
	//}
}

func (d Dialector) getSchemaBytesType(field *schema.Field) string {
	if field.Size > 0 && field.Size < 32768 {
		return fmt.Sprintf("VARBINARY(%d)", field.Size)
	}

	return "BLOB"
}

func (d Dialector) getSchemaCustomType(field *schema.Field) string {
	sqlType := string(field.DataType)

	if field.AutoIncrement && !strings.Contains(strings.ToLower(sqlType), " auto_increment") && !strings.Contains(strings.ToLower(sqlType), " identity") {
		sqlType += " IDENTITY(1,1)"
	}

	return sqlType
}

func (d Dialector) SavePoint(tx *gorm.DB, name string) error {
	return tx.Exec("SAVEPOINT " + name).Error
}

func (d Dialector) RollbackTo(tx *gorm.DB, name string) error {
	return tx.Exec("ROLLBACK TO SAVEPOINT " + name).Error
}
