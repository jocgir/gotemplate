package implementation

import (
	"fmt"

	"github.com/coveooss/gotemplate/v3/collections"
)

func (d baseDict) String() string {
	// Unlike go maps, we render dictionary keys in order
	keys := d.KeysAsString()
	for i, k := range keys {
		keys[i] = String(fmt.Sprintf("%s:%v", k, d.Get(k)))
	}
	return fmt.Sprintf("dict[%s]", keys.Join(" "))
}

func (d baseDict) PrettyPrint() string { return d.String() }

// DictHelper implements basic functionalities required for IDictionary.
type DictHelper struct {
	BaseHelper
}

// AsDictionary returns the object casted as IDictionary.
func (dh DictHelper) AsDictionary(object interface{}) IDictionary {
	return must(dh.TryAsDictionary(object)).(IDictionary)
}

// Clone returns a distinct copy of the object with only supplied keys. If no keys are supplied, all keys from d are copied.
func (dh DictHelper) Clone(dict IDictionary, keys []interface{}) IDictionary {
	if len(keys) == 0 {
		keys = dict.GetKeys().AsArray()
	}
	newDict := dh.CreateDictionary(dict.Len())
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
func (dh DictHelper) Default(dict IDictionary, key, defVal interface{}) interface{} {
	if !dict.Has(key) {
		return defVal
	}
	return dict.Get(key)
}

// Delete removes the entry value associated with key. The entry must exist.
func (dh DictHelper) Delete(dict IDictionary, first interface{}, rest []interface{}) (IDictionary, error) {
	return dh.delete(dict, append([]interface{}{first}, rest...), true)
}

// Flush removes all specified keys from the dictionary. If no key is specified, all keys are removed.
func (dh DictHelper) Flush(dict IDictionary, keys []interface{}) IDictionary {
	if len(keys) == 0 {
		keys = dict.GetKeys().AsArray()
	}
	dh.delete(dict, keys, false)
	return dict
}

// Get returns the value associated with key.
func (dh DictHelper) Get(dict IDictionary, keys []interface{}) interface{} {
	switch len(keys) {
	case 0:
		return nil
	case 1:
		return dict.AsMap()[fmt.Sprint(keys[0])]
	}
	result := dict.CreateList(len(keys))
	for i := range result.AsArray() {
		result.Set(i, dict.Get(keys[i]))
	}
	return result
}

// Has returns true if the dictionary object contains all the keys.
func (dh DictHelper) Has(dict IDictionary, keys []interface{}) bool {
	for _, key := range keys {
		if _, ok := dict.AsMap()[fmt.Sprint(key)]; !ok {
			return false
		}
	}
	return dict.Len() > 0
}

// GetKeys returns the keys in the dictionary in alphabetical order.
func (dh DictHelper) GetKeys(dict IDictionary) IGenericList {
	keys := dict.KeysAsString()
	result := dh.CreateList(dict.Len())

	for i := range keys {
		result.Set(i, keys[i])
	}
	return result
}

// KeysAsString returns the keys in the dictionary in alphabetical order.
func (dh DictHelper) KeysAsString(dict IDictionary) StringArray {
	keys := make(StringArray, 0, dict.Len())
	for key := range dict.AsMap() {
		keys = append(keys, String(key))
	}
	return keys.Sorted()
}

// Merge merges the other dictionaries into the current dictionary.
func (dh DictHelper) Merge(target IDictionary, first IDictionary, rest []IDictionary) IDictionary {
	sources := append([]IDictionary{first}, rest...)
	for i := range sources {
		if sources[i] == nil {
			continue
		}
		target = dh.deepMerge(target, dh.ConvertDict(sources[i]))
	}
	return target
}

func (dh DictHelper) deepMerge(target IDictionary, source IDictionary) IDictionary {
	targetMap := target.AsMap()
	sourceMap := source.AsMap()
	for key := range sourceMap {
		sourceValue, sourceHasKey := sourceMap[key]
		targetValue, targetHasKey := targetMap[key]

		if sourceHasKey && !targetHasKey {
			targetMap[key] = sourceValue
			continue
		}

		targetValueDict, targetValueIsDict := targetValue.(IDictionary)
		sourceValueDict, sourceValueIsDict := sourceValue.(IDictionary)

		if sourceValueIsDict && targetValueIsDict {
			targetMap[key] = dh.deepMerge(targetValueDict, sourceValueDict)
		}
	}

	return target
}

// Omit returns a distinct copy of the object including all keys except specified ones.
func (dh DictHelper) Omit(dict IDictionary, first interface{}, rest []interface{}) IDictionary {
	keys := append([]interface{}{first}, rest...)
	omitKeys := make(map[string]bool, len(keys))
	for i := range keys {
		omitKeys[fmt.Sprint(keys[i])] = true
	}
	keep := make([]interface{}, 0, dict.Len())
	for key := range dict.AsMap() {
		if !omitKeys[key] {
			keep = append(keep, key)
		}
	}
	return dh.Clone(dict, keep)
}

// Pop returns and remove the objects with the specified keys.
func (dh DictHelper) Pop(dict IDictionary, keys []interface{}) interface{} {
	if len(keys) == 0 {
		return nil
	}
	result := dh.Get(dict, keys)
	dh.delete(dict, keys, false)
	return result
}

// Set sets key to value in the dictionary.
func (dh DictHelper) Set(dict IDictionary, key interface{}, value interface{}) IDictionary {
	if dict.AsMap() == nil {
		dict = dh.CreateDictionary()
	}
	dict.AsMap()[fmt.Sprint(key)] = dh.Convert(value)
	return dict
}

// Add adds value to an existing key instead of replacing the value as done by set.
func (dh DictHelper) Add(dict IDictionary, key interface{}, value interface{}) IDictionary {
	if dict.AsMap() == nil {
		dict = dh.CreateDictionary()
	}
	m := dict.AsMap()
	k := fmt.Sprint(key)

	if current, ok := m[k]; ok {
		if list, err := collections.TryAsList(current); err == nil {
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
func (dh DictHelper) GetValues(dict IDictionary) IGenericList {
	result := dh.CreateList(dict.Len())
	for i, key := range dict.KeysAsString() {
		result.Set(i, dict.Get(key))
	}
	return result
}

// Transpose returns a dictionary with values as keys and keys as values. The resulting dictionary
// is a dictionary where each key could contains single value or list of values if there are multiple matches.
func (dh DictHelper) Transpose(dict IDictionary) IDictionary {
	result := dh.CreateDictionary()
	for _, key := range dict.GetKeys().AsArray() {
		value := dict.Get(key)
		if list, err := collections.TryAsList(value); err == nil {
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

func (dh DictHelper) delete(dict IDictionary, keys []interface{}, mustExist bool) (IDictionary, error) {
	for i := range keys {
		if mustExist && !dict.Has(keys[i]) {
			return dict, fmt.Errorf("key %v not found", keys[i])
		}
		delete(dict.AsMap(), fmt.Sprint(keys[i]))
	}
	return dict, nil
}

// Register the default implementation of dictionary helper
var _ = func() int {
	collections.DictionaryHelper = dh
	return 0
}()
