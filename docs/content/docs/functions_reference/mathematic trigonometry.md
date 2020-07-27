---
bookToC: 2
weight: 2
---
# Mathematic Trigonometry
<!-- markdownlint-disable MD033 MD024 --->

## __acos__

```go
func acos(x interface{}) interface{}
```

<pre>
Returns the arccosine, in radians, of x.
Special case is:
    acos(x) = NaN if x &lt; -1 or x > 1
</pre>

### Aliases

- _arcCosine_
- _arcCosinus_

### Examples

```go
Razor:    @ceil(acos(0.5) / 3.1416 * 180)
Template: {{ ceil (mul (div (acos 0.5) 3.1416) 180) }}
Result:   60
```

## __acosh__

```go
func acosh(x interface{}) interface{}
```

<pre>
Returns the inverse hyperbolic cosine of x.
Special cases are:
    acosh(+Inf) = +Inf
    acosh(x) = NaN if x &lt; 1
    acosh(NaN) = NaN
</pre>

### Aliases

- _arcHyperbolicCosine_
- _arcHyperbolicCosinus_

## __asin__

```go
func asin(x interface{}) interface{}
```

<pre>
Returns the arcsine, in radians, of x.
Special cases are:
    asin(±0) = ±0
    asin(x) = NaN if x &lt; -1 or x > 1
</pre>

### Aliases

- _arcSine_
- _arcSinus_

## __asinh__

```go
func asinh(x interface{}) interface{}
```

<pre>
Returns the inverse hyperbolic sine of x.
Special cases are:
    asinh(±0) = ±0
    asinh(±Inf) = ±Inf
    asinh(NaN) = NaN
</pre>

### Aliases

- _arcHyperbolicSine_
- _arcHyperbolicSinus_

## __atan__

```go
func atan(x interface{}) interface{}
```

<pre>
Returns the arctangent, in radians, of x.
Special cases are:
    atan(±0) = ±0
    atan(±Inf) = ±Pi/2
</pre>

### Aliases

- _arcTangent_

## __atan2__

```go
func atan2(x interface{}, y interface{}) interface{}
```

<pre>
Returns the arc tangent of y/x, using the signs of the two to
determine the quadrant of the return value.
Special cases are (in order):
    atan2(y, NaN) = NaN
    atan2(NaN, x) = NaN
    atan2(+0, x>=0) = +0
    atan2(-0, x>=0) = -0
    atan2(+0, x&lt;=-0) = +Pi
    atan2(-0, x&lt;=-0) = -Pi
    atan2(y>0, 0) = +Pi/2
    atan2(y&lt;0, 0) = -Pi/2
    atan2(+Inf, +Inf) = +Pi/4
    atan2(-Inf, +Inf) = -Pi/4
    atan2(+Inf, -Inf) = 3Pi/4
    atan2(-Inf, -Inf) = -3Pi/4
    atan2(y, +Inf) = 0
    atan2(y>0, -Inf) = +Pi
    atan2(y&lt;0, -Inf) = -Pi
    atan2(+Inf, x) = +Pi/2
    atan2(-Inf, x) = -Pi/2
</pre>

### Aliases

- _arcTangent2_

## __atanh__

```go
func atanh(x interface{}) interface{}
```

<pre>
Returns the inverse hyperbolic tangent of x.
Special cases are:
    atanh(1) = +Inf
    atanh(±0) = ±0
    atanh(-1) = -Inf
    atanh(x) = NaN if x &lt; -1 or x > 1
    atanh(NaN) = NaN
</pre>

### Aliases

- _arcHyperbolicTangent_

## __cos__

```go
func cos(x interface{}) interface{}
```

<pre>
Returns the cosine of the radian argument x.
Special cases are:
    cos(±Inf) = NaN
    cos(NaN) = NaN
</pre>

### Aliases

- _cosine_
- _cosinus_

## __cosh__

```go
func cosh(x interface{}) interface{}
```

<pre>
Returns the hyperbolic cosine of x.
Special cases are:
    cosh(±0) = 1
    cosh(±Inf) = +Inf
    cosh(NaN) = NaN
</pre>

### Aliases

- _hyperbolicCosine_
- _hyperbolicCosinus_

## __deg__

```go
func deg(arg1 interface{}) interface{}
```

<pre>
Returns the decimal degree of the specified radian number.
</pre>

### Aliases

- _degree_

## __ilogb__

```go
func ilogb(x interface{}) interface{}
```

<pre>
Returns the binary exponent of x as an integer.
Special cases are:
    ilogb(±Inf) = MaxInt32
    ilogb(0) = MinInt32
    ilogb(NaN) = MaxInt32
</pre>

## __j0__

```go
func j0(x interface{}) interface{}
```

<pre>
Returns the order-zero Bessel function of the first kind.
Special cases are:
    j0(±Inf) = 0
    j0(0) = 1
    j0(NaN) = NaN
</pre>

### Aliases

- _firstBessel0_

## __j1__

```go
func j1(x interface{}) interface{}
```

<pre>
Returns the order-one Bessel function of the first kind.
Special cases are:
    j1(±Inf) = 0
    j1(NaN) = NaN
</pre>

### Aliases

- _firstBessel1_

## __jn__

```go
func jn(n interface{}, x interface{}) interface{}
```

<pre>
Returns the order-n Bessel function of the first kind.
Special cases are:
    jn(n, ±Inf) = 0
    jn(n, NaN) = NaN
</pre>

### Aliases

- _firstBesselN_

## __log__

```go
func log(x interface{}) interface{}
```

<pre>
Returns the natural logarithm of x.
Special cases are:
    log(+Inf) = +Inf
    log(0) = -Inf
    log(x &lt; 0) = NaN
    log(NaN) = NaN
</pre>

## __log10__

```go
func log10(x interface{}) interface{}
```

<pre>
Returns the decimal logarithm of x. The special cases are the same as
for log.
</pre>

## __log1p__

```go
func log1p(x interface{}) interface{}
```

<pre>
Returns the natural logarithm of 1 plus its argument x. It is more
accurate than log(1 + x) when x is near zero.
Special cases are:
    log1p(+Inf) = +Inf
    log1p(±0) = ±0
    log1p(-1) = -Inf
    log1p(x &lt; -1) = NaN
    log1p(NaN) = NaN
</pre>

## __log2__

```go
func log2(x interface{}) interface{}
```

<pre>
Returns the binary logarithm of x. The special cases are the same as
for log.
</pre>

## __logb__

```go
func logb(x interface{}) interface{}
```

<pre>
Returns the binary exponent of x.
Special cases are:
    logb(±Inf) = +Inf
    logb(0) = -Inf
    logb(NaN) = NaN
</pre>

## __rad__

```go
func rad(arg1 interface{}) interface{}
```

<pre>
Returns the radian of the specified decimal degree number.
</pre>

### Aliases

- _radian_

## __sin__

```go
func sin(x interface{}) interface{}
```

<pre>
Returns the sine of the radian argument x.
Special cases are:
    sin(±0) = ±0
    sin(±Inf) = NaN
    sin(NaN) = NaN
</pre>

### Aliases

- _sine_
- _sinus_

## __sincos__

```go
func sincos(x interface{}) interface{}
```

<pre>
Returns Sin(x), Cos(x).
Special cases are:
    sincos(±0) = ±0, 1
    sincos(±Inf) = NaN, NaN
    sincos(NaN) = NaN, NaN
</pre>

### Aliases

- _sineCosine_
- _sinusCosinus_

## __sinh__

```go
func sinh(x interface{}) interface{}
```

<pre>
Returns the hyperbolic sine of x.
Special cases are:
    sinh(±0) = ±0
    sinh(±Inf) = ±Inf
    sinh(NaN) = NaN
</pre>

### Aliases

- _hyperbolicSine_
- _hyperbolicSinus_

## __tan__

```go
func tan(x interface{}) interface{}
```

<pre>
Returns the tangent of the radian argument x.
Special cases are:
    tan(±0) = ±0
    tan(±Inf) = NaN
    tan(NaN) = NaN
</pre>

### Aliases

- _tangent_

## __tanh__

```go
func tanh(x interface{}) interface{}
```

<pre>
Returns the hyperbolic tangent of x.
Special cases are:
    tanh(±0) = ±0
    tanh(±Inf) = ±1
    tanh(NaN) = NaN
</pre>

### Aliases

- _hyperbolicTangent_

## __y0__

```go
func y0(x interface{}) interface{}
```

<pre>
Returns the order-zero Bessel function of the second kind.
Special cases are:
    y0(+Inf) = 0
    y0(0) = -Inf
    y0(x &lt; 0) = NaN
    y0(NaN) = NaN
</pre>

### Aliases

- _secondBessel0_

## __y1__

```go
func y1(x interface{}) interface{}
```

<pre>
Returns the order-one Bessel function of the second kind.
Special cases are:
    y1(+Inf) = 0
    y1(0) = -Inf
    y1(x &lt; 0) = NaN
    y1(NaN) = NaN
</pre>

### Aliases

- _secondBessel1_

## __yn__

```go
func yn(n interface{}, x interface{}) interface{}
```

<pre>
Returns the order-n Bessel function of the second kind.
Special cases are:
    yn(n, +Inf) = 0
    yn(n ≥ 0, 0) = -Inf
    yn(n &lt; 0, 0) = +Inf if n is odd, -Inf if n is even
    yn(n, x &lt; 0) = NaN
    yn(n, NaN) = NaN
</pre>

### Aliases

- _secondBesselN_
