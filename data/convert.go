package data

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/errors"
)

// UnmarshalString returns a go representation of the supplied string (YAML, JSON or HCL)
func UnmarshalString(data string, out interface{}) (err error) { return Unmarshal([]byte(data), out) }

// Unmarshal returns a go representation of the supplied buffer (YAML, JSON or HCL)
func Unmarshal(bs []byte, out interface{}) (err error) {
	data := string(bs)
	trySimplified := func() error {
		if strings.Count(data, "=") == 0 {
			return fmt.Errorf("Not simplifiable")
		}
		// Special case where we want to have a map and the supplied string is simplified such as "a = 10 b = string"
		// so we try transform the supplied string in valid YAML
		simplified := regexp.MustCompile(`[ \t]*=[ \t]*`).ReplaceAllString(data, ":")
		simplified = regexp.MustCompile(`[ \t]+`).ReplaceAllString(simplified, "\n")
		simplified = strings.Replace(simplified, ":", ": ", -1) + "\n"
		if err := Convert(simplified, out); err != nil {
			return err
		}
		if out, generic := out.(*interface{}); generic {
			if yaml, ok := (*out).(yamlDict); ok {
				// The YAML converter did the job, but in fact, the data looks more as Hcl
				// so we convert the result to Hcl
				*out = hclDict(yaml)
			}
		}
		return nil
	}
	var errs errors.Array

	defer func() {
		if err == nil {
			// YAML converter returns a string if it encounter invalid data, so we check the result to ensure that is is different from the input.
			if out, isItf := out.(*interface{}); isItf && data == fmt.Sprint(*out) && strings.ContainsAny(data, "=:{}") {
				if _, isString := (*out).(string); isString {
					if trySimplified() == nil && data != fmt.Sprint(*out) {
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

	for _, c := range collections.Converters() {
		err = c.Unmarshal(bs, out)
		if err == nil {
			return
		}
		errs = append(errs, fmt.Errorf("Trying %s: %v", c.Name(), err))
	}

	return errs.AsError()
}

// Convert returns a go representation of the supplied content (file, YAML, JSON or HCL)
func Convert(data string, out interface{}) error {
	if !strings.ContainsAny(data, "\t\n\":={}[]") && Load(data, out) == nil {
		return nil
	}
	return UnmarshalString(data, out)
}

// Load returns a go representation of the supplied file
func Load(uri string, out interface{}) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(uri); err == nil {
		return Unmarshal(content, out)
	}
	return
}

// ToBash returns the bash 4 variable representation of value
func ToBash(value interface{}) string {
	return toBash(ToNativeRepresentation(value), 0)
}

func toBash(value interface{}, level int) (result string) {
	if value, isString := value.(string); isString {
		result = value
		if strings.ContainsAny(value, " \t\n[]()") {
			result = fmt.Sprintf("%q", value)
		}
		return
	}

	if value, err := TryAsList(value); err == nil {
		results := value.Strings()
		for i := range results {
			results[i] = quote(results[i])
		}
		fmt.Println(results)
		switch level {
		case 2:
			result = strings.Join(results, ",")
		default:
			result = fmt.Sprintf("(%s)", strings.Join(results, " "))
		}
		return
	}

	if value, err := TryAsDictionary(value); err == nil {
		results := make([]string, value.Len())
		vMap := value.AsMap()
		switch level {
		case 0:
			for i, key := range value.KeysAsString() {
				key := key.Str()
				val := toBash(vMap[key], level+1)
				if _, err := TryAsList(vMap[key]); err == nil {
					results[i] = fmt.Sprintf("declare -a %[1]s\n%[1]s=%[2]v", key, val)
				} else if _, err := TryAsDictionary(vMap[key]); err == nil {
					results[i] = fmt.Sprintf("declare -A %[1]s\n%[1]s=%[2]v", key, val)
				} else {
					results[i] = fmt.Sprintf("%s=%v", key, val)
				}
			}
			result = strings.Join(results, "\n")
		case 1:
			for i, key := range value.KeysAsString() {
				key := key.Str()
				val := toBash(vMap[key], level+1)
				val = strings.Replace(val, `$`, `\$`, -1)
				results[i] = fmt.Sprintf("[%s]=%s", key, val)
			}
			result = fmt.Sprintf("(%s)", strings.Join(results, " "))
		default:
			for i, key := range value.KeysAsString() {
				key := key.Str()
				val := toBash(vMap[key], level+1)
				results[i] = fmt.Sprintf("%s=%s", key, quote(val))
			}
			result = strings.Join(results, ",")
		}
		return
	}
	return fmt.Sprint(value)
}

func quote(s string) string {
	if strings.ContainsAny(s, " \t,[]()") {
		s = fmt.Sprintf("%q", s)
	}
	return s
}
