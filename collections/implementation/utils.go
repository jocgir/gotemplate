package implementation

import "github.com/coveo/gotemplate/collections"

// Register the default implementation of dictionary and list
var _ = func() int {
	collections.DictionaryHelper, collections.ListHelper = baseList{}.GetHelpers()
	return 0
}()
