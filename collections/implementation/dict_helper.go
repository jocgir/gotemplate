package implementation

import (
	"fmt"

	"github.com/coveo/gotemplate/collections"

	"github.com/imdario/mergo"
)

func (d baseDict) String() string {
	// Unlike go maps, we render dictionary keys in order
	keys := d.GetKeys()
	for i, k := range keys.AsArray() {
		k := k.(String)
		keys.Set(i, k.Format("%s:%v", d.Get(k)))
	}
	return fmt.Sprintf("dict[%s]", keys.Join(" "))
}

// DictHelper implements basic functionalities required for IDictionary.
type DictHelper struct {
	BaseHelper
}

// Clone returns a distinct copy of the object with only supplied keys. If no keys are supplied, all keys from d are copied.
func (dh DictHelper) Clone(dict baseIDict, keys []interface{}) baseIDict {
	if len(keys) == 0 {
		keys = dict.GetKeys().AsArray()
	}
	newDict := dh.CreateDictionary(dict.Count())
	for i := range keys {
		value := dict.Get(keys[i])
		if value != nil {
			if v, err := dh.TryAsDictionary(value); err == nil {
				value = dh.Clone(v, nil)
			} else if v, err := dh.TryAsList(value); err == nil {
				value = dh.ConvertList(v.Clone())
			}
		}
		newDict.Set(keys[i], value)
	}
	return newDict
}

// Default returns defVal if dictionary doesn't contain key, otherwise, simply returns entry corresponding to key.
func (dh DictHelper) Default(dict baseIDict, key, defVal interface{}) interface{} {
	if !dict.Has(key) {
		return defVal
	}
	return dict.Get(key)
}

// Delete removes the entry value associated with key. The entry must exist.
func (dh DictHelper) Delete(dict baseIDict, keys []interface{}) (baseIDict, error) {
	return dh.delete(dict, keys, true)
}

// Flush removes all specified keys from the dictionary. If no key is specified, all keys are removed.
func (dh DictHelper) Flush(dict baseIDict, keys []interface{}) baseIDict {
	if len(keys) == 0 {
		keys = dict.GetKeys().AsArray()
	}
	dh.delete(dict, keys, false)
	return dict
}

// Get returns the value associated with key.
func (dh DictHelper) Get(dict baseIDict, keys []interface{}) interface{} {
	switch len(keys) {
	case 0:
		return nil
	case 1:
		return dict.AsMap()[AsStdString(keys[0])]
	}
	result := dict.CreateList(len(keys))
	for i := range result.AsArray() {
		result.Set(i, dict.Get(keys[i]))
	}
	return result
}

// Has returns true if the dictionary object contains all the keys.
func (dh DictHelper) Has(dict baseIDict, keys []interface{}) bool {
	for _, key := range keys {
		if _, ok := dict.AsMap()[AsStdString(key)]; !ok {
			return false
		}
	}
	return dict.Count() > 0
}

// GetKeys returns the keys in the dictionary in alphabetical order.
func (dh DictHelper) GetKeys(dict baseIDict) StringArray {
	keys := make([]string, 0, dict.Count())
	for k := range dict.AsMap() {
		keys = append(keys, k)
	}
	return dh.NewStringList(keys...).Sorted().Strings()
}

// Merge merges the other dictionaries into the current dictionary.
func (dh DictHelper) Merge(target baseIDict, sources []baseIDict) baseIDict {
	m := target.AsMap()
	for i := range sources {
		if sources[i] == nil {
			continue
		}
		collections.Must(mergo.Merge(&m, dh.ConvertDict(sources[i]).AsMap()))
	}
	return target
}

// Omit returns a distinct copy of the object including all keys except specified ones.
func (dh DictHelper) Omit(dict baseIDict, keys []interface{}) baseIDict {
	omitKeys := make(map[string]bool, len(keys))
	for i := range keys {
		omitKeys[AsStdString(keys[i])] = true
	}
	keep := make([]interface{}, 0, dict.Count())
	for key := range dict.AsMap() {
		if !omitKeys[key] {
			keep = append(keep, key)
		}
	}
	return dh.Clone(dict, keep)
}

// Pop returns and remove the objects with the specified keys.
func (dh DictHelper) Pop(dict baseIDict, keys []interface{}) interface{} {
	if len(keys) == 0 {
		return nil
	}
	result := dh.Get(dict, keys)
	dh.delete(dict, keys, false)
	return result
}

// Set sets key to value in the dictionary.
func (dh DictHelper) Set(dict baseIDict, key interface{}, value interface{}) baseIDict {
	if dict.AsMap() == nil {
		dict = dh.CreateDictionary()
	}
	dict.AsMap()[AsStdString(key)] = dh.Convert(value)
	return dict
}

// Add adds value to an existing key instead of replacing the value as done by set.
func (dh DictHelper) Add(dict baseIDict, key interface{}, value interface{}) baseIDict {
	if dict.AsMap() == nil {
		dict = dh.CreateDictionary()
	}
	m := dict.AsMap()
	k := AsStdString(key)

	if current, ok := m[k]; ok {
		if list, err := dh.TryAsList(current); err == nil {
			m[k] = list.Append(value)
		} else {
			// Convert the current value into a list
			m[k] = dict.CreateList().Append(current, value)
		}
	} else {
		m[k] = dh.Convert(value)
	}
	return dict
}

// GetValues returns the values in the dictionary in key alphabetical order.
func (dh DictHelper) GetValues(dict baseIDict) baseIList {
	result := dh.CreateList(dict.Count())
	for i, key := range dict.GetKeys() {
		result.Set(i, dict.Get(key))
	}
	return result
}

// Transpose returns a dictionary with values as keys and keys as values. The resulting dictionary
// is a dictionary where each key could contains single value or list of values if there are multiple matches.
func (dh DictHelper) Transpose(dict baseIDict) baseIDict {
	result := dh.CreateDictionary()
	for _, key := range dict.GetKeys() {
		value := dict.Get(key)
		if list, err := dh.TryAsList(value); err == nil {
			// If the element is a list, we scan each element
			for _, value := range list.AsArray() {
				result.Add(value, key)
			}
		} else {
			result.Add(value, key)
		}
	}
	return result
}

func (dh DictHelper) delete(dict baseIDict, keys []interface{}, mustExist bool) (baseIDict, error) {
	for i := range keys {
		if mustExist && !dict.Has(keys[i]) {
			return dict, fmt.Errorf("key %v not found", keys[i])
		}
		delete(dict.AsMap(), AsStdString(keys[i]))
	}
	return dict, nil
}
