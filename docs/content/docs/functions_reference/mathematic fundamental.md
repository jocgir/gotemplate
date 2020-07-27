---
bookToC: 2
weight: 2
---
# Mathematic Fundamental
<!-- markdownlint-disable MD033 MD024 --->

## __add__

```go
func add(arg1 interface{}, args ...interface{}) interface{}
```

<pre>
Returns the result of the addition of all arguments.
</pre>

### Aliases

- _sum_

## __cbrt__

```go
func cbrt(x interface{}) interface{}
```

<pre>
Returns the cube root of x.
Special cases are:
    cbrt(±0) = ±0
    cbrt(±Inf) = ±Inf
    cbrt(NaN) = NaN
</pre>

## __ceil__

```go
func ceil(x interface{}) interface{}
```

<pre>
Returns the least integer value greater than or equal to x.
Special cases are:
    ceil(±0) = ±0
    ceil(±Inf) = ±Inf
    ceil(NaN) = NaN
</pre>

### Aliases

- _roundUp_
- _roundup_

## __dim__

```go
func dim(x interface{}, y interface{}) interface{}
```

<pre>
Returns the maximum of x-y or 0.
Special cases are:
    dim(+Inf, +Inf) = NaN
    dim(-Inf, -Inf) = NaN
    dim(x, NaN) = dim(NaN, x) = NaN
</pre>

## __div__

```go
func div(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Returns the result of the division of the two arguments.
</pre>

### Aliases

- _divide_
- _quotient_

## __exp__

```go
func exp(x interface{}) interface{}
```

<pre>
Returns e**x, the base-e exponential of x.
Special cases are:
    exp(+Inf) = +Inf
    exp(NaN) = NaN
Very large values overflow to 0 or +Inf. Very small values underflow
to 1.
</pre>

### Aliases

- _exponent_

## __exp2__

```go
func exp2(x interface{}) interface{}
```

<pre>
Returns 2**x, the base-2 exponential of x.
Special cases are the same as exp.
</pre>

### Aliases

- _exponent2_

## __expm1__

```go
func expm1(x interface{}) interface{}
```

<pre>
Returns e**x - 1, the base-e exponential of x minus 1. It is more
accurate than exp(x) - 1 when x is near zero.
Special cases are:
    expm1(+Inf) = +Inf
    expm1(-Inf) = -1
    expm1(NaN) = NaN
Very large values overflow to -1 or +Inf
</pre>

## __floor__

```go
func floor(x interface{}) interface{}
```

<pre>
Returns the greatest integer value less than or equal to x.
Special cases are:
    floor(±0) = ±0
    floor(±Inf) = ±Inf
    floor(NaN) = NaN
</pre>

### Aliases

- _roundDown_
- _rounddown_
- _int64_
- _integer64_

## __int__

```go
func int(arg1 interface{}) int
```

<pre>
Returns the integer value (type int).
</pre>

### Aliases

- _integer_

## __mod__

```go
func mod(x interface{}, y interface{}) interface{}
```

<pre>
Returns the floating-point remainder of x/y. The magnitude of the
result is less than y and its sign agrees with that of x.
Special cases are:
    mod(±Inf, y) = NaN
    mod(NaN, y) = NaN
    mod(x, 0) = NaN
    mod(x, ±Inf) = x
    mod(x, NaN) = NaN
</pre>

### Aliases

- _modulo_

## __modf__

```go
func modf(f interface{}) interface{}
```

<pre>
Returns integer and fractional floating-point numbers that sum to f.
Both values have the same sign as f.
Special cases are:
    modf(±Inf) = ±Inf, NaN
    modf(NaN) = NaN, NaN
</pre>

## __mul__

```go
func mul(arg1 interface{}, args ...interface{}) interface{}
```

<pre>
Returns the result of the multiplication of all arguments.
</pre>

### Aliases

- _multiply_
- _prod_
- _product_

## __pow__

```go
func pow(x interface{}, y interface{}) interface{}
```

<pre>
Returns x**y, the base-x exponential of y.
Special cases are (in order):
    pow(x, ±0) = 1 for any x
    pow(1, y) = 1 for any y
    pow(x, 1) = x for any x
    pow(NaN, y) = NaN
    pow(x, NaN) = NaN
    pow(±0, y) = ±Inf for y an odd integer &lt; 0
    pow(±0, -Inf) = +Inf
    pow(±0, +Inf) = +0
    pow(±0, y) = +Inf for finite y &lt; 0 and not an odd integer
    pow(±0, y) = ±0 for y an odd integer > 0
    pow(±0, y) = +0 for finite y > 0 and not an odd integer
    pow(-1, ±Inf) = 1
    pow(x, +Inf) = +Inf for |x| > 1
    pow(x, -Inf) = +0 for |x| > 1
    pow(x, +Inf) = +0 for |x| &lt; 1
    pow(x, -Inf) = +Inf for |x| &lt; 1
    pow(+Inf, y) = +Inf for y > 0
    pow(+Inf, y) = +0 for y &lt; 0
    pow(-Inf, y) = Pow(-0, -y)
    pow(x, y) = NaN for finite x &lt; 0 and finite non-integer y
</pre>

### Aliases

- _power_

## __pow10__

```go
func pow10(n interface{}) interface{}
```

<pre>
Returns 10**n, the base-10 exponential of n.
Special cases are:
    pow10(n) =0 for n &lt; -323
    pow10(n) = +Inf for n > 308
</pre>

### Aliases

- _power10_

## __rem__

```go
func rem(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Returns the IEEE 754 floating-point remainder of x/y.
Special cases are:
    rem(±Inf, y) = NaN
    rem(NaN, y) = NaN
    rem(x, 0) = NaN
    rem(x, ±Inf) = x
    rem(x, NaN) = NaN
</pre>

### Aliases

- _remainder_

## __sub__

```go
func sub(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Returns the result of the substraction of the two arguments.
</pre>

### Aliases

- _subtract_

## __trunc__

```go
func trunc(x interface{}) interface{}
```

<pre>
Returns the integer value of x.
Special cases are:
    trunc(±0) = ±0
    trunc(±Inf) = ±Inf
    trunc(NaN) = NaN
</pre>

### Aliases

- _truncate_
