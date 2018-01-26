package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// LoadYaml returns the content of a YAML file as go object
func LoadYaml(filename string) (result map[string]interface{}, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(filename); err == nil {
		err = YamlUnmarshal(content, &result)
	}
	return
}

// MapKeyInterface2string convert maps with interface{} key to map with a string as the key
func MapKeyInterface2string(source interface{}) interface{} {
	switch value := source.(type) {
	case map[string]interface{}:
		for key, val := range value {
			value[key] = MapKeyInterface2string(val)
		}
		return value
	case map[interface{}]interface{}:
		result := make(map[string]interface{}, len(value))
		for key, val := range value {
			result[fmt.Sprintf("%v", key)] = MapKeyInterface2string(val)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(value), len(value))
		for i, val := range value {
			result[i] = MapKeyInterface2string(val)
		}
		return result
	}
	return source
}

// ToYaml returns a yaml representation of the supplied object
func ToYaml(v interface{}) (string, error) {
	result, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// YamlUnmarshal calls yaml.Unmarshal, but replace tabs by spaces if there are
func YamlUnmarshal(in []byte, out interface{}) (err error) {
	// Yaml does not support tab, so we repace tabs by spaces if there are
	in = []byte(strings.Replace(string(in), "\t", "    ", -1))
	return yaml.Unmarshal(in, out)
}
