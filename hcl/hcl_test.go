package hcl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/coveo/gotemplate/collections"
)

func Test_list_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    hclList
		want string
	}{
		{"Nil", nil, "[]"},
		{"Empty list", hclList{}, "[]"},
		{"List of int", hclList{1, 2, 3}, "[1,2,3]"},
		{"List of string", strFixture, `["Hello","World,","I'm","Foo","Bar!"]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("hclList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    hclDict
		want string
	}{
		{"nil", nil, ""},
		{"Empty dict", hclDict{}, ""},
		{"Map", dictFixture, `float=1.23 int=123 list=[1,"two"] listInt=[1,2,3] map{sub1=1 sub2="two"} mapInt{"1"=1 "2"="two"} string="Foo bar"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("hclList.String():\n got %v\nwant %v", got, tt.want)
			}
		})
	}
}

func TestMarshalHCL(t *testing.T) {
	t.Parallel()

	type test struct {
		Name  string `hcl:",omitempty"`
		Value int    `hcl:",omitempty"`
	}
	const (
		noIndent = ""
		indent   = "  "
	)
	var testNilPtr *test

	type args struct {
		value  interface{}
		indent string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Integer", args{2, noIndent}, "2"},
		{"Boolean", args{true, noIndent}, "true"},
		{"String", args{"Hello world", noIndent}, `"Hello world"`},
		{"String with newline", args{"Hello\nworld\n", noIndent}, `"Hello\nworld\n"`},
		{"String with newline (pretty)", args{"Hello\n\"world\"\n", indent}, "<<-EOF\nHello\n\"world\"\nEOF"},
		{"Null value", args{nil, noIndent}, "null"},
		{"Null struct", args{testNilPtr, noIndent}, "null"},
		{"List of integer", args{[]int{0, 1, 2, 3}, noIndent}, "[0,1,2,3]"},
		{"One level map", args{hclDict{"a": hclDict{"b": 10}}, noIndent}, "a {b=10}"},
		{"One level map (pretty)", args{hclDict{"a": hclDict{"b": 10}}, indent}, "a {\n  b = 10\n}"},
		{"Two level map 1", args{hclDict{"a": hclDict{"b": hclDict{"c": 10, "d": 20}}}, noIndent}, "a b {c=10 d=20}"},
		{"Two level map 1 (pretty)", args{hclDict{"a": hclDict{"b": hclDict{"c": 10, "d": 20}}}, indent}, "a b {\n  c = 10\n  d = 20\n}"},
		{"Two level map 2", args{hclDict{"a": hclDict{"b": hclDict{"c": 10, "d": 20}}, "e": 30}, noIndent}, "a b {c=10 d=20} e=30"},
		{"Two level map 2 (pretty)", args{hclDict{"a": hclDict{"b": hclDict{"c": 10, "d": 20}}, "e": 30}, indent}, "e = 30\n\na b {\n  c = 10\n  d = 20\n}"},
		{"Two level map 3", args{hclDict{"a": hclDict{"b": 10, "c": hclDict{"d": 20, "e": 30}}}, noIndent}, "a {b=10 c{d=20 e=30}}"},
		{"Two level map 3 (pretty)", args{hclDict{"a": hclDict{"b": 10, "c": hclDict{"d": 20, "e": 30}}}, indent}, "a {\n  b = 10\n  \n  c {\n    d = 20\n    e = 30\n  }\n}"},
		{"Map", args{hclDict{"a": 0, "bb": 1}, noIndent}, "a=0 bb=1"},
		{"Map (pretty)", args{hclDict{"a": 0, "bb": 1}, indent}, "a  = 0\nbb = 1"},
		{"Structure (pretty)", args{test{"name", 1}, indent}, "Name  = \"name\"\nValue = 1"},
		{"Structure Ptr (pretty)", args{&test{"name", 1}, indent}, "Name  = \"name\"\nValue = 1"},
		{"Array of 1 structure (pretty)", args{[]test{{"name", 1}}, indent}, "Name  = \"name\"\nValue = 1"},
		{"Array of 2 structures (pretty)", args{[]test{{"val1", 1}, {"val2", 2}}, indent}, "[\n  {\n    Name  = \"val1\"\n    Value = 1\n  },\n  {\n    Name  = \"val2\"\n    Value = 2\n  },\n]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := collections.ToNativeRepresentation(tt.args.value)
			if got, _ := marshalHCL(value, true, true, "", tt.args.indent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalHCL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombo(t *testing.T) {
	t.Parallel()

	unindent := func(s string) string { return collections.String(s).UnIndent().TrimSpace().Str() }
	type r map[string]string

	tests := []struct {
		hcl     string
		results r
	}{
		{`1`, nil},
		{`3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862`, r{"": "3.141592653589793"}},
		{`true`, nil},
		{`false`, nil},
		{`"Hello world"`, nil},
		{`"Hello\nworld\n"`, r{
			"indent=true": unindent(`
				<<-EOF
				Hello
				world
				EOF
			`)},
		},
		{`[0, 1, 2, 3]`, nil},
		{`[0, "one", 2.0, "III"]`, nil},
		{`a = 0 b = "Hello"`, nil},
		{`a { b = 1 }`, nil},
		{`a b { c = 1 d = 2 }`, nil},
		{`a b { c = 1 d = 2 } e = 3`, nil},
		{`a { b = 1 c { d = 2 e = 3 }}`, nil},
	}
	for _, tt := range tests {
		t.Run(tt.hcl, func(t *testing.T) {
			var value interface{}
			if err := Unmarshal([]byte(tt.hcl), &value); err != nil {
				t.Errorf("Unmarshal(%v) => %v", tt.hcl, err)
			}
			native := collections.ToNativeRepresentation(value)

			for _, indent := range []string{"", "\t"} {
				for _, hcl := range []bool{false, true} {
					for _, head := range []bool{false, true} {
						out, err := marshalHCL(native, hcl, head, "", indent)
						if err != nil {
							t.Errorf("marshalHCL(%v, hcl=%v, head=%v, indent=%q) => %v", tt.hcl, hcl, head, indent, err)
						}
						result++
						index := result
						if index >= len(tt) {
							index = len(tt) - 1
						}
						if out != tt[index] {
							t.Errorf("marshalHCL(%v, hcl=%v, head=%v, indent=%q) =\n%v\nWant:\n%v", tt.hcl, hcl, head, indent, out, tt[index])
						}
					}
				}
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		hcl     string
		want    interface{}
		wantErr bool
	}{
		{"Empty", "", hclDict{}, false},
		{"Empty list", "[]", hclList{}, false},
		{"List of int", "[1,2,3]", hclList{1, 2, 3}, false},
		{"Array of map", "a { b { c { d = 1 e = 2 }}}", hclDict{"a": hclDict{"b": hclDict{"c": hclDict{"d": 1, "e": 2}}}}, false},
		{"Map", fmt.Sprint(dictFixture), dictFixture, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out interface{}
			err := Unmarshal([]byte(tt.hcl), &out)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Unmarshal:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", out, tt.want)
			}
		})
	}
}

func TestUnmarshalStrict(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		hcl     string
		want    interface{}
		wantErr bool
	}{
		{"Empty", "", map[string]interface{}{}, false},
		{"Empty list", "[]", nil, true},
		{"List of int", "[1,2,3]", nil, true},
		{"Array of map", "a { b { c { d = 1 e = 2 }}}", map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": map[string]interface{}{"d": 1, "e": 2}}}}, false},
		{"Map", fmt.Sprint(dictFixture), dictFixture.Native(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out map[string]interface{}
			err := Unmarshal([]byte(tt.hcl), &out)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Unmarshal:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", out, tt.want)
			}
		})
	}
}
