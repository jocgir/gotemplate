package implementation

import (
	"github.com/coveooss/gotemplate/v3/collections"
	"github.com/coveooss/gotemplate/v3/errors"
	"github.com/coveooss/gotemplate/v3/stringclass"
)

// SetAsDefault configure the current implementation as default list & dictionary manager.
func SetAsDefault() { collections.ListHelper, collections.DictionaryHelper = lh, dh }

type (
	itf = interface{}

	// List is the base implementation of IGenericList
	List     = baseList
	baseList []itf
	il       = IGenericList
	bl       = baseList

	// Dictionary is the base implementation of IDictionary
	Dictionary = baseDict
	baseDict   map[string]itf
	id         = IDictionary
	bd         = baseDict
)

// List implementation
func (l bl) Append(values ...itf) il                { return lh.Add(l, false, values...) }
func (l bl) AsArray() []itf                         { return []itf(l) }
func (l bl) Cap() int                               { return cap(l) }
func (l bl) Capacity() int                          { return cap(l) }
func (l bl) Clone() il                              { return lh.Clone(l) }
func (l bl) Contains(values ...itf) bool            { return lh.Contains(l, values...) }
func (l bl) Count() int                             { return len(l) }
func (l bl) Create(args ...int) il                  { return lh.CreateList(args...) }
func (l bl) CreateDict(args ...int) id              { return lh.CreateDictionary(args...) }
func (l bl) First() itf                             { return lh.GetIndexes(l, 0) }
func (l bl) Get(indexes ...int) itf                 { return lh.GetIndexes(l, indexes...) }
func (l bl) GetHelpers() (IDictHelper, IListHelper) { return dh, lh }
func (l bl) Has(values ...itf) bool                 { return l.Contains(values...) }
func (l bl) Intersect(values ...itf) il             { return lh.Intersect(l, values...) }
func (l bl) Join(sep itf) String                    { return l.StringArray().Join(sep) }
func (l bl) Last() itf                              { return lh.GetIndexes(l, len(l)-1) }
func (l bl) Len() int                               { return len(l) }
func (l bl) New(args ...itf) il                     { return lh.NewList(args...) }
func (l bl) Pop(indexes ...int) (itf, il)           { return lh.Pop(l, indexes...) }
func (l bl) Prepend(values ...itf) il               { return lh.Add(l, true, values...) }
func (l bl) Remove(indexes ...int) il               { return lh.Remove(l, indexes...) }
func (l bl) Reverse() il                            { return lh.Reverse(l) }
func (l bl) Set(i int, v itf) (il, error)           { return lh.SetIndex(l, i, v) }
func (l bl) StringArray() StringArray               { return lh.GetStringArray(l) }
func (l bl) Strings() []string                      { return lh.GetStrings(l) }
func (l bl) TypeName() String                       { return "base" }
func (l bl) Union(values ...itf) il                 { return lh.Add(l, false, values...).Unique() }
func (l bl) Unique() il                             { return lh.Unique(l) }
func (l bl) Without(values ...itf) il               { return lh.Without(l, values...) }

// Dictionary implementation
func (d bd) Add(key, v itf) id                         { return dh.Add(d, key, v) }
func (d bd) AsMap() map[string]itf                     { return (map[string]itf)(d) }
func (d bd) Clone(keys ...itf) id                      { return dh.Clone(d, keys) }
func (d bd) Count() int                                { return len(d) }
func (d bd) Create(args ...int) id                     { return lh.CreateDictionary(args...) }
func (d bd) CreateList(args ...int) il                 { return dh.CreateList(args...) }
func (d bd) Default(key, defVal itf) itf               { return dh.Default(d, key, defVal) }
func (d bd) Delete(first itf, rest ...itf) (id, error) { return dh.Delete(d, first, rest) }
func (d bd) Flush(keys ...itf) id                      { return dh.Flush(d, keys) }
func (d bd) Get(keys ...itf) itf                       { return dh.Get(d, keys) }
func (d bd) GetHelpers() (IDictHelper, IListHelper)    { return dh, lh }
func (d bd) GetKeys() il                               { return dh.GetKeys(d) }
func (d bd) GetValues() il                             { return dh.GetValues(d) }
func (d bd) Has(keys ...itf) bool                      { return dh.Has(d, keys) }
func (d bd) KeysAsString() StringArray                 { return dh.KeysAsString(d) }
func (d bd) Len() int                                  { return len(d) }
func (d bd) Merge(first id, rest ...id) id             { return dh.Merge(d, first, rest) }
func (d bd) Native() itf                               { return must(collections.MarshalGo(d)) }
func (d bd) Omit(first itf, rest ...itf) id            { return dh.Omit(d, first, rest) }
func (d bd) Pop(keys ...itf) itf                       { return dh.Pop(d, keys) }
func (d bd) Set(key, v itf) id                         { return dh.Set(d, key, v) }
func (d bd) Transpose() id                             { return dh.Transpose(d) }
func (d bd) TypeName() String                          { return "base" }

// Generic helpers to simplify physical implementation
var (
	helper = helperBase{
		ConvertList:    func(list il) il { return baseList(list.AsArray()) },
		ConvertDict:    func(dict id) id { return baseDict(dict.AsMap()) },
		NeedConversion: func(object itf, strict bool) bool { return needConversionImpl(object, strict, "base") },
	}
	lh = helperList{BaseHelper: helper}
	dh = helperDict{BaseHelper: helper}
)

// Imported types
type (
	// IGenericList is imported from collections
	IGenericList = collections.IGenericList

	// IListHelper is imported from collections
	IListHelper = collections.IListHelper

	// IDictionary is imported from collections
	IDictionary = collections.IDictionary

	// IDictHelper is imported from collections
	IDictHelper = collections.IDictionaryHelper

	// String is imported from stringclass
	String = stringclass.String

	// StringArray is imported from stringclass
	StringArray = stringclass.StringArray
)

// Imported functions
var (
	iif           = collections.IIf
	must          = errors.Must
	TrimmedString = stringclass.TrimmedString
)
