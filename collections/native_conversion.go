package collections

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToNativeRepresentation converts any object to native (literals, maps, slices)
func ToNativeRepresentation(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	typ, val := reflect.TypeOf(value), reflect.ValueOf(value)
	if typ.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
		typ = val.Type()
	}
	switch typ.Kind() {
	case reflect.String:
		return reflect.ValueOf(value).String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return int(val.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return uint(val.Uint())

	case reflect.Int64:
		return val.Int()

	case reflect.Uint64:
		return val.Uint()

	case reflect.Float32, reflect.Float64:
		return must(strconv.ParseFloat(fmt.Sprint(value), 64)).(float64)

	case reflect.Bool:
		return must(strconv.ParseBool(fmt.Sprint(value))).(bool)

	case reflect.Slice, reflect.Array:
		result := make([]interface{}, val.Len())
		for i := range result {
			result[i] = ToNativeRepresentation(val.Index(i).Interface())
		}
		if len(result) == 1 && reflect.TypeOf(result[0]).Kind() == reflect.Map {
			// If the result is an array of one map, we just return the inner element
			return result[0]
		}
		return result

	case reflect.Map:
		result := make(map[string]interface{}, val.Len())
		for _, key := range val.MapKeys() {
			result[fmt.Sprintf("%v", key)] = ToNativeRepresentation(val.MapIndex(key).Interface())
		}
		return result

	case reflect.Struct:
		result := make(map[string]interface{}, typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			sf := typ.Field(i)
			if sf.Anonymous {
				t := sf.Type
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
				// If embedded, StructField.PkgPath is not a reliable
				// indicator of whether the field is exported.
				// See https://golang.org/issue/21122
				if !IsExported(t.Name()) && t.Kind() != reflect.Struct {
					// Ignore embedded fields of unexported non-struct collections.
					// Do not ignore embedded fields of unexported struct types
					// since they may have exported fields.
					continue
				}
			} else if sf.PkgPath != "" {
				// Ignore unexported non-embedded fields.
				continue
			}
			tag := sf.Tag.Get("hcl")
			if tag == "" {
				// If there is no hcl specific tag, we rely on json tag if there is
				tag = sf.Tag.Get("json")
			}
			if tag == "-" {
				continue
			}

			split := strings.Split(tag, ",")
			name := split[0]
			if name == "" {
				name = sf.Name
			}
			options := make(map[string]bool, len(split[1:]))
			for i := range split[1:] {
				options[split[i+1]] = true
			}

			if !IsExported(sf.Name) || options["omitempty"] && IsEmptyValue(val.Field(i)) {
				continue
			}

			if options["inline"] {
				for key, value := range ToNativeRepresentation(val.Field(i).Interface()).(map[string]interface{}) {
					result[key] = value
				}
			} else {
				result[name] = ToNativeRepresentation(val.Field(i).Interface())
			}
		}
		return result
	default:
		fmt.Fprintf(os.Stderr, "Unknown type %T %v : %v\n", value, typ.Kind(), value)
		return fmt.Sprintf("%v", value)
	}
}

// IsExported reports whether the identifier is exported.
func IsExported(id string) bool {
	r, _ := utf8.DecodeRuneInString(id)
	return unicode.IsUpper(r)
}
