package collections

import (
	"sort"
	"sync"
)

// Register provides a mechanism to register a data converter
func Register(name string, priority int, handler func([]byte, interface{}) error) error {
	convertersMutex.Lock()
	defer convertersMutex.Unlock()
	converters = append(converters, TypeConverter{name, priority, handler})
	convertersSorted = false
	return nil
}

// Converters returns the list of currently registered converters
func Converters() []TypeConverter {
	if !convertersSorted {
		sort.SliceStable(converters, func(i, j int) bool {
			ci, cj := converters[i], converters[j]
			return ci.priority < cj.priority || ci.priority == cj.priority && ci.name < cj.name
		})
		convertersSorted = true
	}
	return converters
}

var (
	converters       []TypeConverter
	convertersSorted bool
	convertersMutex  sync.Mutex
)

// TypeConverter is used to register the available converters
type TypeConverter struct {
	name      string
	priority  int
	Unmarshal func([]byte, interface{}) error
}

type typeConverter = TypeConverter

func (tc *typeConverter) Name() string  { return tc.name }
func (tc *typeConverter) Priority() int { return tc.priority }
