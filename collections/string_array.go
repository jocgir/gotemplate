package collections

// StringArray represents an array of String that complies to IGenericList interface.
type StringArray []String

// Local type used to avoid having to specify comment to all implemented function (see IGenericList for help).
type strArray = StringArray

func (sa strArray) AsArray() []interface{}                       { return AsList(sa).AsArray() }
func (sa strArray) Capacity() int                                { return cap(sa) }
func (sa strArray) Clone() IGenericList                          { return AsList(sa).Clone().Strings() }
func (sa strArray) Contains(values ...interface{}) bool          { return AsList(sa).Contains(values...) }
func (sa strArray) Count() int                                   { return len(sa) }
func (sa strArray) Create(args ...int) IGenericList              { return AsList(sa).Create(args...).Strings() }
func (sa strArray) CreateDict(args ...int) IDictionary           { return AsList(sa).CreateDict(args...) }
func (sa strArray) First() interface{}                           { return AsList(sa).First() }
func (sa strArray) Get(indexes ...int) interface{}               { return AsList(sa).Get(indexes...) }
func (sa strArray) GetHelpers() (IDictionaryHelper, IListHelper) { return AsList(sa).GetHelpers() }
func (sa strArray) Has(values ...interface{}) bool               { return AsList(sa).Has(values...) }
func (sa strArray) Join(sep IString) String                      { return AsList(sa).Join(sep) }
func (sa strArray) JoinLines() String                            { return AsList(sa).JoinLines() }
func (sa strArray) Last() interface{}                            { return AsList(sa).Last() }
func (sa strArray) New(args ...interface{}) IGenericList         { return AsList(sa).New(args...).Strings() }
func (sa strArray) Remove(indexes ...int) IGenericList           { return AsList(sa).Remove(indexes...).Strings() }
func (sa strArray) Reverse() IGenericList                        { return AsList(sa).Reverse().Strings() }
func (sa strArray) Sorted() IGenericList                         { return AsList(sa).Sorted().Strings() }
func (sa strArray) String() string                               { return AsList(sa).String() }
func (sa strArray) Strings() strArray                            { return sa }
func (sa strArray) StdStrings() []string                         { return AsList(sa).StdStrings() }
func (sa strArray) TypeName() String                             { return "StringArray" }
func (sa strArray) Unique() IGenericList                         { return AsList(sa).Unique().Strings() }

func (sa strArray) Append(values ...interface{}) IGenericList {
	return AsList(sa).Append(values).Strings()
}

func (sa strArray) Intersect(values ...interface{}) IGenericList {
	return AsList(sa).Intersect(values...).Strings()
}

func (sa strArray) Pop(indexes ...int) (interface{}, IGenericList) {
	r, l := AsList(sa).Pop(indexes...)
	return r, l.Strings()
}

func (sa strArray) Prepend(values ...interface{}) IGenericList {
	return AsList(sa).Prepend(values).Strings()
}

func (sa strArray) Set(index int, value interface{}) (IGenericList, error) {
	l, e := AsList(sa).Set(index, value)
	return l.Strings(), e
}

func (sa strArray) Union(values ...interface{}) IGenericList {
	return AsList(sa).Union(values...).Strings()
}

func (sa strArray) Without(values ...interface{}) IGenericList {
	return AsList(sa).Without(values...).Strings()
}
