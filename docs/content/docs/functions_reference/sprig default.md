---
bookToC: 2
weight: 2
---
# [Sprig Default](http://masterminds.github.io/sprig/defaults.html)
<!-- markdownlint-disable MD033 MD024 --->

## __coalesce__

```go
func coalesce(args ...interface{}) interface{}
```

<pre>
Takes a list of values and returns the first non-empty one.
</pre>

## __compact__

```go
func compact(list interface{}) []interface{}
```

<pre>
Removes entries with empty values.
</pre>

## __default__

```go
func default(default interface{}, value ...interface{}) interface{}
```

<pre>
Set a simple default value.
</pre>

## __empty__

```go
func empty(value interface{}) bool
```

<pre>
Returns true if the given value is considered empty.
</pre>

## __ternarySprig__

```go
func ternarySprig(true interface{}, false interface{}, condition bool)
interface{}
```

<pre>
If the test value is true, the first value will be returned,
otherwise, the second is returned.
</pre>

### Aliases

- _ternary_

## __toJsonSprig__

```go
func toJsonSprig(object interface{}) string
```

<pre>
Encodes an item into a JSON string.
</pre>

### Aliases

- _toJson_

## __toPrettyJsonSprig__

```go
func toPrettyJsonSprig(object interface{}) string
```

<pre>
Encodes an item into a pretty (indented) JSON string.
</pre>

### Aliases

- _toPrettyJson_

## __toRawJson__

```go
func toRawJson(object interface{}) string
```

<pre>
Encodes an item into JSON string with HTML characters unescaped.
</pre>
