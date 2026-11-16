package dameng

import (
	"bytes"
	"testing"
)

func TestDialector_QuoteTo(t *testing.T) {
	testdatas := []struct {
		raw    string
		expect string
	}{
		{"datadase.tableUser", "\"datadase\".\"tableUser\""},
		{"datadase.table`User", "\"datadase\".\"table`User\""},
		{"`a`.`b`", "\"`a`\".\"`b`\""},
		{"`a`.b`", "\"`a`\".\"b`\""},
		{"a.`b`", "\"a\".\"`b`\""},
		{"`a`.b`c", "\"`a`\".\"b`c\""},
		{"`a`.`b`c`", "\"`a`\".\"`b`c`\""},
		{"`a`.b", "\"`a`\".\"b\""},
		{"`ab`", "\"`ab`\""},
		{"`a``b`", "\"`a``b`\""},
		{"`a```b`", "\"`a```b`\""},
		{"a`b", "\"a`b\""},
		{"ab", "\"ab\""},
		{"`a.b`", "\"`a\".\"b`\""},
		{"a.b", "\"a\".\"b\""},
	}

	dailor := Open("")
	for _, item := range testdatas {
		buf := &bytes.Buffer{}
		dailor.QuoteTo(buf, item.raw)
		if buf.String() != item.expect {
			t.Fatalf("quote %q fail, got %q, expect %q", item.raw, buf.String(), item.expect)
		}
	}
}

// BenchmarkDialector_QuoteTo
// Result:
// goos: darwin
// goarch: amd64
// pkg: gorm.io/driver/mysql
// cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
// BenchmarkDialector_QuoteTo               9184232               113.2 ns/op
// BenchmarkDialector_QuoteTo-2             9782818               112.3 ns/op
// BenchmarkDialector_QuoteTo-4            10726722               109.0 ns/op
// BenchmarkDialector_QuoteTo-8             9656778               113.1 ns/op
// BenchmarkDialector_QuoteTo-12           10729615               112.7 ns/op
func BenchmarkDialector_QuoteTo(b *testing.B) {
	dailor := Open("")
	buf := &bytes.Buffer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dailor.QuoteTo(buf, "datadase.table`User")
		buf.Reset()
	}
}

func TestBuildUrl(t *testing.T) {
	type args struct {
		user       string
		password   string
		host       string
		port       int
		urlOptions []map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"escape", args{"SYSDBA", "abc!@#$%^", "127.0.0.1", 54321, []map[string]string{{"escapeProcess": "true"}}}, "dm://SYSDBA:abc%21%40%23$%25%5E@127.0.0.1:54321?escapeProcess=true"},
		{"wrong", args{"SYSDBA", "abc!@#$%^", "127.0.0.1", 54321, []map[string]string{}}, "dm://SYSDBA:abc%21%40%23$%25%5E@127.0.0.1:54321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUrl(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.urlOptions...); got != tt.want {
				t.Errorf("BuildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
