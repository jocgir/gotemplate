package collections

import (
	"fmt"
	"sync"
)

// IGenericList represents objects that act as []interface{}.
type IGenericList interface {
	Append(...interface{}) IGenericList                     // Add elements (with transformation) to the current list. If list is not large enough, it is enlarged to fit the required size.
	AppendRaw(...interface{}) IGenericList                  // Add raw elements to the current list. If list is not large enough, it is enlarged to fit the required size.
	AsArray() []interface{}                                 // Returns the current list as standard array of interface{}.
	Cap() int                                               // Returns the capacity of the list.
	Capacity() int                                          // Simply an alias for Cap.
	Clone() IGenericList                                    // Returns a distinct copy of the object.
	Contains(...interface{}) bool                           // Indicates if the list contains all specified elements (matches even if type are not the same).
	ContainsStrict(...interface{}) bool                     // Indicates if the list contains all specified elements (matches only for same types).
	Count() int                                             // Simply an alias for Len.
	Create(...int) IGenericList                             // Allocates a new list of the same type implementation as this list. Optional arguments are size and capacity.
	CreateDict(...int) IDictionary                          // Instantiates a new dictionary of the same type with optional size.
	Find(interface{}) IGenericList                          // Return the index positions of the searched element in a list (match elements even if type is not the same).
	FindStrict(interface{}) IGenericList                    // Return the index positions of the searched element in a list (must have the same type).
	First() interface{}                                     // Returns the first element of the list.
	Get(...int) interface{}                                 // Returns the element at position index in the list. If index is out of bound, nil is returned.
	GetHelpers() (IDictionaryHelper, IListHelper)           // Returns the helpers implementation associated with the current type.
	Has(...interface{}) bool                                // Alias for Contains
	HasStrict(...interface{}) bool                          // Alias for ContainsStrict
	Intersect(...interface{}) IGenericList                  // Returns a list that is the result of the intersection of the list and the parameters (removing duplicates).
	Join(sep interface{}) String                            // Returns the string representation of the list.
	Last() interface{}                                      // Returns the last element of the list.
	Len() int                                               // Returns the number of elements in the list.
	New(...interface{}) IGenericList                        // Creates a new generic list from the supplied arguments.
	Pop(indexes ...int) (interface{}, IGenericList)         // Removes and returns the elements of the list (if nothing is specified, remove the last element).
	Prepend(...interface{}) IGenericList                    // Add elements (with transformation) to the beginning of the current list. If list is not large enough, it is enlarged to fit the required size.
	PrependRaw(...interface{}) IGenericList                 // Add raw elements to the beginning of the current list. If list is not large enough, it is enlarged to fit the required size.
	PrettyPrint() string                                    // Returns the pretty string representation of the list.
	Remove(indexes ...int) IGenericList                     // Returns a new list without the element specified.
	RemoveEmpty() IGenericList                              // Returns a new list without the element that evaluates to empty or zero.
	RemoveNil() IGenericList                                // Returns a new list without the nil elements.
	Reverse() IGenericList                                  // Returns a copy of the current list in reverse order.
	Set(index int, value interface{}) (IGenericList, error) // Sets the value at position index into the list. If list is not large enough, it is enlarged to fit the index.
	String() string                                         // Returns the string representation of the list.
	StringArray() StringArray                               // Returns the current list as StringArray.
	Strings() []string                                      // Returns the current list as list of strings.
	TypeName() String                                       // Returns the actual type name
	Union(...interface{}) IGenericList                      // Returns a list that represents the union of the list and the elements (removing duplicates).
	Unique() IGenericList                                   // Returns a copy of the list removing all duplicate elements.
	Without(...interface{}) IGenericList                    // Returns a copy of the list removing specified elements.
}

// IListHelper represents objects that implement IGenericList compatible objects
type IListHelper interface {
	AsList(interface{}) IGenericList                    // Converts object to IGenericList object. It panics if conversion is impossible.
	Convert(object interface{}) interface{}             // Tries to convert the supplied object into IDictionary or IGenericList.
	CreateList(...int) IGenericList                     // Creates a new IGenericList with optional size/capacity arguments.
	NewList(...interface{}) IGenericList                // Creates a new IGenericList from supplied arguments.
	NewStringList(...string) IGenericList               // Creates a new IGenericList from supplied arguments.
	TryAsList(object interface{}) (IGenericList, error) // Tries to convert any object to IGenericList object.
	TryConvert(object interface{}) (interface{}, bool)  // Tries to convert any object to IGenericList or IDictionary object.
}

var listHelper IListHelper
var listHelperMutex sync.Mutex

// GetListHelper fetches the default list manager
func GetListHelper() IListHelper {
	listHelperMutex.Lock()
	defer listHelperMutex.Unlock()
	return listHelper
}

// SetListHelper configures the default list manager
func SetListHelper(newListHelper IListHelper) {
	listHelperMutex.Lock()
	defer listHelperMutex.Unlock()
	listHelper = newListHelper
}

func assertListHelper() {
	if GetListHelper() == nil {
		panic(fmt.Errorf("ListHelper not configured"))
	}
}

// AsList returns the object casted as IGenericList.
func AsList(object interface{}) IGenericList {
	result, err := TryAsList(object)
	if err != nil {
		return NewList(object)
	}
	return result
}

// CreateList instantiates a new generic list with optional size and capacity.
func CreateList(args ...int) IGenericList {
	assertListHelper()
	return GetListHelper().CreateList(args...)
}

// NewList instantiates a new generic list from supplied arguments.
func NewList(objects ...interface{}) IGenericList {
	assertListHelper()
	return GetListHelper().NewList(objects...)
}

// NewStringList creates a new IGenericList from supplied arguments.
func NewStringList(objects ...string) IGenericList {
	assertListHelper()
	return GetListHelper().NewStringList(objects...)
}

// TryAsList returns the object casted as IGenericList if possible.
func TryAsList(object interface{}) (IGenericList, error) {
	if result, ok := object.(IGenericList); ok {
		return result, nil
	}
	assertListHelper()
	return GetListHelper().TryAsList(object)
}
