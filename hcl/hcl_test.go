package hcl

import (
	"reflect"
	"testing"

	"github.com/coveo/gotemplate/collections"
)

func Test_list_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    hclIList
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

func TestMarshalHCLVars(t *testing.T) {
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
		{"Map", args{hclDict{"a": 0, "bb": 1}, noIndent}, "a=0 bb=1"},
		{"Map (pretty)", args{hclDict{"a": 0, "bb": 1}, indent}, "a  = 0\nbb = 1"},
		{"Structure (pretty)", args{test{"name", 1}, indent}, "Name  = \"name\"\nValue = 1"},
		{"Structure Ptr (pretty)", args{&test{"name", 1}, indent}, "Name  = \"name\"\nValue = 1"},
		{"Array of 1 structure (pretty)", args{[]test{{"name", 1}}, indent}, "Name  = \"name\"\nValue = 1"},
		{"Array of 2 structures (pretty)", args{[]test{{"val1", 1}, {"val2", 1}}, indent}, "[\n  {\n    Name  = \"val1\"\n    Value = 1\n  },\n  {\n    Name  = \"val2\"\n    Value = 1\n  },\n]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := collections.ToNativeRepresentation(tt.args.value)
			if got, _ := marshalHCL(value, true, true, "", tt.args.indent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalHCLVars() = %v, want %v", got, tt.want)
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
		{"Map", AsStdString(dictFixture), dictFixture, false},
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
		{"Map", AsStdString(dictFixture), dictFixture.Native(), false},
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
