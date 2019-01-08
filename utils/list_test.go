package utils

import (
	"reflect"
	"testing"

	"github.com/coveo/gotemplate/collections"
	"github.com/coveo/gotemplate/collections/implementation"
)

type IGenericList = collections.IGenericList
type list = implementation.ListTypeName

func TestFormatList(t *testing.T) {
	type args struct {
		format string
		v      interface{}
	}
	tests := []struct {
		name string
		args args
		want IGenericList
	}{
		{"quote", args{`"%v"`, []int{1, 2}}, list{`"1"`, `"2"`}},
		{"greating", args{"Hello %v", []int{1, 2}}, list{"Hello 1", "Hello 2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatList(tt.args.format, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeLists(t *testing.T) {
	tests := []struct {
		name string
		args []IGenericList
		want IGenericList
	}{
		{"Empty list", nil, nil},
		{"Simple list", []IGenericList{list{1, 2, 3}}, list{1, 2, 3}},
		{"Two lists", []IGenericList{list{1, 2, 3}, list{4, 5, 6}}, list{1, 2, 3, 4, 5, 6}},
		{"Three lists mixed", []IGenericList{list{"One", 2, "3"}, list{4, 5, 6}, list{"7", "8", "9"}}, list{"One", 2, "3", 4, 5, 6, "7", "8", "9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeLists(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
