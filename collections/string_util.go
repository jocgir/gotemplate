package collections

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig"

	"github.com/coveo/gotemplate/errors"
)

var (
	Must = errors.Must
	eol  = fmt.Sprintln()
)

// FormatMessage analyses the arguments to determine if printf or println should be used.
func FormatMessage(args ...interface{}) String {
	switch len(args) {
	case 0:
		return ""
	case 1:
		return AsString(args[0])
	default:
		if format, args := AsString(args[0]), args[1:]; format.Contains("%") {
			if result := fmt.Sprintf(format.Str(), args...); !strings.Contains(result, "%!") {
				return String(result)
			}
		}
		return String(fmt.Sprintln(args...)).TrimSuffix(eol)
	}
}

// Interface2string returns the string representation of any interface.
func Interface2string(str interface{}) string {
	switch str := str.(type) {
	case string:
		return str
	default:
		return fmt.Sprintf("%v", str)
	}
}

// Concat returns a string with all string representation of object concatenated without space.
func Concat(objects ...IString) String {
	var result string
	for _, object := range objects {
		result += AsStdString(object)
	}
	return String(result)
}

// ToStrings converts the supplied parameter into an array of string.
func ToStrings(args interface{}) StringArray {
	source := ToStdStrings(args)
	result := make(StringArray, len(source))
	for i := range source {
		result[i] = String(source[i])
	}
	return result
}

// ToStdStrings converts the supplied parameter into an array of string.
var ToStdStrings = sprig.GenericFuncMap()["toStrings"].(func(interface{}) []string)

// ToInterfaces converts an array of strings into an array of interfaces
func ToInterfaces(values ...string) []interface{} {
	result := make([]interface{}, len(values))
	for i := range values {
		result[i] = values[i]
	}
	return result
}

// SplitLines return a list of interface object for each line in the supplied content.
func SplitLines(content IString) StringArray { return AsString(content).Lines() }

// JoinLines concatenate the representation of supplied arguments as a string separated by newlines.
func JoinLines(objects ...IString) String { return String("\n").Join(objects) }

// Split2 returns left and right part of a split.
func Split2(source, sep IString) (String, String) { return AsString(source).Split2(sep) }

// Indent returns the indented version of the supplied string.
func Indent(s, indent IString) String { return AsString(s).Indent(indent) }

// IndentN returns the indented version (indent as a number of spaces) of the supplied string.
func IndentN(s IString, indent int) String { return AsString(s).IndentN(indent) }

// PrettyPrintStruct returns a readable version of an object.
func PrettyPrintStruct(object interface{}) string {
	var out string
	isZero := func(x interface{}) bool {
		return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
	}

	val := reflect.ValueOf(object)
	switch val.Kind() {
	case reflect.Interface:
		fallthrough
	case reflect.Ptr:
		val = val.Elem()
	}

	result := make([][2]string, 0, val.NumField())
	max := 0
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		if !field.CanInterface() {
			continue
		}

		itf := val.Field(i).Interface()
		if isZero(itf) {
			continue
		}

		itf = reflect.Indirect(val.Field(i)).Interface()
		value := AsString(itf).UnIndent().TrimSpace().Lines()
		if val.Field(i).Kind() == reflect.Struct {
			value[0] = "\n" + value.JoinLines().IndentN(4)
		} else if len(value) > 1 {
			value[0] += " ..."
		}
		result = append(result, [2]string{fieldType.Name, value[0].Str()})
		if len(fieldType.Name) > max {
			max = len(fieldType.Name)
		}
	}

	for _, entry := range result {
		out += fmt.Sprintf("%*s = %v\n", -max, entry[0], entry[1])
	}

	return out
}
