package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_list_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want string
	}{
		{"Nil", nil, "[]\n"},
		{"Empty List", yamlList{}, "[]\n"},
		{"List of int", yamlList{1, 2, 3}, String(`
			- 1
			- 2
			- 3
			`).UnIndent()[1:]},
		{"List of string", strFixture, String(`
			- Hello
			- World,
			- I'm
			- Foo
			- Bar!
			`).UnIndent()[1:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("yamlList.String():\ngot:\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}

func Test_dict_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		want string
	}{
		{"nil", nil, "{}\n"},
		{"Map", dictFixture, String(`
			float: 1.23
			int: 123
			list:
			- 1
			- two
			listInt:
			- 1
			- 2
			- 3
			map:
			  sub1: 1
			  sub2: two
			mapInt:
			  "1": 1
			  "2": two
			string: Foo bar
			`).UnIndent()[1:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("yamlDict.String():\ngot:\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		yaml string
		want interface{}
	}{
		{"nil", "{}\n", yamlDict{}},
		{"Map", asStdString(dictFixture), dictFixture},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out interface{}
			err := Unmarshal([]byte(tt.yaml), &out)
			if err == nil && !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Unmarshal:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", out, tt.want)
			}
		})
	}
}

func TestUnmarshalWithError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		yaml string
	}{
		{"Error", "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out map[string]interface{}
			err := Unmarshal([]byte(tt.yaml), &out)
			if err == nil {
				t.Errorf("Unmarshal() expected error")
			}
		})
	}
}

func TestUnmarshalStrict(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		yaml    string
		want    interface{}
		wantErr bool
	}{
		{"nil", "{}\n", map[string]interface{}{}, false},
		{"Map", fmt.Sprint(dictFixture), dictFixture.Native(), false},
		{"Error", "Invalid", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out map[string]interface{}
			err := UnmarshalStrict([]byte(tt.yaml), &out)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Unmarshal:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", out, tt.want)
			}
		})
	}
}
