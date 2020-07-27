---
bookToC: 2
weight: 2
---
# [Sprig List](http://masterminds.github.io/sprig/lists.html)
<!-- markdownlint-disable MD033 MD024 --->

## __appendSprig__

```go
func appendSprig(arg1 interface{}, arg2 interface{}) []interface{}
```

<pre>
Append a new item to an existing list, creating a new list.
</pre>

### Aliases

- _append_

## __concatSprig__

```go
func concatSprig(args ...interface{}) interface{}
```

<pre>
Concatenate arbitrary number of lists into one.
</pre>

### Aliases

- _concat_

## __first__

```go
func first(arg1 interface{}) interface{}
```

<pre>
Get the head item on a list.
</pre>

## __hasSprig__

```go
func hasSprig(arg1 interface{}, arg2 interface{}) bool
```

<pre>
Test to see if a list has a particular element.
</pre>

### Aliases

- _has_

## __initialSprig__

```go
func initialSprig(arg1 interface{}) []interface{}
```

<pre>
Complements last by returning all but the last element.
</pre>

### Aliases

- _initial_

## __last__

```go
func last(arg1 interface{}) interface{}
```

<pre>
Get the last item on a list.
</pre>

## __prependSprig__

```go
func prependSprig(arg1 interface{}, arg2 interface{}) []interface{}
```

<pre>
Push an element onto the front of a list, creating a new list.
</pre>

### Aliases

- _prepend_

## __pushSprig__

```go
func pushSprig(arg1 interface{}, arg2 interface{}) []interface{}
```

<pre>
Append a new item to an existing list, creating a new list.
</pre>

### Aliases

- _push_

## __restSprig__

```go
func restSprig(arg1 interface{}) []interface{}
```

<pre>
Get the tail of the list (everything but the first item).
</pre>

### Aliases

- _rest_

## __reverseSprig__

```go
func reverseSprig(arg1 interface{}) []interface{}
```

<pre>
Produce a new list with the reversed elements of the given list.
</pre>

### Aliases

- _reverse_

## __sliceSprig__

```go
func sliceSprig(arg1 interface{}, args ...interface{}) interface{}
```

<pre>
Get partial elements of a list.
</pre>

### Aliases

- _slice_

## __uniqSprig__

```go
func uniqSprig(arg1 interface{}) []interface{}
```

<pre>
Generate a list with all of the duplicates removed.
</pre>

### Aliases

- _uniq_

## __withoutSprig__

```go
func withoutSprig(arg1 interface{}, args ...interface{}) []interface{}
```

<pre>
Filters items out of a list.
</pre>

### Aliases

- _without_
