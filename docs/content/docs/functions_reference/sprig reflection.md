---
bookToC: 2
weight: 2
---
# [Sprig Reflection](http://masterminds.github.io/sprig/reflection.html)
<!-- markdownlint-disable MD033 MD024 --->

## __deepEqual__

```go
func deepEqual(arg1 interface{}, arg2 interface{}) bool
```

<pre>
returns true if two values are deeply equal.
</pre>

## __kindIs__

```go
func kindIs(arg1 string, arg2 interface{}) bool
```

<pre>
Let you verify that a value is a particular kind.
</pre>

### Aliases

- _kindis_

## __kindOf__

```go
func kindOf(arg1 interface{}) string
```

<pre>
Returns the kind of an object.
</pre>

### Aliases

- _kindof_

## __typeIs__

```go
func typeIs(arg1 string, arg2 interface{}) bool
```

<pre>
Like `kindIs`, but for types.
</pre>

### Aliases

- _typeis_

## __typeIsLike__

```go
func typeIsLike(arg1 string, arg2 interface{}) bool
```

<pre>
Works as `typeIs`, except that it also dereferences pointers.
</pre>

### Aliases

- _typeisLike_

## __typeOf__

```go
func typeOf(arg1 interface{}) string
```

<pre>
Returns the underlying type of a value.
</pre>

### Aliases

- _typeof_
