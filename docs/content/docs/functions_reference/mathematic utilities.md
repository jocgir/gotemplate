---
bookToC: 2
weight: 2
---
# Mathematic Utilities
<!-- markdownlint-disable MD033 MD024 --->

## __abs__

```go
func abs(x interface{}) interface{}
```

<pre>
Returns the absolute value of x.
Special cases are:
    abs(±Inf) = +Inf
    abs(NaN) = NaN
</pre>

### Aliases

- _absolute_

### Examples

```go
Razor:    @abs(0)
Template: {{ abs 0 }}
Result:   0
```

```go
Razor:    @abs(22)
Template: {{ abs 22 }}
Result:   22
```

```go
Razor:    @abs(-10)
Template: {{ abs -10 }}
Result:   10
```

## __dec__

```go
func dec(arg1 interface{}) interface{}
```

<pre>
Converts an hexadecimal number to decimal.
</pre>

### Aliases

- _decimal_

## __frexp__

```go
func frexp(f interface{}) interface{}
```

<pre>
Breaks f into a normalized fraction and an integral power of two.
Returns frac and exp satisfying f == frac × 2**exp, with the absolute
value of frac in the interval [½, 1).
Special cases are:
    frexp(±0) = ±0, 0
    frexp(±Inf) = ±Inf, 0
    frexp(NaN) = NaN, 0
</pre>

## __gamma__

```go
func gamma(x interface{}) interface{}
```

<pre>
Returns the Gamma function of x.
Special cases are:
    gamma(+Inf) = +Inf
    gamma(+0) = +Inf
    gamma(-0) = -Inf
    gamma(x) = NaN for integer x &lt; 0
    gamma(-Inf) = NaN
    gamma(NaN) = NaN
</pre>

## __hex__

```go
func hex(arg1 interface{}) interface{}
```

<pre>
Formats a number as hexadecimal number.
</pre>

### Aliases

- _hexa_
- _hexaDecimal_

## __hypot__

```go
func hypot(p interface{}, q interface{}) interface{}
```

<pre>
Returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and
underflow.
Special cases are:
    hypot(±Inf, q) = +Inf
    hypot(p, ±Inf) = +Inf
    hypot(NaN, q) = NaN
    hypot(p, NaN) = NaN
</pre>

### Aliases

- _hypotenuse_

## __isInf__

```go
func isInf(f interface{}, arg2 interface{}) interface{}
```

<pre>
Reports whether f is an infinity, according to sign. If sign > 0,
isInf reports whether f is positive infinity. If sign &lt; 0, IsInf
reports whether f is negative infinity. If sign == 0, IsInf reports
whether f is either infinity
</pre>

### Aliases

- _isInfinity_

## __isNaN__

```go
func isNaN(f interface{}) interface{}
```

<pre>
Reports whether f is an IEEE 754 'not-a-number' value
</pre>

## __ldexp__

```go
func ldexp(frac interface{}, exp interface{}) interface{}
```

<pre>
Ldexp is the inverse of Frexp. Returns frac × 2**exp.
Special cases are:
    ldexp(±0, exp) = ±0
    ldexp(±Inf, exp) = ±Inf
    ldexp(NaN, exp) = NaN
</pre>

## __lgamma__

```go
func lgamma(x interface{}) interface{}
```

<pre>
Returns the natural logarithm and sign (-1 or +1) of Gamma(x).
Special cases are:
    lgamma(+Inf) = +Inf
    lgamma(0) = +Inf
    lgamma(-integer) = +Inf
    lgamma(-Inf) = -Inf
    lgamma(NaN) = NaN
</pre>

## __nextAfter__

```go
func nextAfter(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Returns the next representable float64 value after x towards y.
Special cases are:
    Nextafter(x, x) = x
Nextafter(NaN, y) = NaN
Nextafter(x, NaN) = NaN
</pre>

## __signBit__

```go
func signBit(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Reports whether x is negative or negative zero.
</pre>

## __sqrt__

```go
func sqrt(x interface{}) interface{}
```

<pre>
Returns the square root of x.
Special cases are:
    sqrt(+Inf) = +Inf
    sqrt(±0) = ±0
    sqrt(x &lt; 0) = NaN
    sqrt(NaN) = NaN
</pre>

### Aliases

- _squareRoot_

## __to__

```go
func to(args ...interface{}) interface{}
```

<pre>
Builds a range of integers starting with 1 by default and including
the upper limit.
</pre>

### Examples

```go
Razor:    @to(10)
Template: {{ to 10 }}
Result:   [1,2,3,4,5,6,7,8,9,10]
```

```go
Razor:    @to(10, 0)
Template: {{ to 10 0 }}
Result:   [10,9,8,7,6,5,4,3,2,1,0]
```

```go
Razor:    @to(0, 10, 2)
Template: {{ to 0 10 2 }}
Result:   [0,2,4,6,8,10]
```

## __until__

```go
func until(args ...interface{}) interface{}
```

<pre>
Builds a range of integers starting with 0 by default and not
including the upper limit.
</pre>

### Examples

```go
Razor:    @until(10)
Template: {{ until 10 }}
Result:   [0,1,2,3,4,5,6,7,8,9]
```

```go
Razor:    @until(10, 0)
Template: {{ until 10 0 }}
Result:   [10,9,8,7,6,5,4,3,2,1]
```

```go
Razor:    @until(0, 10, 2)
Template: {{ until 0 10 2 }}
Result:   [0,2,4,6,8]
```
