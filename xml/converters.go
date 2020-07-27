package customxml

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/errors"
	"github.com/coveo/gotemplate/v3/utils"
)

var (
	must = errors.Must
	trap = errors.Trap
)

func (l xmlList) String() string { result, _ := Marshal(l.AsArray()); return string(result) }
func (d xmlDict) String() string { result, _ := Marshal(d.AsMap()); return string(result) }

func (l xmlList) PrettyPrint() string {
	result, _ := MarshalIndent(l.AsArray(), "", "  ")
	return string(result)
}

func (d xmlDict) PrettyPrint() string {
	result, _ := MarshalIndent(d.AsMap(), "", "  ")
	return string(result)
}

// MarshalXML marshals an xml dictionary into XML.
func (d xmlDict) MarshalXML(encoder *xml.Encoder, start xml.StartElement) (err error) {
	defer func() { err = trap(err, recover()) }()

	must(encoder.EncodeToken(start))

	for _, key := range d.KeysAsString() {
		token := xml.StartElement{
			Name: xml.Name{Local: "Key"},
			Attr: []xml.Attr{
				{Name: xml.Name{Local: "name"}, Value: key.Str()},
			},
		}
		must(encoder.EncodeToken(token))
		must(encoder.Encode(transformIn(d.Get(key))))
		must(encoder.EncodeToken(token.End()))
	}

	return encoder.EncodeToken(start.End())
}

// UnmarshalXML restores back an xml dictionary.
func (d *xmlDict) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*d = xmlDict{}
	for {
		var x interface{}
		decoder.Decode(&x)
		fmt.Println(x)
		// var e dictElement

		// err := decoder.Decode(&e)
		// if err == io.EOF {
		// 	break
		// } else if err != nil {
		// 	return err
		// }

		// (*d)[e.Name] = e.Value
	}
	return nil
}

type Xml struct {
	Value string `xml:",innerxml"`
}

type Xml2 struct {
	Value interface{}
}

func transformIn(v interface{}) interface{} {
	if v == nil {
		return v
	} else if dict, err := collections.TryAsDictionary(v); err == nil {
		return xmlDictConvert(dict)
	} else if list, err := collections.TryAsList(v); err == nil {
		return xmlListConvert(list)
	}
	return v
}

func transform(out interface{}) {
	result := transformElement(reflect.ValueOf(out).Elem().Interface())
	if _, isMap := out.(*map[string]interface{}); isMap {
		// If the result is expected to be map[string]interface{}, we convert it back from internal dict type.
		result = result.(xmlIDict).Native()
	}
	reflect.ValueOf(out).Elem().Set(reflect.ValueOf(result))
}

func transformElement(source interface{}) interface{} {
	if value, err := xmlHelper.TryAsDictionary(source); err == nil {
		for _, key := range value.KeysAsString() {
			value.Set(key, transformElement(value.Get(key)))
		}
		source = value
	} else if value, err := xmlHelper.TryAsList(source); err == nil {
		for i, sub := range value.AsArray() {
			value.Set(i, transformElement(sub))
		}
		source = value
	}
	return source
}

// Optimizes the xml result by compressing self ending tags
func compress(result []byte, err error) ([]byte, error) {
	if err != nil {
		return result, err
	}
	result = []byte(reToken.ReplaceAllStringFunc(string(result), func(match string) string {
		matches, _ := utils.MultiMatch(match, reToken)
		if start := matches["start"]; start == matches["end"] {
			return fmt.Sprintf("<%s%s/>", start, matches["attr"])
		}
		return match
	}))
	return result, err
}

var reToken = regexp.MustCompile(`<(?P<start>\w+)(?P<attr>(?:\s+\w+=["'].*?["'])*)><\/(?P<end>\w+)>`)
