---
bookToC: 2
weight: 2
---
# [Sprig Mathematics](http://masterminds.github.io/sprig/math.html)
<!-- markdownlint-disable MD033 MD024 --->

## __add1__

```go
func add1(arg1 interface{}) int64
```

<pre>
Increments a value by 1
</pre>

## __addSprig__

```go
func addSprig(args ...interface{}) int64
```

<pre>
Sum numbers with `add`. Accepts two or more inputs.
</pre>

### Aliases

- _add_

## __ceilSprig__

```go
func ceilSprig(arg1 interface{}) float64
```

<pre>
Returns the greatest float value greater than or equal to input value.
</pre>

### Aliases

- _ceil_

## __divSprig__

```go
func divSprig(arg1 interface{}, arg2 interface{}) int64
```

<pre>
Performs integer division.
</pre>

### Aliases

- _div_

## __floorSprig__

```go
func floorSprig(arg1 interface{}) float64
```

<pre>
Returns the greatest float value less than or equal to input value
</pre>

### Aliases

- _floor_

## __maxSprig__

```go
func maxSprig(arg1 interface{}, args ...interface{}) int64
```

<pre>
Returns the largest of a series of integers.
</pre>

### Aliases

- _max_
- _biggest_
- _biggestSprig_

## __minSprig__

```go
func minSprig(arg1 interface{}, args ...interface{}) int64
```

<pre>
Returns the smallest of a series of integers.
</pre>

### Aliases

- _min_

## __modSprig__

```go
func modSprig(arg1 interface{}, arg2 interface{}) int64
```

<pre>
Performs integer modulo.
</pre>

### Aliases

- _mod_

## __mulSprig__

```go
func mulSprig(arg1 interface{}, args ...interface{}) int64
```

<pre>
Multiplies numbers. Accepts two or more inputs.
</pre>

### Aliases

- _mul_

## __round__

```go
func round(arg1 interface{}, arg2 int, args ...float64) float64
```

<pre>
Returns a float value with the remainder rounded to the given number
to digits after the decimal point.
</pre>

## __seq__

```go
func seq(args ...int) string
```

<pre>
Works like the bash seq command.
</pre>

## __subSprig__

```go
func subSprig(arg1 interface{}, arg2 interface{}) int64
```

<pre>
Subtracts a number from another number.`
</pre>

### Aliases

- _sub_

## __untilStep__

```go
func untilStep(arg1 int, arg2 int, arg3 int) []int
```

<pre>
Generates a list of counting integers. But it allows you to define a
start, stop, and step
</pre>
