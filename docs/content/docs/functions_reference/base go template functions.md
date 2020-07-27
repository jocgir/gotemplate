---
bookToC: 2
weight: 2
---
# Base go template functions
<!-- markdownlint-disable MD033 MD024 --->

## __and__

```go
func and(arg0 reflect.Value, args ...reflect.Value) reflect.Value
```

<pre>
Returns the boolean AND of its arguments by returning the first empty
argument or the last argument, that is, "and x y" behaves as "if x
then y else x". All the arguments are evaluated.
</pre>

## __call__

```go
func call(fn reflect.Value, args ...reflect.Value) reflect.Value
```

<pre>
Returns the result of calling the first argument, which must be a
function, with the remaining arguments as parameters. Thus "call .X.Y
1 2" is, in Go notation, dot.X.Y(1, 2) where Y is a func-valued field,
map entry, or the like. The first argument must be the result of an
evaluation that yields a value of function type (as distinct from a
predefined function such as print). The function must return either
one or two result values, the second of which is of type error. If the
arguments don't match the function or the returned error value is
non-nil, execution stops.
</pre>

## __eq__

```go
func eq(arg1 reflect.Value, arg2 ...reflect.Value) bool
```

<pre>
Returns the boolean truth of arg1 == arg2
</pre>

## __ge__

```go
func ge(arg1 reflect.Value, arg2 ...reflect.Value) bool
```

<pre>
Returns the boolean truth of arg1 >= arg2
</pre>

## __gt__

```go
func gt(arg1 reflect.Value, arg2 ...reflect.Value) bool
```

<pre>
Returns the boolean truth of arg1 > arg2
</pre>

## __html__

```go
func html(args ...interface{}) string
```

<pre>
Returns the escaped HTML equivalent of the textual representation of
its arguments. This function is unavailable in html/template, with a
few exceptions.
</pre>

## __index__

```go
func index(item reflect.Value, indices ...reflect.Value) reflect.Value
```

<pre>
Returns the result of indexing its first argument by the following
arguments. Thus "index x 1 2 3" is, in Go syntax, x[1][2][3]. Each
indexed item must be a map, slice, or array.
</pre>

## __js__

```go
func js(args ...interface{}) string
```

<pre>
Returns the escaped JavaScript equivalent of the textual
representation of its arguments.
</pre>

## __le__

```go
func le(arg1 reflect.Value, arg2 ...reflect.Value) bool
```

<pre>
Returns the boolean truth of arg1 &lt;= arg2
</pre>

## __len__

```go
func len(item interface{}) int
```

<pre>
Returns the integer length of its argument.
</pre>

## __lt__

```go
func lt(arg1 reflect.Value, arg2 ...reflect.Value) bool
```

<pre>
Returns the boolean truth of arg1 &lt; arg2
</pre>

## __ne__

```go
func ne(arg1 reflect.Value, arg2 ...reflect.Value) bool
```

<pre>
Returns the boolean truth of arg1 != arg2
</pre>

## __not__

```go
func not(not(arg reflect.Value) bool
```

<pre>
Returns the boolean negation of its single argument.
</pre>

## __or__

```go
func or(or(arg0 reflect.Value, args ...reflect.Value) reflect.Value
```

<pre>
Returns the boolean OR of its arguments by returning the first
non-empty argument or the last argument, that is, "or x y" behaves as
"if x then x else y". All the arguments are evaluated.
</pre>

## __print__

```go
func print(args ...interface{}) string
```

<pre>
An alias for fmt.Sprint
</pre>

## __printf__

```go
func printf(format string, args ...interface{}) string
```

<pre>
An alias for fmt.Sprintf
</pre>

## __println__

```go
func println(args ...interface{}) string
```

<pre>
An alias for fmt.Sprintln
</pre>

## __urlquery__

```go
func urlquery(args ...interface{}) string
```

<pre>
Returns the escaped value of the textual representation of its
arguments in a form suitable for embedding in a URL query. This
function is unavailable in html/template, with a few exceptions.
</pre>
