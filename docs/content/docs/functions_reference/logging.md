---
bookToC: 2
weight: 2
---
# Logging
<!-- markdownlint-disable MD033 MD024 --->

## __critical__

```go
func critical(args ...interface{}) string
```

<pre>
Deprecated: Use error instead. Logs a message using ERROR log level
(2).
</pre>

### Aliases

- _criticalf_

## __debug__

```go
func debug(args ...interface{}) string
```

<pre>
Logs a message using DEBUG as log level (5).
</pre>

### Aliases

- _debugf_

## __error__

```go
func error(args ...interface{}) string
```

<pre>
Logs a message using ERROR as log level (2).
</pre>

### Aliases

- _errorf_

## __fatal__

```go
func fatal(args ...interface{}) string
```

<pre>
Logs a message using FATAL as log level (1) followed by a call to
os.Exit(1).
</pre>

### Aliases

- _fatalf_

## __info__

```go
func info(args ...interface{}) string
```

<pre>
Logs a message using INFO as log level (4).
</pre>

### Aliases

- _infof_

## __notice__

```go
func notice(args ...interface{}) string
```

<pre>
Deprecated: Use info instead. Logs a message using INFO log level (4).
</pre>

### Aliases

- _noticef_

## __panic__

```go
func panic(args ...interface{}) string
```

<pre>
Logs a message using PANIC as log level (0) followed by a call to
panic.
</pre>

### Aliases

- _panicf_

## __trace__

```go
func trace(args ...interface{}) string
```

<pre>
Logs a message using TRACE as log level (6).
</pre>

### Aliases

- _tracef_

## __warning__

```go
func warning(args ...interface{}) string
```

<pre>
Logs a message using WARNING as log level (3).
</pre>

### Aliases

- _warn_
- _warnf_
- _warningf_
