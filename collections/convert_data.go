package collections

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/coveo/gotemplate/errors"
)

// TypeConverters is used to register the available converters
var TypeConverters = make(map[string]func([]byte, interface{}) error)

// ConvertData returns a go representation of the supplied string (YAML, JSON or HCL)
func ConvertData(data string, out interface{}) (err error) {
	trySimplified := func() error {
		if strings.Count(data, "=") == 0 {
			return fmt.Errorf("Not simplifiable")
		}
		// Special case where we want to have a map and the supplied string is simplified such as "a = 10 b = string"
		// so we try transform the supplied string in valid YAML
		simplified := regexp.MustCompile(`[ \t]*=[ \t]*`).ReplaceAllString(data, ":")
		simplified = regexp.MustCompile(`[ \t]+`).ReplaceAllString(simplified, "\n")
		simplified = strings.Replace(simplified, ":", ": ", -1) + "\n"
		return ConvertData(simplified, out)
	}
	var errs errors.Array

	defer func() {
		if err == nil {
			// YAML converter returns a string if it encounter invalid data, so we check the result to ensure that is is different from the input.
			if out, isItf := out.(*interface{}); isItf && data == AsStdString(*out) && strings.ContainsAny(data, "=:{}") {
				if _, isString := (*out).(string); isString {
					if trySimplified() == nil && data != AsStdString(*out) {
						err = nil
						return
					}

					err = errs
					*out = nil
				}
			}
		} else {
			if _, e := TryAsList(out); e == nil && trySimplified() == nil {
				err = nil
			}
		}
	}()

	for _, key := range AsDictionary(TypeConverters).GetKeys() {
		err = TypeConverters[key.Str()]([]byte(data), out)
		if err == nil {
			return
		}
		errs = append(errs, fmt.Errorf("Trying %s: %v", key, err))
	}

	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}

// LoadData returns a go representation of the supplied file name (YAML, JSON or HCL)
func LoadData(filename string, out interface{}) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(filename); err == nil {
		return ConvertData(string(content), out)
	}
	return
}

// ToBash returns the bash 4 variable representation of value
func ToBash(value interface{}) String {
	return toBash(ToNativeRepresentation(value), 0)
}

func toBash(value interface{}, level int) (result String) {
	if v, isString := value.(string); isString {
		value = String(v)
	}

	if value, isString := value.(String); isString {
		result = value
		if result.ContainsAny(" \t\n[]()") {
			result = result.Format("%q")
		}
		return
	}

	if value, err := TryAsList(value); err == nil {
		results := value.Clone()
		for i := range results.AsArray() {
			results.Set(i, quote(results.Get(i)))
		}
		switch level {
		case 2:
			result = results.Join(",")
		default:
			result = results.Join(" ").Format("(%s)")
		}
		return
	}

	if value, err := TryAsDictionary(value); err == nil {
		results := value.CreateList(value.Count())
		vMap := value.AsMap()
		switch level {
		case 0:
			for i, key := range value.GetKeys() {
				k := key.Str()
				val := toBash(vMap[k], level+1)
				if _, err := TryAsList(vMap[k]); err == nil {
					results.Set(i, key.Format("declare -a %[1]s\n%[1]s=%[2]v", val))
				} else if _, err := TryAsDictionary(vMap[k]); err == nil {
					results.Set(i, key.Format("declare -A %[1]s\n%[1]s=%[2]v", val))
				} else {
					results.Set(i, key.Format("%s=%v", val))
				}
			}
			return results.Join("\n")
		case 1:
			for i, key := range value.GetKeys() {
				val := toBash(vMap[key.Str()], level+1)
				val = val.Replace(`$`, `\$`)
				results.Set(i, key.Format("[%s]=%s", val))
			}
			return results.Join(" ").Format("(%s)")
		default:
			for i, key := range value.GetKeys() {
				val := toBash(vMap[key.Str()], level+1)
				results.Set(i, key.Format("%s=%s", quote(val)))
			}
			return results.Join(",")
		}
	}
	return AsString(value)
}

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

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return Must(strconv.Atoi(AsStdString(value))).(int)

	case reflect.Int64, reflect.Uint64:
		return Must(strconv.ParseInt(AsStdString(value), 10, 64)).(int64)

	case reflect.Float32, reflect.Float64:
		return Must(strconv.ParseFloat(AsStdString(value), 64)).(float64)

	case reflect.Bool:
		return Must(strconv.ParseBool(AsStdString(value))).(bool)

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

			if options["omitempty"] && IsEmptyValue(val.Field(i)) {
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

func quote(i interface{}) String {
	s := AsString(i)
	if s.ContainsAny(" \t,[]()") {
		s = String(fmt.Sprintf("%q", s))
	}
	return s
}
