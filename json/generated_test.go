// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package json

import (
	"reflect"
	"testing"

	"github.com/coveo/gotemplate/errors"
	"github.com/stretchr/testify/assert"
)

var (
	strFixture = String("Hello World, I'm Foo Bar!").Split(" ")
	strList    = jsonListHelper.NewStringList
)

func Test_list_Append(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      jsonIList
		values []interface{}
		want   jsonIList
	}{
		{"Empty", jsonList{}, []interface{}{1, 2, 3}, jsonList{1, 2, 3}},
		{"List of int", jsonList{1, 2, 3}, []interface{}{4, 5}, jsonList{1, 2, 3, 4, 5}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, jsonList{"Hello", "World,", "I'm", "Foo", "Bar!", "That's all folks!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Append(tt.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Append():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Prepend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      jsonIList
		values []interface{}
		want   jsonIList
	}{
		{"Empty", jsonList{}, []interface{}{1, 2, 3}, jsonList{1, 2, 3}},
		{"List of int", jsonList{1, 2, 3}, []interface{}{4, 5}, jsonList{4, 5, 1, 2, 3}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, jsonList{"That's all folks!", "Hello", "World,", "I'm", "Foo", "Bar!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Prepend(tt.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Prepend():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_AsArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonIList
		want []interface{}
	}{
		{"Empty List", jsonList{}, []interface{}{}},
		{"List of int", jsonList{1, 2, 3}, []interface{}{1, 2, 3}},
		{"List of string", strFixture, []interface{}{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.AsArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.AsList():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_JsonList_Strings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonIList
		want StringArray
	}{
		{"Empty List", jsonList{}, StringArray{}},
		{"List of int", jsonList{1, 2, 3}, StringArray{"1", "2", "3"}},
		{"List of string", strFixture, StringArray{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Strings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Capacity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonIList
		want int
	}{
		{"Empty List with 100 spaces", jsonListHelper.CreateList(0, 100), 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Capacity(); got != tt.want {
				t.Errorf("JsonList.Capacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonIList
		want jsonIList
	}{
		{"Empty List", jsonList{}, jsonList{}},
		{"List of int", jsonList{1, 2, 3}, jsonList{1, 2, 3}},
		{"List of string", strFixture, jsonList{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Clone():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l       jsonIList
		indexes []int
		want    interface{}
	}{
		{"Empty List", jsonList{}, []int{0}, nil},
		{"Negative index", jsonList{}, []int{-1}, nil},
		{"List of int", jsonList{1, 2, 3}, []int{0}, 1},
		{"List of string", strFixture, []int{1}, "World,"},
		{"Get last", strFixture, []int{-1}, "Bar!"},
		{"Get before last", strFixture, []int{-2}, "Foo"},
		{"A way to before last", strFixture, []int{-12}, nil},
		{"Get nothing", strFixture, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Count(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonIList
		want int
	}{
		{"Empty List", jsonList{}, 0},
		{"List of int", jsonList{1, 2, 3}, 3},
		{"List of string", strFixture, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Count(); got != tt.want {
				t.Errorf("JsonList.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []int
		want    jsonIList
		wantErr bool
	}{
		{"Empty", nil, jsonList{}, false},
		{"With nil elements", []int{10}, make(jsonList, 10), false},
		{"With capacity", []int{0, 10}, make(jsonList, 0, 10), false},
		{"Too much args", []int{0, 10, 1}, nil, true},
	}
	for _, tt := range tests {
		var err error
		t.Run(tt.name, func(t *testing.T) {
			defer func() { err = errors.Trap(err, recover()) }()
			got := jsonListHelper.CreateList(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateList():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if got.Capacity() != tt.want.Capacity() {
				t.Errorf("CreateList() capacity:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got.Capacity(), tt.want.Capacity())
			}
		})
		if (err != nil) != tt.wantErr {
			t.Errorf("CreateList() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func Test_list_Create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonList
		args []int
		want jsonIList
	}{
		{"Empty", nil, nil, jsonList{}},
		{"Existing List", jsonList{1, 2}, nil, jsonList{}},
		{"With Empty spaces", jsonList{1, 2}, []int{5}, jsonList{nil, nil, nil, nil, nil}},
		{"With Capacity", jsonList{1, 2}, []int{0, 5}, jsonListHelper.CreateList(0, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.Create(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Create():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if got.Capacity() != tt.want.Capacity() {
				t.Errorf("JsonList.Create() capacity:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got.Capacity(), tt.want.Capacity())
			}
		})
	}
}

func Test_list_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonList
		args []interface{}
		want jsonIList
	}{
		{"Empty", nil, nil, jsonList{}},
		{"Existing List", jsonList{1, 2}, nil, jsonList{}},
		{"With elements", jsonList{1, 2}, []interface{}{3, 4, 5}, jsonList{3, 4, 5}},
		{"With strings", jsonList{1, 2}, []interface{}{"Hello", "World"}, jsonList{"Hello", "World"}},
		{"With nothing", jsonList{1, 2}, []interface{}{}, jsonList{}},
		{"With nil", jsonList{1, 2}, nil, jsonList{}},
		{"Adding array", jsonList{1, 2}, []interface{}{jsonList{3, 4}}, jsonList{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.New(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Create():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_CreateDict(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l       jsonList
		args    []int
		want    jsonIDict
		wantErr bool
	}{
		{"Empty", nil, nil, jsonDict{}, false},
		{"With capacity", nil, []int{10}, jsonDict{}, false},
		{"With too much parameter", nil, []int{10, 1}, nil, true},
	}
	for _, tt := range tests {
		var err error
		t.Run(tt.name, func(t *testing.T) {
			defer func() { err = errors.Trap(err, recover()) }()
			got := tt.l.CreateDict(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.CreateDict():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
		if (err != nil) != tt.wantErr {
			t.Errorf("JsonList.CreateDict() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func Test_list_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonList
		args []interface{}
		want bool
	}{
		{"Empty List", nil, []interface{}{}, false},
		{"Search nothing", jsonList{1}, nil, true},
		{"Search nothing 2", jsonList{1}, []interface{}{}, true},
		{"Not there", jsonList{1}, []interface{}{2}, false},
		{"Included", jsonList{1, 2}, []interface{}{2}, true},
		{"Partially there", jsonList{1, 2}, []interface{}{2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Contains(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Contains():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if got := tt.l.Has(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Has():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_First_Last(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		l         jsonList
		wantFirst interface{}
		wantLast  interface{}
	}{
		{"Nil", nil, nil, nil},
		{"Empty", jsonList{}, nil, nil},
		{"One element", jsonList{1}, 1, 1},
		{"Many element ", jsonList{1, "two", 3.1415, "four"}, 1, "four"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.First(); !reflect.DeepEqual(got, tt.wantFirst) {
				t.Errorf("JsonList.First():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.wantFirst)
			}
			if got := tt.l.Last(); !reflect.DeepEqual(got, tt.wantLast) {
				t.Errorf("JsonList.Last():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.wantLast)
			}
		})
	}
}

func Test_list_Pop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		l        jsonList
		args     []int
		want     interface{}
		wantList jsonList
	}{
		{"Nil", nil, nil, nil, jsonList{}},
		{"Empty", jsonList{}, nil, nil, jsonList{}},
		{"Non existent", jsonList{}, []int{1}, nil, jsonList{}},
		{"Empty with args", jsonList{}, []int{1, 3}, jsonList{nil, nil}, jsonList{}},
		{"List with bad index", jsonList{0, 1, 2, 3, 4, 5}, []int{1, 3, 8}, jsonList{1, 3, nil}, jsonList{0, 2, 4, 5}},
		{"Pop last element", jsonList{0, 1, 2, 3, 4, 5}, nil, 5, jsonList{0, 1, 2, 3, 4}},
		{"Pop before last", jsonList{0, 1, 2, 3, 4, 5}, []int{-2}, 4, jsonList{0, 1, 2, 3, 5}},
		{"Pop first element", jsonList{0, 1, 2, 3, 4, 5}, []int{0}, 0, jsonList{1, 2, 3, 4, 5}},
		{"Pop all", jsonList{0, 1, 2, 3}, []int{0, 1, 2, 3}, jsonList{0, 1, 2, 3}, jsonList{}},
		{"Pop same element many time", jsonList{0, 1, 2, 3}, []int{1, 1, 2, 2}, jsonList{1, 1, 2, 2}, jsonList{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotL := tt.l.Pop(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Pop():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if !reflect.DeepEqual(gotL, tt.wantList) {
				t.Errorf("JsonList.Pop():\ngotList %[1]v (%[1]T)\n   want %[2]v (%[2]T)", gotL, tt.wantList)
			}
		})
	}
}

func Test_list_Intersect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonList
		args []interface{}
		want jsonList
	}{
		{"Empty List", nil, []interface{}{}, jsonList{}},
		{"Intersect nothing", jsonList{1}, nil, jsonList{}},
		{"Intersect nothing 2", jsonList{1}, []interface{}{}, jsonList{}},
		{"Not there", jsonList{1}, []interface{}{2}, jsonList{}},
		{"Included", jsonList{1, 2}, []interface{}{2}, jsonList{2}},
		{"Partially there", jsonList{1, 2}, []interface{}{2, 3}, jsonList{2}},
		{"With duplicates", jsonList{1, 2, 3, 4, 5, 4, 3, 2, 1}, []interface{}{3, 4, 5, 6, 7, 8, 7, 6, 5, 5, 4, 3}, jsonList{3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Intersect(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Intersect():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Union(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonList
		args []interface{}
		want jsonList
	}{
		{"Empty List", nil, []interface{}{}, jsonList{}},
		{"Intersect nothing", jsonList{1}, nil, jsonList{1}},
		{"Intersect nothing 2", jsonList{1}, []interface{}{}, jsonList{1}},
		{"Not there", jsonList{1}, []interface{}{2}, jsonList{1, 2}},
		{"Included", jsonList{1, 2}, []interface{}{2}, jsonList{1, 2}},
		{"Partially there", jsonList{1, 2}, []interface{}{2, 3}, jsonList{1, 2, 3}},
		{"With duplicates", jsonList{1, 2, 3, 4, 5, 4, 3, 2, 1}, []interface{}{8, 7, 6, 5, 6, 7, 8}, jsonList{1, 2, 3, 4, 5, 8, 7, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Union(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Union():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Without(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    jsonList
		args []interface{}
		want jsonList
	}{
		{"Empty List", nil, []interface{}{}, jsonList{}},
		{"Remove nothing", jsonList{1}, nil, jsonList{1}},
		{"Remove nothing 2", jsonList{1}, []interface{}{}, jsonList{1}},
		{"Not there", jsonList{1}, []interface{}{2}, jsonList{1}},
		{"Included", jsonList{1, 2}, []interface{}{2}, jsonList{1}},
		{"Partially there", jsonList{1, 2}, []interface{}{2, 3}, jsonList{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Without(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Without():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l, want jsonList
	}{
		{"Empty List", nil, jsonList{}},
		{"Remove nothing", jsonList{1}, jsonList{1}},
		{"Duplicates following", jsonList{1, 1, 2, 3}, jsonList{1, 2, 3}},
		{"Duplicates not following", jsonList{1, 2, 3, 1, 2, 3, 4}, jsonList{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Unique(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Unique():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}
func Test_list_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l, want jsonIList
	}{
		{"Empty List", jsonList{}, jsonList{}},
		{"List of int", jsonList{1, 2, 3}, jsonList{3, 2, 1}},
		{"List of string", strFixture, jsonList{"Bar!", "Foo", "I'm", "World,", "Hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l.Clone()
			if got := l.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Reverse():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Set(t *testing.T) {
	t.Parallel()

	type args struct {
		i int
		v interface{}
	}
	tests := []struct {
		name    string
		l       jsonIList
		args    args
		want    jsonIList
		wantErr bool
	}{
		{"Empty", jsonList{}, args{2, 1}, jsonList{nil, nil, 1}, false},
		{"List of int", jsonList{1, 2, 3}, args{0, 10}, jsonList{10, 2, 3}, false},
		{"List of string", strFixture, args{2, "You're"}, jsonList{"Hello", "World,", "You're", "Foo", "Bar!"}, false},
		{"Negative", jsonList{}, args{-1, "negative value"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Clone().Set(tt.args.i, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonList.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Set():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

var mapFixture = map[string]interface{}{
	"int":     123,
	"float":   1.23,
	"string":  "Foo bar",
	"list":    []interface{}{1, "two"},
	"listInt": []int{1, 2, 3},
	"map": map[string]interface{}{
		"sub1": 1,
		"sub2": "two",
	},
	"mapInt": map[int]interface{}{
		1: 1,
		2: "two",
	},
}

var dictFixture = jsonDict(jsonDictHelper.AsDictionary(mapFixture).AsMap())

func dumpKeys(t *testing.T, d1, d2 jsonIDict) {
	t.Parallel()

	for key := range d1.AsMap() {
		v1, v2 := d1.Get(key), d2.Get(key)
		if reflect.DeepEqual(v1, v2) {
			continue
		}
		t.Logf("'%[1]v' = %[2]v (%[2]T) vs %[3]v (%[3]T)", key, v1, v2)
	}
}

func Test_dict_AsMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    jsonDict
		want map[string]interface{}
	}{
		{"Nil", nil, nil},
		{"Empty", jsonDict{}, map[string]interface{}{}},
		{"Map", dictFixture, map[string]interface{}(dictFixture)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.AsMap():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    jsonDict
		keys []interface{}
		want jsonIDict
	}{
		{"Nil", nil, nil, jsonDict{}},
		{"Empty", jsonDict{}, nil, jsonDict{}},
		{"Map", dictFixture, nil, dictFixture},
		{"Map with Fields", dictFixture, []interface{}{"int", "list"}, jsonDict(dictFixture).Omit("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.Clone(tt.keys...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Clone():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}

			// Ensure that the copy is distinct from the original
			got.Set("NewFields", "Test")
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Should be different:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if !got.Has("NewFields") || !reflect.DeepEqual(got.Get("NewFields"), "Test") {
				t.Errorf("Element has not been added")
			}
		})
	}
}

func Test_JsonDict_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		d            jsonDict
		args         []int
		want         jsonIList
		wantCount    int
		wantCapacity int
	}{
		{"Nil", nil, nil, jsonList{}, 0, 0},
		{"Empty", jsonDict{}, nil, jsonList{}, 0, 0},
		{"Map", dictFixture, nil, jsonList{}, 0, 0},
		{"Map with size", dictFixture, []int{3}, jsonList{nil, nil, nil}, 3, 3},
		{"Map with capacity", dictFixture, []int{0, 10}, jsonList{}, 0, 10},
		{"Map with size&capacity", dictFixture, []int{3, 10}, jsonList{nil, nil, nil}, 3, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.CreateList(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.CreateList() = %v, want %v", got, tt.want)
			}
			if got.Count() != tt.wantCount || got.Capacity() != tt.wantCapacity {
				t.Errorf("JsonDict.CreateList() size: %d, %d vs %d, %d", got.Count(), got.Capacity(), tt.wantCount, tt.wantCapacity)
			}
		})
	}
}

func Test_dict_Create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		d       jsonDict
		args    []int
		want    jsonIDict
		wantErr bool
	}{
		{"Empty", nil, nil, jsonDict{}, false},
		{"With capacity", nil, []int{10}, jsonDict{}, false},
		{"With too much parameter", nil, []int{10, 1}, nil, true},
	}
	for _, tt := range tests {
		var err error
		t.Run(tt.name, func(t *testing.T) {
			defer func() { err = errors.Trap(err, recover()) }()
			got := tt.d.Create(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Create():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
		if (err != nil) != tt.wantErr {
			t.Errorf("JsonList.Create() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func Test_dict_Default(t *testing.T) {
	t.Parallel()

	type args struct {
		key    interface{}
		defVal interface{}
	}
	tests := []struct {
		name string
		d    jsonDict
		args args
		want interface{}
	}{
		{"Empty", nil, args{"Foo", "Bar"}, "Bar"},
		{"Map int", dictFixture, args{"int", 1}, 123},
		{"Map float", dictFixture, args{"float", 1}, 1.23},
		{"Map Non existant", dictFixture, args{"Foo", "Bar"}, "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Default(tt.args.key, tt.args.defVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		key  interface{}
		keys []interface{}
	}
	tests := []struct {
		name    string
		d       jsonDict
		args    args
		want    jsonIDict
		wantErr bool
	}{
		{"Empty", nil, args{}, jsonDict{}, true},
		{"Map", dictFixture, args{}, dictFixture, true},
		{"Non existant key", dictFixture, args{"Test", nil}, dictFixture, true},
		{"Map with keys", dictFixture, args{"int", []interface{}{"list"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), false},
		{"Map with keys + non existant", dictFixture, args{"int", []interface{}{"list", "Test"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got, err := d.Delete(tt.args.key, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDict.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Delete():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
		})
	}
}

func Test_dict_Flush(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    jsonDict
		keys []interface{}
		want jsonIDict
	}{
		{"Empty", nil, nil, jsonDict{}},
		{"Map", dictFixture, nil, jsonDict{}},
		{"Non existant key", dictFixture, []interface{}{"Test"}, dictFixture},
		{"Map with keys", dictFixture, []interface{}{"int", "list"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
		{"Map with keys + non existant", dictFixture, []interface{}{"int", "list", "Test"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Flush(tt.keys...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Flush():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
			if !reflect.DeepEqual(d, got) {
				t.Errorf("Should be equal after: %v, want %v", d, got)
				dumpKeys(t, d, got)
			}
		})
	}
}

func Test_dict_Keys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    jsonDict
		want jsonIList
	}{
		{"Empty", nil, jsonList{}},
		{"Map", dictFixture, strList("float", "int", "list", "listInt", "map", "mapInt", "string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.GetKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.GetKeys():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Merge(t *testing.T) {
	t.Parallel()

	adding1 := jsonDict{
		"int":        1000,
		"Add1Int":    1,
		"Add1String": "string",
	}
	adding2 := jsonDict{
		"Add2Int":    1,
		"Add2String": "string",
		"map": map[string]interface{}{
			"sub1":   2,
			"newVal": "NewValue",
		},
	}
	type args struct {
		jsonDict jsonIDict
		dicts    []jsonIDict
	}
	tests := []struct {
		name string
		d    jsonDict
		args args
		want jsonIDict
	}{
		{"Empty", nil, args{nil, []jsonIDict{}}, jsonDict{}},
		{"Add map to empty", nil, args{dictFixture, []jsonIDict{}}, dictFixture},
		{"Add map to same map", dictFixture, args{dictFixture, []jsonIDict{}}, dictFixture},
		{"Add empty to map", dictFixture, args{nil, []jsonIDict{}}, dictFixture},
		{"Add new1 to map", dictFixture, args{adding1, []jsonIDict{}}, dictFixture.Clone().Merge(adding1)},
		{"Add new2 to map", dictFixture, args{adding2, []jsonIDict{}}, dictFixture.Clone().Merge(adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []jsonIDict{adding2}}, dictFixture.Clone().Merge(adding1, adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []jsonIDict{adding2}}, dictFixture.Clone().Merge(adding1).Merge(adding2)},
	}
	for _, tt := range tests {
		go t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Merge(tt.args.jsonDict, tt.args.dicts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Merge():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
		})
	}
}

func Test_dict_Values(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    jsonDict
		want jsonIList
	}{
		{"Empty", nil, jsonList{}},
		{"Map", dictFixture, jsonList{1.23, 123, jsonList{1, "two"}, jsonList{1, 2, 3}, jsonDict{"sub1": 1, "sub2": "two"}, jsonDict{"1": 1, "2": "two"}, "Foo bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.GetValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.GetValues():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Pop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		d          jsonDict
		args       []interface{}
		want       interface{}
		wantObject jsonIDict
	}{
		{"Nil", dictFixture, nil, nil, dictFixture},
		{"Pop one element", dictFixture, []interface{}{"float"}, 1.23, dictFixture.Omit("float")},
		{"Pop missing element", dictFixture, []interface{}{"undefined"}, nil, dictFixture},
		{"Pop element twice", dictFixture, []interface{}{"int", "int", "string"}, jsonList{123, 123, "Foo bar"}, dictFixture.Omit("int", "string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Pop(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Pop():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if !reflect.DeepEqual(d, tt.wantObject) {
				t.Errorf("JsonDict.Pop() object:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", d, tt.wantObject)
			}
		})
	}
}

func Test_dict_Add(t *testing.T) {
	t.Parallel()

	type args struct {
		key interface{}
		v   interface{}
	}
	tests := []struct {
		name string
		d    jsonDict
		args args
		want jsonIDict
	}{
		{"Empty", nil, args{"A", 1}, jsonDict{"A": 1}},
		{"With element", jsonDict{"A": 1}, args{"A", 2}, jsonDict{"A": jsonList{1, 2}}},
		{"With element, another value", jsonDict{"A": 1}, args{"B", 2}, jsonDict{"A": 1, "B": 2}},
		{"With list element", jsonDict{"A": jsonList{1, 2}}, args{"A", 3}, jsonDict{"A": jsonList{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Add(tt.args.key, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Set(t *testing.T) {
	t.Parallel()

	type args struct {
		key interface{}
		v   interface{}
	}
	tests := []struct {
		name string
		d    jsonDict
		args args
		want jsonIDict
	}{
		{"Empty", nil, args{"A", 1}, jsonDict{"A": 1}},
		{"With element", jsonDict{"A": 1}, args{"A", 2}, jsonDict{"A": 2}},
		{"With element, another value", jsonDict{"A": 1}, args{"B", 2}, jsonDict{"A": 1, "B": 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Set(tt.args.key, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Transpose(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    jsonDict
		want jsonIDict
	}{
		{"Empty", nil, jsonDict{}},
		{"Base", jsonDict{"A": 1}, jsonDict{"1": String("A")}},
		{"Multiple", jsonDict{"A": 1, "B": 2, "C": 1}, jsonDict{"1": strList("A", "C"), "2": String("B")}},
		{"List", jsonDict{"A": []int{1, 2, 3}, "B": 2, "C": 3}, jsonDict{"1": String("A"), "2": strList("A", "B"), "3": strList("A", "C")}},
		{"Complex", jsonDict{"A": jsonDict{"1": 1, "2": 2}, "B": 2, "C": 3}, jsonDict{"2": String("B"), "3": String("C"), AsStdString(jsonDict{"1": 1, "2": 2}): String("A")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDict.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_JsonList_Get(t *testing.T) {
	type args struct {
		indexes []int
	}
	tests := []struct {
		name string
		l    jsonList
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_JsonList_TypeName(t *testing.T) {
	tests := []struct {
		name string
		l    jsonList
		want String
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.TypeName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonList.TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Json_TypeName(t *testing.T) {
	t.Run("list", func(t *testing.T) { assert.Equal(t, jsonList{}.TypeName(), String("Json")) })
	t.Run("dict", func(t *testing.T) { assert.Equal(t, jsonDict{}.TypeName(), String("Json")) })
}

func Test_Json_GetHelper(t *testing.T) {
	t.Run("list", func(t *testing.T) {
		gotD, gotL := jsonList{}.GetHelpers()
		assert.Equal(t, gotD.CreateDictionary().TypeName(), jsonDictHelper.CreateDictionary().TypeName())
		assert.Equal(t, gotL.CreateList().TypeName(), jsonListHelper.CreateList().TypeName())
	})
	t.Run("dict", func(t *testing.T) {
		gotD, gotL := jsonDict{}.GetHelpers()
		assert.Equal(t, gotD.CreateDictionary().TypeName(), jsonDictHelper.CreateDictionary().TypeName())
		assert.Equal(t, gotL.CreateList().TypeName(), jsonListHelper.CreateList().TypeName())
	})
}
