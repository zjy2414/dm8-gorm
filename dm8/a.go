/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm8

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/zjy2414/dm8-gorm/dm8/security"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	Dm_build_0 = 8192
	Dm_build_1 = 2 * time.Second
)

type dm_build_2 struct {
	dm_build_3 Dm_build_268
	dm_build_4 *Dm_build_1357
	dm_build_5 *DmConnection
	dm_build_6 security.Cipher
	dm_build_7 bool
	dm_build_8 bool
	dm_build_9 *security.DhKey

	dm_build_10 bool
	dm_build_11 string
	dm_build_12 bool
}

func dm_build_13(dm_build_14 context.Context, dm_build_15 *DmConnection) (*dm_build_2, error) {

	dm_build_16 := &Dm_build_270{}
	if err := dm_build_16.Dm_build_269(dm_build_14, dm_build_15); err != nil {
		return nil, err
	}

	dm_build_17 := &dm_build_2{}
	dm_build_17.dm_build_3 = dm_build_16
	dm_build_17.dm_build_4 = Dm_build_1360(Dm_build_352)
	dm_build_17.dm_build_5 = dm_build_15
	dm_build_17.dm_build_7 = false
	dm_build_17.dm_build_8 = false
	dm_build_17.dm_build_10 = false
	dm_build_17.dm_build_11 = ""
	dm_build_17.dm_build_12 = false
	dm_build_15.Access = dm_build_17

	return dm_build_17, nil
}

func (dm_build_19 *dm_build_2) dm_build_18(dm_build_20 dm_build_474) bool {
	var dm_build_21 = dm_build_19.dm_build_5.dmConnector.compress
	if dm_build_20.dm_build_489() == Dm_build_380 || dm_build_21 == Dm_build_429 {
		return false
	}

	if dm_build_21 == Dm_build_427 {
		return true
	} else if dm_build_21 == Dm_build_428 {
		return !dm_build_19.dm_build_5.Local && dm_build_20.dm_build_487() > Dm_build_426
	}

	return false
}

func (dm_build_23 *dm_build_2) dm_build_22(dm_build_24 dm_build_474) bool {
	var dm_build_25 = dm_build_23.dm_build_5.dmConnector.compress
	if dm_build_24.dm_build_489() == Dm_build_380 || dm_build_25 == Dm_build_429 {
		return false
	}

	if dm_build_25 == Dm_build_427 {
		return true
	} else if dm_build_25 == Dm_build_428 {
		return dm_build_23.dm_build_4.Dm_build_1624(Dm_build_388) == 1
	}

	return false
}

func (dm_build_27 *dm_build_2) dm_build_26(dm_build_28 dm_build_474) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_30 := dm_build_28.dm_build_487()

	if dm_build_30 > 0 {

		if dm_build_27.dm_build_18(dm_build_28) {
			var retBytes, err = Compress(dm_build_27.dm_build_4, Dm_build_381, int(dm_build_30), int(dm_build_27.dm_build_5.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_27.dm_build_4.Dm_build_1371(Dm_build_381)

			dm_build_27.dm_build_4.Dm_build_1412(dm_build_30)

			dm_build_27.dm_build_4.Dm_build_1440(retBytes)

			dm_build_28.dm_build_488(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_27.dm_build_4.Dm_build_1544(Dm_build_388, 1)
		}

		if dm_build_27.dm_build_8 {
			dm_build_30 = dm_build_28.dm_build_487()
			var retBytes = dm_build_27.dm_build_6.Encrypt(dm_build_27.dm_build_4.Dm_build_1651(Dm_build_381, int(dm_build_30)), true)

			dm_build_27.dm_build_4.Dm_build_1371(Dm_build_381)

			dm_build_27.dm_build_4.Dm_build_1440(retBytes)

			dm_build_28.dm_build_488(int32(len(retBytes)))
		}
	}

	if dm_build_27.dm_build_4.Dm_build_1369() > Dm_build_353 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_28.dm_build_483()

	dm_build_27.dm_build_4.Dm_build_1374(0)
	if _, err := dm_build_27.dm_build_4.Dm_build_1393(dm_build_27.dm_build_3); err != nil {
		return err
	}

	return nil
}

func (dm_build_32 *dm_build_2) dm_build_31(dm_build_33 dm_build_474) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_35 := int32(0)

	dm_build_32.dm_build_4.Dm_build_1371(0)

	if _, err := dm_build_32.dm_build_4.Dm_build_1387(dm_build_32.dm_build_3, Dm_build_381); err != nil {
		return err
	}
	dm_build_35 = dm_build_33.dm_build_487()

	if dm_build_35 > 0 {
		if _, err := dm_build_32.dm_build_4.Dm_build_1387(dm_build_32.dm_build_3, int(dm_build_35)); err != nil {
			return err
		}
	}

	dm_build_33.dm_build_484()
	dm_build_35 = dm_build_33.dm_build_487()
	if dm_build_35 <= 0 {
		return nil
	}

	if dm_build_32.dm_build_8 {
		ebytes := dm_build_32.dm_build_4.Dm_build_1651(Dm_build_381, int(dm_build_35))
		bytes, err := dm_build_32.dm_build_6.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_32.dm_build_4.Dm_build_1371(Dm_build_381)
		dm_build_32.dm_build_4.Dm_build_1440(bytes)
		dm_build_33.dm_build_488(int32(len(bytes)))
	}

	if dm_build_32.dm_build_22(dm_build_33) {

		dm_build_35 = dm_build_33.dm_build_487()
		cbytes := dm_build_32.dm_build_4.Dm_build_1651(Dm_build_381+ULINT_SIZE, int(dm_build_35-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_32.dm_build_5.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_32.dm_build_4.Dm_build_1371(Dm_build_381)
		dm_build_32.dm_build_4.Dm_build_1440(bytes)
		dm_build_33.dm_build_488(int32(len(bytes)))
	}
	return nil
}

func (dm_build_37 *dm_build_2) dm_build_36(dm_build_38 dm_build_474) (dm_build_39 interface{}, dm_build_40 error) {
	if dm_build_37.dm_build_12 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_41 := dm_build_37.dm_build_5
	dm_build_41.mu.Lock()
	defer dm_build_41.mu.Unlock()
	dm_build_40 = dm_build_38.dm_build_478(dm_build_38)
	if dm_build_40 != nil {
		return nil, dm_build_40
	}

	dm_build_40 = dm_build_37.dm_build_26(dm_build_38)
	if dm_build_40 != nil {
		return nil, dm_build_40
	}

	dm_build_40 = dm_build_37.dm_build_31(dm_build_38)
	if dm_build_40 != nil {
		return nil, dm_build_40
	}

	return dm_build_38.dm_build_482(dm_build_38)
}

func (dm_build_43 *dm_build_2) dm_build_42() (*dm_build_936, error) {

	Dm_build_44 := dm_build_942(dm_build_43)
	_, dm_build_45 := dm_build_43.dm_build_36(Dm_build_44)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	return Dm_build_44, nil
}

func (dm_build_47 *dm_build_2) dm_build_46() error {

	dm_build_48 := dm_build_800(dm_build_47)
	_, dm_build_49 := dm_build_47.dm_build_36(dm_build_48)
	if dm_build_49 != nil {
		return dm_build_49
	}

	return nil
}

func (dm_build_51 *dm_build_2) dm_build_50() error {

	var dm_build_52 *dm_build_936
	var err error
	if dm_build_52, err = dm_build_51.dm_build_42(); err != nil {
		return err
	}

	if err = dm_build_51.dm_build_265(); err != nil {
		return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
	}

	if dm_build_51.dm_build_8 || dm_build_51.dm_build_7 {
		k, err := dm_build_51.dm_build_255()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_52.Dm_build_940)
		encryptType := dm_build_52.dm_build_938
		hashType := int(dm_build_52.Dm_build_939)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_51.dm_build_258(encryptType, sessionKey, dm_build_51.dm_build_5.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_51.dm_build_46(); err != nil {
		return err
	}
	return nil
}

func (dm_build_55 *dm_build_2) Dm_build_54(dm_build_56 *DmStatement) error {
	dm_build_57 := dm_build_969(dm_build_55, dm_build_56)
	_, dm_build_58 := dm_build_55.dm_build_36(dm_build_57)
	if dm_build_58 != nil {
		return dm_build_58
	}

	return nil
}

func (dm_build_60 *dm_build_2) Dm_build_59(dm_build_61 int32) error {
	dm_build_62 := dm_build_979(dm_build_60, dm_build_61)
	_, dm_build_63 := dm_build_60.dm_build_36(dm_build_62)
	if dm_build_63 != nil {
		return dm_build_63
	}

	return nil
}

func (dm_build_65 *dm_build_2) Dm_build_64(dm_build_66 *DmStatement, dm_build_67 bool, dm_build_68 int16) (*execRetInfo, error) {
	dm_build_69 := dm_build_842(dm_build_65, dm_build_66, dm_build_67, dm_build_68)
	dm_build_70, dm_build_71 := dm_build_65.dm_build_36(dm_build_69)
	if dm_build_71 != nil {
		return nil, dm_build_71
	}
	return dm_build_70.(*execRetInfo), nil
}

func (dm_build_73 *dm_build_2) Dm_build_72(dm_build_74 *DmStatement, dm_build_75 int16) (*execRetInfo, error) {
	return dm_build_73.Dm_build_64(dm_build_74, false, Dm_build_433)
}

func (dm_build_77 *dm_build_2) Dm_build_76(dm_build_78 *DmStatement, dm_build_79 []OptParameter) (*execRetInfo, error) {
	dm_build_80, dm_build_81 := dm_build_77.dm_build_36(dm_build_577(dm_build_77, dm_build_78, dm_build_79))
	if dm_build_81 != nil {
		return nil, dm_build_81
	}

	return dm_build_80.(*execRetInfo), nil
}

func (dm_build_83 *dm_build_2) Dm_build_82(dm_build_84 *DmStatement, dm_build_85 int16) (*execRetInfo, error) {
	return dm_build_83.Dm_build_64(dm_build_84, true, dm_build_85)
}

func (dm_build_87 *dm_build_2) Dm_build_86(dm_build_88 *DmStatement, dm_build_89 [][]interface{}) (*execRetInfo, error) {
	dm_build_90 := dm_build_609(dm_build_87, dm_build_88, dm_build_89)
	dm_build_91, dm_build_92 := dm_build_87.dm_build_36(dm_build_90)
	if dm_build_92 != nil {
		return nil, dm_build_92
	}
	return dm_build_91.(*execRetInfo), nil
}

func (dm_build_94 *dm_build_2) Dm_build_93(dm_build_95 *DmStatement, dm_build_96 [][]interface{}, dm_build_97 bool) (*execRetInfo, error) {
	var dm_build_98, dm_build_99 = 0, 0
	var dm_build_100 = len(dm_build_96)
	var dm_build_101 [][]interface{}
	var dm_build_102 = NewExceInfo()
	dm_build_102.updateCounts = make([]int64, dm_build_100)
	var dm_build_103 = false
	for dm_build_98 < dm_build_100 {
		for dm_build_99 = dm_build_98; dm_build_99 < dm_build_100; dm_build_99++ {
			paramData := dm_build_96[dm_build_99]
			bindData := make([]interface{}, dm_build_95.paramCount)
			dm_build_103 = false
			for icol := 0; icol < int(dm_build_95.paramCount); icol++ {
				if dm_build_95.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_94.dm_build_238(bindData, paramData, icol) {
					dm_build_103 = true
					break
				}
			}

			if dm_build_103 {
				break
			}
			dm_build_101 = append(dm_build_101, bindData)
		}

		if dm_build_99 != dm_build_98 {
			tmpExecInfo, err := dm_build_94.Dm_build_86(dm_build_95, dm_build_101)
			if err != nil {
				return nil, err
			}
			dm_build_101 = dm_build_101[0:0]
			dm_build_102.union(tmpExecInfo, dm_build_98, dm_build_99-dm_build_98)
		}

		if dm_build_99 < dm_build_100 {
			tmpExecInfo, err := dm_build_94.Dm_build_112(dm_build_95, dm_build_96[dm_build_99], dm_build_97)
			if err != nil {
				return nil, err
			}

			dm_build_97 = true
			dm_build_102.union(tmpExecInfo, dm_build_99, 1)
		}

		dm_build_98 = dm_build_99 + 1
	}
	for _, i := range dm_build_102.updateCounts {
		if i > 0 {
			dm_build_102.updateCount += i
		}
	}
	return dm_build_102, nil
}

func (dm_build_105 *dm_build_2) dm_build_104(dm_build_106 *DmStatement, dm_build_107 []parameter) error {
	if !dm_build_106.prepared {
		retInfo, err := dm_build_105.Dm_build_64(dm_build_106, false, Dm_build_433)
		if err != nil {
			return nil
		}
		dm_build_106.serverParams = retInfo.serverParams
		dm_build_106.paramCount = int32(len(dm_build_106.serverParams))
		dm_build_106.prepared = true
	}

	dm_build_108 := dm_build_831(dm_build_105, dm_build_106, dm_build_106.bindParams)
	dm_build_109, err := dm_build_105.dm_build_36(dm_build_108)
	if err != nil {
		return nil
	}
	retInfo := dm_build_109.(*execRetInfo)
	if retInfo.serverParams != nil && len(retInfo.serverParams) > 0 {
		dm_build_106.serverParams = retInfo.serverParams
		dm_build_106.paramCount = int32(len(dm_build_106.serverParams))
	}
	dm_build_106.preExec = true
	return nil
}

func (dm_build_113 *dm_build_2) Dm_build_112(dm_build_114 *DmStatement, dm_build_115 []interface{}, dm_build_116 bool) (*execRetInfo, error) {

	var dm_build_117 = make([]interface{}, dm_build_114.paramCount)
	for icol := 0; icol < int(dm_build_114.paramCount); icol++ {
		if dm_build_114.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_113.dm_build_238(dm_build_117, dm_build_115, icol) {

			if !dm_build_116 {
				dm_build_113.dm_build_104(dm_build_114, dm_build_114.bindParams)

				dm_build_116 = true
			}

			dm_build_113.dm_build_244(dm_build_114, dm_build_114.bindParams[icol], icol, dm_build_115[icol].(iOffRowBinder))
			dm_build_117[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_118 = make([][]interface{}, 1, 1)
	dm_build_118[0] = dm_build_117

	dm_build_119 := dm_build_609(dm_build_113, dm_build_114, dm_build_118)
	dm_build_120, dm_build_121 := dm_build_113.dm_build_36(dm_build_119)
	if dm_build_121 != nil {
		return nil, dm_build_121
	}
	return dm_build_120.(*execRetInfo), nil
}

func (dm_build_123 *dm_build_2) Dm_build_122(dm_build_124 *DmStatement, dm_build_125 int16) (*execRetInfo, error) {
	dm_build_126 := dm_build_818(dm_build_123, dm_build_124, dm_build_125)

	dm_build_127, dm_build_128 := dm_build_123.dm_build_36(dm_build_126)
	if dm_build_128 != nil {
		return nil, dm_build_128
	}
	return dm_build_127.(*execRetInfo), nil
}

func (dm_build_130 *dm_build_2) Dm_build_129(dm_build_131 *innerRows, dm_build_132 int64) (*execRetInfo, error) {
	dm_build_133 := dm_build_717(dm_build_130, dm_build_131, dm_build_132, INT64_MAX)
	dm_build_134, dm_build_135 := dm_build_130.dm_build_36(dm_build_133)
	if dm_build_135 != nil {
		return nil, dm_build_135
	}
	return dm_build_134.(*execRetInfo), nil
}

func (dm_build_137 *dm_build_2) Commit() error {
	dm_build_138 := dm_build_562(dm_build_137)
	_, dm_build_139 := dm_build_137.dm_build_36(dm_build_138)
	if dm_build_139 != nil {
		return dm_build_139
	}

	return nil
}

func (dm_build_141 *dm_build_2) Rollback() error {
	dm_build_142 := dm_build_880(dm_build_141)
	_, dm_build_143 := dm_build_141.dm_build_36(dm_build_142)
	if dm_build_143 != nil {
		return dm_build_143
	}

	return nil
}

func (dm_build_145 *dm_build_2) Dm_build_144(dm_build_146 *DmConnection) error {
	dm_build_147 := dm_build_885(dm_build_145, dm_build_146.IsoLevel)
	_, dm_build_148 := dm_build_145.dm_build_36(dm_build_147)
	if dm_build_148 != nil {
		return dm_build_148
	}

	return nil
}

func (dm_build_150 *dm_build_2) Dm_build_149(dm_build_151 *DmStatement, dm_build_152 string) error {
	dm_build_153 := dm_build_567(dm_build_150, dm_build_151, dm_build_152)
	_, dm_build_154 := dm_build_150.dm_build_36(dm_build_153)
	if dm_build_154 != nil {
		return dm_build_154
	}

	return nil
}

func (dm_build_156 *dm_build_2) Dm_build_155(dm_build_157 []uint32) ([]int64, error) {
	dm_build_158 := dm_build_987(dm_build_156, dm_build_157)
	dm_build_159, dm_build_160 := dm_build_156.dm_build_36(dm_build_158)
	if dm_build_160 != nil {
		return nil, dm_build_160
	}
	return dm_build_159.([]int64), nil
}

func (dm_build_162 *dm_build_2) Close() error {
	if dm_build_162.dm_build_12 {
		return nil
	}

	dm_build_163 := dm_build_162.dm_build_3.Close()
	if dm_build_163 != nil {
		return dm_build_163
	}

	dm_build_162.dm_build_12 = true
	return nil
}

func (dm_build_165 *dm_build_2) dm_build_164(dm_build_166 *lob) (int64, error) {
	dm_build_167 := dm_build_750(dm_build_165, dm_build_166)
	dm_build_168, dm_build_169 := dm_build_165.dm_build_36(dm_build_167)
	if dm_build_169 != nil {
		return 0, dm_build_169
	}
	return dm_build_168.(int64), nil
}

func (dm_build_171 *dm_build_2) dm_build_170(dm_build_172 *lob, dm_build_173 int32, dm_build_174 int32) (*lobRetInfo, error) {
	dm_build_175 := dm_build_735(dm_build_171, dm_build_172, int(dm_build_173), int(dm_build_174))
	dm_build_176, dm_build_177 := dm_build_171.dm_build_36(dm_build_175)
	if dm_build_177 != nil {
		return nil, dm_build_177
	}
	return dm_build_176.(*lobRetInfo), nil
}

func (dm_build_179 *dm_build_2) dm_build_178(dm_build_180 *DmBlob, dm_build_181 int32, dm_build_182 int32) ([]byte, error) {
	var dm_build_183 = make([]byte, dm_build_182)
	var dm_build_184 int32 = 0
	var dm_build_185 int32 = 0
	var dm_build_186 *lobRetInfo
	var dm_build_187 []byte
	var dm_build_188 error
	for dm_build_184 < dm_build_182 {
		dm_build_185 = dm_build_182 - dm_build_184
		if dm_build_185 > Dm_build_467 {
			dm_build_185 = Dm_build_467
		}
		dm_build_186, dm_build_188 = dm_build_179.dm_build_170(&dm_build_180.lob, dm_build_181+dm_build_184, dm_build_185)
		if dm_build_188 != nil {
			return nil, dm_build_188
		}
		dm_build_187 = dm_build_186.data
		if dm_build_187 == nil || len(dm_build_187) == 0 {
			break
		}
		Dm_build_998.Dm_build_1054(dm_build_183, int(dm_build_184), dm_build_187, 0, len(dm_build_187))
		dm_build_184 += int32(len(dm_build_187))
		if dm_build_180.readOver {
			break
		}
	}
	return dm_build_183, nil
}

func (dm_build_190 *dm_build_2) dm_build_189(dm_build_191 *DmClob, dm_build_192 int32, dm_build_193 int32) (string, error) {
	var dm_build_194 bytes.Buffer
	var dm_build_195 int32 = 0
	var dm_build_196 int32 = 0
	var dm_build_197 *lobRetInfo
	var dm_build_198 []byte
	var dm_build_199 string
	var dm_build_200 error
	for dm_build_195 < dm_build_193 {
		dm_build_196 = dm_build_193 - dm_build_195
		if dm_build_196 > Dm_build_467/2 {
			dm_build_196 = Dm_build_467 / 2
		}
		dm_build_197, dm_build_200 = dm_build_190.dm_build_170(&dm_build_191.lob, dm_build_192+dm_build_195, dm_build_196)
		if dm_build_200 != nil {
			return "", dm_build_200
		}
		dm_build_198 = dm_build_197.data
		if dm_build_198 == nil || len(dm_build_198) == 0 {
			break
		}
		dm_build_199 = Dm_build_998.Dm_build_1155(dm_build_198, 0, len(dm_build_198), dm_build_191.serverEncoding, dm_build_190.dm_build_5)

		dm_build_194.WriteString(dm_build_199)
		var strLen = dm_build_197.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_199))
		}
		dm_build_195 += int32(strLen)
		if dm_build_191.readOver {
			break
		}
	}
	return dm_build_194.String(), nil
}

func (dm_build_202 *dm_build_2) dm_build_201(dm_build_203 *DmClob, dm_build_204 int, dm_build_205 string, dm_build_206 string) (int, error) {
	var dm_build_207 = Dm_build_998.Dm_build_1214(dm_build_205, dm_build_206, dm_build_202.dm_build_5)
	var dm_build_208 = 0
	var dm_build_209 = len(dm_build_207)
	var dm_build_210 = 0
	var dm_build_211 = 0
	var dm_build_212 = 0
	var dm_build_213 = dm_build_209/Dm_build_466 + 1
	var dm_build_214 byte = 0
	var dm_build_215 byte = 0x01
	var dm_build_216 byte = 0x02
	for i := 0; i < dm_build_213; i++ {
		dm_build_214 = 0
		if i == 0 {
			dm_build_214 |= dm_build_215
		}
		if i == dm_build_213-1 {
			dm_build_214 |= dm_build_216
		}
		dm_build_212 = dm_build_209 - dm_build_211
		if dm_build_212 > Dm_build_466 {
			dm_build_212 = Dm_build_466
		}

		setLobData := dm_build_899(dm_build_202, &dm_build_203.lob, dm_build_214, dm_build_204, dm_build_207, dm_build_208, dm_build_212)
		ret, err := dm_build_202.dm_build_36(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_210, nil
		} else {
			dm_build_204 += int(tmp)
			dm_build_210 += int(tmp)
			dm_build_211 += dm_build_212
			dm_build_208 += dm_build_212
		}
	}
	return dm_build_210, nil
}

func (dm_build_218 *dm_build_2) dm_build_217(dm_build_219 *DmBlob, dm_build_220 int, dm_build_221 []byte) (int, error) {
	var dm_build_222 = 0
	var dm_build_223 = len(dm_build_221)
	var dm_build_224 = 0
	var dm_build_225 = 0
	var dm_build_226 = 0
	var dm_build_227 = dm_build_223/Dm_build_466 + 1
	var dm_build_228 byte = 0
	var dm_build_229 byte = 0x01
	var dm_build_230 byte = 0x02
	for i := 0; i < dm_build_227; i++ {
		dm_build_228 = 0
		if i == 0 {
			dm_build_228 |= dm_build_229
		}
		if i == dm_build_227-1 {
			dm_build_228 |= dm_build_230
		}
		dm_build_226 = dm_build_223 - dm_build_225
		if dm_build_226 > Dm_build_466 {
			dm_build_226 = Dm_build_466
		}

		setLobData := dm_build_899(dm_build_218, &dm_build_219.lob, dm_build_228, dm_build_220, dm_build_221, dm_build_222, dm_build_226)
		ret, err := dm_build_218.dm_build_36(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_224, nil
		} else {
			dm_build_220 += int(tmp)
			dm_build_224 += int(tmp)
			dm_build_225 += dm_build_226
			dm_build_222 += dm_build_226
		}
	}
	return dm_build_224, nil
}

func (dm_build_232 *dm_build_2) dm_build_231(dm_build_233 *lob, dm_build_234 int) (int64, error) {
	dm_build_235 := dm_build_761(dm_build_232, dm_build_233, dm_build_234)
	dm_build_236, dm_build_237 := dm_build_232.dm_build_36(dm_build_235)
	if dm_build_237 != nil {
		return dm_build_233.length, dm_build_237
	}
	return dm_build_236.(int64), nil
}

func (dm_build_239 *dm_build_2) dm_build_238(dm_build_240 []interface{}, dm_build_241 []interface{}, dm_build_242 int) bool {
	var dm_build_243 = false
	dm_build_240[dm_build_242] = dm_build_241[dm_build_242]

	if binder, ok := dm_build_241[dm_build_242].(iOffRowBinder); ok {
		dm_build_243 = true
		dm_build_240[dm_build_242] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_239.dm_build_5) {
			dm_build_240[dm_build_242] = &lobCtl{lob.buildCtlData()}
			dm_build_243 = false
		}
	} else {
		dm_build_240[dm_build_242] = dm_build_241[dm_build_242]
	}
	return dm_build_243
}

func (dm_build_245 *dm_build_2) dm_build_244(dm_build_246 *DmStatement, dm_build_247 parameter, dm_build_248 int, dm_build_249 iOffRowBinder) error {
	var dm_build_250 = Dm_build_1283()
	dm_build_249.read(dm_build_250)
	var dm_build_251 = 0
	for !dm_build_249.isReadOver() || dm_build_250.Dm_build_1284() > 0 {
		if !dm_build_249.isReadOver() && dm_build_250.Dm_build_1284() < Dm_build_466 {
			dm_build_249.read(dm_build_250)
		}
		if dm_build_250.Dm_build_1284() > Dm_build_466 {
			dm_build_251 = Dm_build_466
		} else {
			dm_build_251 = dm_build_250.Dm_build_1284()
		}

		putData := dm_build_870(dm_build_245, dm_build_246, int16(dm_build_248), dm_build_250, int32(dm_build_251))
		_, err := dm_build_245.dm_build_36(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_253 *dm_build_2) dm_build_252() ([]byte, error) {
	var dm_build_254 error
	if dm_build_253.dm_build_9 == nil {
		if dm_build_253.dm_build_9, dm_build_254 = security.NewClientKeyPair(); dm_build_254 != nil {
			return nil, dm_build_254
		}
	}
	return security.Bn2Bytes(dm_build_253.dm_build_9.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_256 *dm_build_2) dm_build_255() (*security.DhKey, error) {
	var dm_build_257 error
	if dm_build_256.dm_build_9 == nil {
		if dm_build_256.dm_build_9, dm_build_257 = security.NewClientKeyPair(); dm_build_257 != nil {
			return nil, dm_build_257
		}
	}
	return dm_build_256.dm_build_9, nil
}

func (dm_build_259 *dm_build_2) dm_build_258(dm_build_260 int, dm_build_261 []byte, dm_build_262 string, dm_build_263 int) (dm_build_264 error) {
	if dm_build_260 > 0 && dm_build_260 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_261 != nil {
		dm_build_259.dm_build_6, dm_build_264 = security.NewSymmCipher(dm_build_260, dm_build_261)
	} else if dm_build_260 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_259.dm_build_6, dm_build_264 = security.NewThirdPartCipher(dm_build_260, dm_build_261, dm_build_262, dm_build_263); dm_build_264 != nil {
			dm_build_264 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_264.Error()).throw()
		}
	}
	return
}

func (dm_build_266 *dm_build_2) dm_build_265() (dm_build_267 error) {
	if dm_build_266.dm_build_5.sslEncrypt == 2 {

		socket := &Dm_build_297{
			dm_build_298: dm_build_266.dm_build_3,
			dm_build_300: true,
		}
		if dm_build_267 = socket.Dm_build_269(context.Background(), dm_build_266.dm_build_5); dm_build_267 != nil {
			return dm_build_267
		}
		dm_build_266.dm_build_3 = socket
	} else if dm_build_266.dm_build_5.sslEncrypt == 1 || dm_build_266.dm_build_5.sslEncrypt == 4 {
		socket := &Dm_build_297{
			dm_build_298: dm_build_266.dm_build_3,
			dm_build_300: false,
		}
		if dm_build_267 = socket.Dm_build_269(context.Background(), dm_build_266.dm_build_5); dm_build_267 != nil {
			return dm_build_267
		}
		dm_build_266.dm_build_3 = socket
	}

	return
}

type Dm_build_268 interface {
	io.ReadWriteCloser

	Dm_build_269(ctx context.Context, dmConn *DmConnection) error
}

type Dm_build_270 struct {
	dm_build_271 net.Conn
}

func (dm_build_273 *Dm_build_270) Dm_build_269(dm_build_274 context.Context, dm_build_275 *DmConnection) (dm_build_276 error) {
	dialsLock.RLock()
	dm_build_277, dm_build_278 := dials[dm_build_275.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_278 {
		dm_build_273.dm_build_271, dm_build_276 = dm_build_277(dm_build_274, dm_build_275.dmConnector.host+":"+strconv.Itoa(int(dm_build_275.dmConnector.port)))
	} else {
		dm_build_273.dm_build_271, dm_build_276 = dm_build_273.dm_build_291(dm_build_275.dmConnector.host+":"+strconv.Itoa(int(dm_build_275.dmConnector.port)), time.Duration(dm_build_275.dmConnector.socketTimeout)*time.Second)
	}
	return dm_build_276
}

func (dm_build_280 *Dm_build_270) Close() error {
	return dm_build_280.dm_build_271.Close()
}

func (dm_build_282 *Dm_build_270) Write(dm_build_283 []byte) (dm_build_284 int, dm_build_285 error) {
	return dm_build_282.dm_build_271.Write(dm_build_283)
}

func (dm_build_287 *Dm_build_270) Read(dm_build_288 []byte) (dm_build_289 int, dm_build_290 error) {
	return dm_build_287.dm_build_271.Read(dm_build_288)
}

func (dm_build_292 *Dm_build_270) dm_build_291(dm_build_293 string, dm_build_294 time.Duration) (net.Conn, error) {
	dm_build_295, dm_build_296 := net.DialTimeout("tcp", dm_build_293, dm_build_294)
	if dm_build_296 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_293).throw()
	}

	if tcpConn, ok := dm_build_295.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1)
		tcpConn.SetNoDelay(true)

	}
	return dm_build_295, nil
}

type Dm_build_297 struct {
	dm_build_298 Dm_build_268

	dm_build_299 tls.Conn

	dm_build_300 bool
}

func (dm_build_302 *Dm_build_297) Dm_build_269(dm_build_303 context.Context, dm_build_304 *DmConnection) (dm_build_305 error) {
	var dm_build_306 *tls.Conn
	dm_build_306, dm_build_305 = dm_build_302.dm_build_319((dm_build_302.dm_build_298).(*Dm_build_270).dm_build_271, dm_build_304)
	if dm_build_305 != nil {
		return
	}
	dm_build_302.dm_build_299 = *dm_build_306
	return nil
}

func (dm_build_308 *Dm_build_297) Close() error {
	if err := dm_build_308.dm_build_299.Close(); err != nil {
		return err
	}
	return dm_build_308.dm_build_298.Close()
}

func (dm_build_310 *Dm_build_297) Write(dm_build_311 []byte) (dm_build_312 int, dm_build_313 error) {
	if dm_build_310.dm_build_300 {
		return dm_build_310.dm_build_298.Write(dm_build_311)
	} else {
		return dm_build_310.dm_build_299.Write(dm_build_311)
	}
}

func (dm_build_315 *Dm_build_297) Read(dm_build_316 []byte) (dm_build_317 int, dm_build_318 error) {
	if dm_build_315.dm_build_300 {
		return dm_build_315.dm_build_298.Read(dm_build_316)
	} else {
		return dm_build_315.dm_build_299.Read(dm_build_316)
	}
}

func (dm_build_320 *Dm_build_297) dm_build_319(dm_build_321 net.Conn, dm_build_322 *DmConnection) (*tls.Conn, error) {
	var dm_build_323 = &tls.Config{
		InsecureSkipVerify: true,
	}

	if dm_build_322.sslEncrypt == 1 || dm_build_322.sslEncrypt == 2 {

		if customCA, err := ioutil.ReadFile(dm_build_322.dmConnector.sslCaPath); err == nil {
			dm_build_323.RootCAs, err = x509.SystemCertPool()
			if err != nil {

				dm_build_323.RootCAs = x509.NewCertPool()
			}
			dm_build_323.RootCAs.AppendCertsFromPEM(customCA)
		}

		if dm_build_322.dmConnector.sslCertPath == "" || dm_build_322.dmConnector.sslKeyPath == "" {

			return nil, errors.New("sslPath can not be empty!")
		}

		cer, err := dm_build_320.dm_build_325(dm_build_322.dmConnector.sslCertPath, dm_build_322.dmConnector.sslKeyPath, dm_build_322.dmConnector.sslPwd)
		if err != nil {
			return nil, err
		}

		dm_build_323.Certificates = []tls.Certificate{cer}

		dm_build_323.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {

			certs := make([]*x509.Certificate, len(rawCerts))
			for i, asn1Data := range rawCerts {
				cert, err := x509.ParseCertificate(asn1Data)
				if err != nil {
					return fmt.Errorf("tls: failed to parse certificate from server: %w", err)
				}
				certs[i] = cert
			}

			opts := x509.VerifyOptions{
				Roots:       dm_build_323.RootCAs,
				CurrentTime: time.Now(),

				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range certs[1:] {
				opts.Intermediates.AddCert(cert)
			}
			if _, err := certs[0].Verify(opts); err != nil {
				return err
			}
			return nil
		}
	}

	dm_build_324 := tls.Client(dm_build_321, dm_build_323)
	if err := dm_build_324.Handshake(); err != nil {
		return nil, err
	}
	return dm_build_324, nil

}

func (dm_build_326 *Dm_build_297) dm_build_325(dm_build_327, dm_build_328, dm_build_329 string) (tls.Certificate, error) {

	dm_build_330, err := ioutil.ReadFile(dm_build_327)
	if err != nil {
		return tls.Certificate{}, err
	}
	dm_build_332, err := ioutil.ReadFile(dm_build_328)
	if err != nil {
		return tls.Certificate{}, err
	}

	dm_build_334, _ := pem.Decode(dm_build_332)
	if dm_build_334 == nil {
		return tls.Certificate{}, fmt.Errorf("tls: failed to find any PEM data in key input")
	}

	if x509.IsEncryptedPEMBlock(dm_build_334) {

		dm_build_332, err = x509.DecryptPEMBlock(dm_build_334, []byte(dm_build_329))
		if err != nil {
			return tls.Certificate{}, err
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(dm_build_332)
		if err != nil {

			key, err2 := x509.ParsePKCS8PrivateKey(dm_build_332)
			if err2 != nil {
				return tls.Certificate{}, err2
			}
			dm_build_332, _ = x509.MarshalPKCS8PrivateKey(key)
		} else {
			dm_build_332 = x509.MarshalPKCS1PrivateKey(privateKey)
		}

		dm_build_332 = pem.EncodeToMemory(&pem.Block{
			Type:  dm_build_334.Type,
			Bytes: dm_build_332,
		})
	} else if dm_build_334.Type == "ENCRYPTED PRIVATE KEY" {

		return tls.Certificate{}, fmt.Errorf("PKCS#8 didn't supported yet")
	}

	dm_build_335, err := tls.X509KeyPair(dm_build_330, dm_build_332)
	if err != nil {
		return tls.Certificate{}, err
	}

	return dm_build_335, nil
}
