package customxml

// The package is renamed because there is a strange vet error when we call it xml.
// Anyway, this library self registers, so it does not need to be referred directly in most cases.

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/collections/implementation"
)

// Expose xml public objects
var (
	_                   = collections.Register("xml", 0, Unmarshal)
	CopyToken           = xml.CopyToken
	Escape              = xml.Escape
	EscapeText          = xml.EscapeText
	NativeMarshal       = xml.Marshal
	NativeMarshalIndent = xml.MarshalIndent
	NewDecoder          = xml.NewDecoder
	NewEncoder          = xml.NewEncoder
	NativeUnmarshal     = xml.Unmarshal
)

// Marshal returns the XML encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	return compress(NativeMarshal(transformIn(v)))
}

// MarshalIndent returns the XML encoding of v with indentation.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return compress(NativeMarshalIndent(transformIn(v), prefix, indent))
}

// Unmarshal calls the native Unmarshal but transform the results
// to returns Dictionary and GenerecList instead of go native collections.
func Unmarshal(data []byte, out interface{}) (err error) {
	if out, isInterface := out.(*interface{}); isInterface {
		if len(bytes.TrimSpace(data)) == 0 {
			*out = nil
			return nil
		}

		x := map[reflect.Kind]interface{}{
			reflect.Bool:       func() interface{} { v := false; return &v }(),
			reflect.Int:        func() interface{} { v := int(0); return &v }(),
			reflect.Int8:       func() interface{} { v := int8(0); return &v }(),
			reflect.Int16:      func() interface{} { v := int16(0); return &v }(),
			reflect.Int32:      func() interface{} { v := int32(0); return &v }(),
			reflect.Int64:      func() interface{} { v := int64(0); return &v }(),
			reflect.Uint:       func() interface{} { v := uint(0); return &v }(),
			reflect.Uint8:      func() interface{} { v := uint8(0); return &v }(),
			reflect.Uint16:     func() interface{} { v := uint16(0); return &v }(),
			reflect.Uint32:     func() interface{} { v := uint32(0); return &v }(),
			reflect.Uint64:     func() interface{} { v := uint64(0); return &v }(),
			reflect.Uintptr:    func() interface{} { v := uintptr(0); return &v }(),
			reflect.Float32:    func() interface{} { v := float32(0); return &v }(),
			reflect.Float64:    func() interface{} { v := float64(0); return &v }(),
			reflect.Complex64:  func() interface{} { v := complex64(0); return &v }(),
			reflect.Complex128: func() interface{} { v := complex128(0); return &v }(),
			reflect.String:     func() interface{} { v := ""; return &v }(),
		}

		y := make(map[string]interface{})
		for kind, value := range x {
			var k interface{} = kind
			y[k.(fmt.Stringer).String()] = value
		}

		dec := xml.NewDecoder(bytes.NewBuffer(data))
		var start xml.StartElement
		if t, err := dec.Token(); err != nil {
			return err
		} else {
			start = t.(xml.StartElement)
		}

		if val, native := y[start.Name.Local]; native {
			if err = dec.DecodeElement(&val, &start); err != nil {
				return
			}
			*out = reflect.ValueOf(val).Elem().Interface()
		}
	} else if err = NativeUnmarshal(data, out); err != nil {
		return
	}

	transform(out)
	return
}

type (
	helperBase = implementation.BaseHelper
	helperList = implementation.ListHelper
	helperDict = implementation.DictHelper
)

var needConversionImpl = implementation.NeedConversion

//go:generate genny -pkg=customxml -in=../collections/implementation/generic.go -out=generated_impl.go gen "ListTypeName=List DictTypeName=Dictionary base=xml"
//go:generate genny -pkg=customxml -in=../collections/implementation/generic_test.go -out=generated_test.go gen "base=xml"
