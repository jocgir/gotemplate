---
bookToC: 2
weight: 2
---
# Mathematic Bit Operations
<!-- markdownlint-disable MD033 MD024 --->

## __band__

```go
func band(arg1 interface{}, arg2 interface{}, args ...interface{})
interface{}
```

<pre>
Returns the result of bitwise AND operation (&) between all arguments.
</pre>

### Aliases

- _bitwiseAND_

## __bclear__

```go
func bclear(arg1 interface{}, arg2 interface{}, args ...interface{})
interface{}
```

<pre>
Performs a bitwise clear (&^) between all arguments.
</pre>

### Aliases

- _bitwiseClear_

## __bor__

```go
func bor(arg1 interface{}, arg2 interface{}, args ...interface{})
interface{}
```

<pre>
Performs a bitwise OR (|) between all arguments.
</pre>

### Aliases

- _bitwiseOR_

## __bxor__

```go
func bxor(arg1 interface{}, arg2 interface{}, args ...interface{})
interface{}
```

<pre>
Performs a bitwise exclusive OR (^) between all arguments.
</pre>

### Aliases

- _bitwiseXOR_

## __lshift__

```go
func lshift(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Performs a left shift (&lt;&lt;) on argument.
</pre>

### Aliases

- _leftShift_

## __rshift__

```go
func rshift(arg1 interface{}, arg2 interface{}) interface{}
```

<pre>
Performs a right shift (>>) on arguments.
</pre>

### Aliases

- _rightShift_
