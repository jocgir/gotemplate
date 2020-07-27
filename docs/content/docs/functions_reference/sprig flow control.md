---
bookToC: 2
weight: 2
---
# [Sprig Flow Control](http://masterminds.github.io/sprig/flow_control.html)
<!-- markdownlint-disable MD033 MD024 --->

## __fail__

```go
func fail(arg1 string) string
```

<pre>
Unconditionally returns an empty string and an error with the
specified text. This is useful in scenarios where other conditionals
have determined that template rendering should fail.
</pre>
