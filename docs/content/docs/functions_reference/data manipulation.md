---
bookToC: 2
weight: 2
---
# Data Manipulation
<!-- markdownlint-disable MD033 MD024 --->

## __String__

```go
func String(value interface{}) String
```

<pre>
Returns a String class object that allows invoking standard string
operations as method.
</pre>

## __append__

```go
func append(list interface{}, elements ...interface{}) IGenericList
```

<pre>
Append new items to an existing list, creating a new list.
</pre>

### Aliases

- _push_

## __array__

```go
func array(value interface{}) interface{}
```

<pre>
Ensures that the supplied argument is an array (if it is already an
array/slice, there is no change, if not, the argument is replaced by
[]interface{} with a single value).
</pre>

## __bool__

```go
func bool(str string) bool
```

<pre>
Converts the `string` into boolean value (`string` must be `True`,
`true`, `TRUE`, `1` or `False`, `false`, `FALSE`, `0`)
</pre>

## __char__

```go
func char(value interface{}) interface{}
```

<pre>
Returns the character corresponging to the supplied integer value
</pre>

## __contains__

```go
func contains(list interface{}, elements ...interface{}) bool
```

<pre>
Tests whether a list contains all given elements (matches any types).
</pre>

### Aliases

- _has_

## __containsStrict__

```go
func containsStrict(list interface{}, elements ...interface{}) bool
```

<pre>
Tests whether a list contains all given elements (matches only the
same types).
</pre>

### Aliases

- _hasStrict_

## __content__

```go
func content(keymap interface{}) interface{}
```

<pre>
Returns the content of a single element map.
Used to retrieve content in a declaration like:
    value "name" { a = 1 b = 3 }
</pre>

## __dict__

```go
func dict(args ...interface{}) IDictionary
```

<pre>
Returns a new dictionary from a list of pairs (key, value).
</pre>

### Aliases

- _dictionary_

## __extract__

```go
func extract(source interface{}, indexes ...interface{}) interface{}
```

<pre>
Extracts values from a slice or a map, indexes could be either
integers for slice or strings for maps.
</pre>

## __find__

```go
func find(list interface{}, element interface{}) interface{}
```

<pre>
Returns all index positions where the element is found in the list
(matches any types).
</pre>

## __findStrict__

```go
func findStrict(list interface{}, element interface{}) interface{}
```

<pre>
Returns all index positions where the element is found in the list
(matches only the same types).
</pre>

## __get__

```go
func get(map interface{}, key interface{}, default ...interface{})
interface{}
```

<pre>
Returns the value associated with the supplied map, key and map could
be inverted for convenience (i.e. when using piping mode).
</pre>

## __hasKey__

```go
func hasKey(dictionary interface{}, key interface{}) interface{}
```

<pre>
Returns true if the dictionary contains the specified key.
</pre>

### Examples

```go
Razor:    @hasKey(dict("key", "value"), "key")
Template: {{ hasKey (dict "key" "value") "key" }}
Result:   true
```

```go
Razor:    @hasKey("key", dict("key", "value"))
Template: {{ hasKey "key" (dict "key" "value") }}
Result:   true
```

```go
Razor:    @hasKey(dict("key", "value"), "otherkey")
Template: {{ hasKey (dict "key" "value") "otherkey" }}
Result:   false
```

## __initial__

```go
func initial(list interface{}) interface{}
```

<pre>
Returns but the last element.
</pre>

## __intersect__

```go
func intersect(list interface{}, elements ...interface{}) IGenericList
```

<pre>
Returns a list that is the intersection of the list and all arguments
(removing duplicates).
</pre>

## __isNil__

```go
func isNil(arg1 interface{}) bool
```

<pre>
Returns true if the supplied value is nil.
</pre>

### Aliases

- _isNull_

## __isSet__

```go
func isSet(arg1 interface{}) bool
```

<pre>
Returns true if the supplied value is not nil.
</pre>

## __isZero__

```go
func isZero(arg1 interface{}) bool
```

<pre>
Returns true if the supplied value is false, 0, nil or empty.
</pre>

### Aliases

- _isEmpty_

## __key__

```go
func key(value interface{}) interface{}
```

<pre>
Returns the key name of a single element map.
Used to retrieve name in a declaration like:
    value "name" { a = 1 b = 3 }
</pre>

## __keys__

```go
func keys(dictionary IDictionary) IGenericList
```

<pre>
Returns a list of all of the keys in a dict (in alphabetical order).
</pre>

## __lenc__

```go
func lenc(str string) int
```

<pre>
Returns the number of actual character in a string.
</pre>

### Aliases

- _nbChars_

## __list__

```go
func list(args ...interface{}) IGenericList
```

<pre>
Returns a generic list from the supplied arguments.
</pre>

### Aliases

- _tuple_

## __merge__

```go
func merge(destination IDictionary, sources IDictionary, args
...IDictionary) IDictionary
```

<pre>
Merges two or more dictionaries into one, giving precedence to the
dest dictionary.
</pre>

## __omit__

```go
func omit(dict IDictionary, keys interface{}, args ...interface{})
IDictionary
```

<pre>
Returns a new dict with all the keys that do not match the given keys.
</pre>

## __pick__

```go
func pick(dict IDictionary, keys ...interface{}) IDictionary
```

<pre>
Selects just the given keys out of a dictionary, creating a new dict.
</pre>

## __pickv__

```go
func pickv(dict IDictionary, message string, keys interface{}, args
...interface{}) interface{}
```

<pre>
Same as pick, but returns an error message if there are intruders in
supplied dictionary.
</pre>

## __pluck__

```go
func pluck(key interface{}, dictionaries ...IDictionary) IGenericList
```

<pre>
Extracts a list of values matching the supplied key from a list of
dictionary.
</pre>

## __prepend__

```go
func prepend(list interface{}, elements ...interface{}) IGenericList
```

<pre>
Push elements onto the front of a list, creating a new list.
</pre>

## __removeEmpty__

```go
func removeEmpty(list interface{}) IGenericList
```

<pre>
Returns a list with all empty elements removed.
</pre>

## __removeNil__

```go
func removeNil(list interface{}) IGenericList
```

<pre>
Returns a list with all nil elements removed.
</pre>

## __rest__

```go
func rest(list interface{}) interface{}
```

<pre>
Gets the tail of the list (everything but the first item)
</pre>

## __reverse__

```go
func reverse(list interface{}) IGenericList
```

<pre>
Produces a new list with the reversed elements of the given list.
</pre>

## __safeIndex__

```go
func safeIndex(value interface{}, index int, default interface{})
interface{}
```

<pre>
Returns the element at index position or default if index is outside
bounds.
</pre>

## __set__

```go
func set(dict interface{}, key interface{}, value interface{}) string
```

<pre>
Adds the value to the supplied map using key as identifier.
</pre>

## __slice__

```go
func slice(value interface{}, args ...interface{}) interface{}
```

<pre>
Returns a slice of the supplied object (equivalent to
object[from:to]).
</pre>

## __string__

```go
func string(value interface{}) string
```

<pre>
Converts the supplied value into its string representation.
</pre>

## __undef__

```go
func undef(default interface{}, values ...interface{}) interface{}
```

<pre>
Returns the default value if value is not set, alias `undef` (differs
from Sprig `default` function as empty value such as 0, false, "" are
not considered as unset).
</pre>

### Aliases

- _ifUndef_

## __union__

```go
func union(list interface{}, elements ...interface{}) IGenericList
```

<pre>
Returns a list that is the union of the list and all arguments
(removing duplicates).
</pre>

## __unique__

```go
func unique(list interface{}) IGenericList
```

<pre>
Generates a list with all of the duplicates removed.
</pre>

### Aliases

- _uniq_

## __unset__

```go
func unset(dictionary interface{}, key interface{}) string
```

<pre>
Removes an element from a dictionary.
</pre>

### Aliases

- _delete_
- _remove_

### Examples

```go
Razor:    @{myDict} := dict("key", "value", "key2", "value2", "key3", "value3")
          @-unset($myDict, "key")
          @-unset("key2", $myDict)
          @-toJson($myDict)
Template: {{- $myDict := dict "key" "value" "key2" "value2" "key3" "value3" }}
          {{- unset $myDict "key" }}
          {{- unset "key2" $myDict }}
          {{- toJson $myDict }}
Result:   {"key3":"value3"}
```

## __values__

```go
func values(arg1 IDictionary) IGenericList
```

<pre>
Returns the list of values contained in a map.
</pre>

## __without__

```go
func without(list interface{}, elements ...interface{}) IGenericList
```

<pre>
Filters items out of a list.
</pre>
