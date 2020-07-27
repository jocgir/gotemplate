---
bookToC: 2
weight: 2
---
# Mathematic Stats
<!-- markdownlint-disable MD033 MD024 --->

## __avg__

```go
func avg(arg1 interface{}, args ...interface{}) interface{}
```

<pre>
Returns the average value of the specified numbers.
</pre>

### Aliases

- _average_

## __max__

```go
func max(x ...interface{}) interface{}
```

<pre>
Returns the larger of x or y.
Special cases are:
    max(x, +Inf) = max(+Inf, x) = +Inf
    max(x, NaN) = max(NaN, x) = NaN
    max(+0, ±0) = max(±0, +0) = +0
    max(-0, -0) = -0
</pre>

### Aliases

- _maximum_
- _biggest_

## __min__

```go
func min(x ...interface{}) interface{}
```

<pre>
Returns the smaller of x or y.
Special cases are:
    min(x, -Inf) = min(-Inf, x) = -Inf
    min(x, NaN) = min(NaN, x) = NaN
    min(-0, ±0) = min(±0, -0) = -0
</pre>

### Aliases

- _minimum_
- _smallest_
