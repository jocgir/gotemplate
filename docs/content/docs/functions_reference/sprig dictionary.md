---
bookToC: 2
weight: 2
---
# [Sprig Dictionary](http://masterminds.github.io/sprig/dicts.html)
<!-- markdownlint-disable MD033 MD024 --->

## __deepCopy__

```go
func deepCopy(arg1 interface{}) interface{}
```

<pre>
Takes a value and makes a deep copy of the value.
</pre>

## __dictSprig__

```go
func dictSprig(args ...interface{}) map[string]interface{}
```

<pre>
Create a dictionary.
</pre>

### Aliases

- _dict_

## __getSprig__

```go
func getSprig(dict map[string]interface{}, key string) interface{}
```

<pre>
Given a map and a key, get the value from the map.
</pre>

### Aliases

- _get_

## __hasKeySprig__

```go
func hasKeySprig(dict map[string]interface{}, key string) bool
```

<pre>
Returns 'true' if the given dict contains the given key.
</pre>

### Aliases

- _hasKey_

## __keysSprig__

```go
func keysSprig(args ...map[string]interface{}) []string
```

<pre>
Returns a list of all of the keys in one or more dict types.
</pre>

### Aliases

- _keys_

## __listSprig__

```go
func listSprig(args ...interface{}) []interface{}
```

<pre>
Create a list of elements.
</pre>

### Aliases

- _list_
- _tuple_
- _tupleSprig_

## __mergeOverwrite__

```go
func mergeOverwrite(arg1 map[string]interface{}, args
...map[string]interface{}) interface{}
```

<pre>
Merge two or more dictionaries into one, giving precedence from
**right to left**, effectively overwriting values in the dest
dictionary.
</pre>

## __mergeSprig__

```go
func mergeSprig(arg1 map[string]interface{}, args
...map[string]interface{}) interface{}
```

<pre>
Merge two or more dictionaries into one, giving precedence to the dest
dictionary.
</pre>

### Aliases

- _merge_

## __omitSprig__

```go
func omitSprig(arg1 map[string]interface{}, args ...string)
map[string]interface{}
```

<pre>
Is similar to 'pick', except it returns a new `dict` with all the keys
that _do not_ match the given keys.
</pre>

### Aliases

- _omit_

## __pickSprig__

```go
func pickSprig(arg1 map[string]interface{}, args ...string)
map[string]interface{}
```

<pre>
Selects just the given keys out of a dictionary, creating a new
`dict`.
</pre>

### Aliases

- _pick_

## __pluckSprig__

```go
func pluckSprig(name string, dicts ...map[string]interface{}) []interface{}
```

<pre>
Makes it possible to give one key and multiple maps, and get a list of
all of the matches.
</pre>

### Aliases

- _pluck_

## __setSprig__

```go
func setSprig(dict map[string]interface{}, key string, value interface{})
map[string]interface{}
```

<pre>
Add a new key/value pair to a dictionary.
</pre>

### Aliases

- _set_

## __unsetSprig__

```go
func unsetSprig(dict map[string]interface{}, key string)
map[string]interface{}
```

<pre>
Given a map and a key, delete the key from the map.
</pre>

### Aliases

- _unset_

## __valuesSprig__

```go
func valuesSprig(arg1 map[string]interface{}) []interface{}
```

<pre>
Is similar to 'keys', except it returns a new 'list' with all the
values of the source 'dict' (only one dictionary is supported).
</pre>

### Aliases

- _values_
