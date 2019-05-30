package data

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/stretchr/testify/assert"
)

func TestConvertData(t *testing.T) {
	type d = map[string]interface{}
	tests := []struct {
		name string
		data string
		want interface{}
		err  error
	}{
		{"Integer", "10", 10, nil},
		{"String", "hello", "hello", nil},
		{"List", "[1,2,3]", jsonList{1, 2, 3}, nil},
		{"YAML List", "[Hello, world]", yamlList{"Hello", "world"}, nil},
		{"Simple value", "a = 10", hclDict{"a": 10}, nil},
		{"YAML", "a: 10", yamlDict{"a": 10}, nil},
		{"HCL", `a = 10 b = "Foo"`, hclDict{"a": 10, "b": "Foo"}, nil},
		{"JSON", `{ "a": 10, "b": "Foo" }`, jsonDict{"a": 10, "b": "Foo"}, nil},
		{"Flexible", `a = 10 b = Foo`, hclDict{"a": 10, "b": "Foo"}, nil},
		{"Invalid", "a = 'value", nil, fmt.Errorf("Trying json: invalid character 'a' looking for beginning of value\nTrying hcl: At 1:5: illegal char")},
		{"File", "test-fixture", yamlDict{"int": 10, "list": yamlList{"a", "b"}, "string": "Hello", "struct": yamlDict{"a": 1, "b": true}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out interface{}
			err := Convert(tt.data, &out)
			assert.Equal(t, tt.want, out)
			if tt.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err.Error())
			}
		})
	}
}

func TestToBash(t *testing.T) {
	type SubStruct struct {
		U int64
		I interface{}
	}
	type a struct {
		private int
		I       int
		F       float64
		S       string
		A       []interface{}
		M       map[string]interface{}
		SS      SubStruct
	}
	tests := []struct {
		name string
		args interface{}
		want interface{}
	}{
		{"Struct conversion", a{
			private: 0,
			I:       123,
			F:       1.23,
			S:       "123",
			A:       []interface{}{1, "2"},
			M: map[string]interface{}{
				"a": "a",
				"b": 2,
			},
			SS: SubStruct{64, "Foo"},
		}, strings.TrimSpace(collections.UnIndent(`
		declare -a A
		A=(1 2)
		F=1.23
		I=123
		declare -A M
		M=([a]=a [b]=2)
		S=123
		declare -A SS
		SS=([I]=Foo [U]=64)
		`))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBash(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNativeRepresentation()\ngot : %q\nwant: %q", got, tt.want)
			}
		})
	}
}

func Test_quote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Simple value", "Foo", "Foo"},
		{"Simple value", "Foo Bar", `"Foo Bar"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quote(tt.arg); got != tt.want {
				t.Errorf("quote() = %v, want %v", got, tt.want)
			}
		})
	}
}
