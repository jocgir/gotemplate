---
bookToC: 2
weight: 2
---
# [Sprig Version comparison](http://masterminds.github.io/sprig/semver.html)
<!-- markdownlint-disable MD033 MD024 --->

## __semver__

```go
func semver(version string) *semver.Version
```

<pre>
Parses a string into a Semantic Version.
</pre>

## __semverCompare__

```go
func semverCompare(constraints string, version string) bool
```

<pre>
A more robust comparison function is provided as semverCompare. This
version supports version ranges.
</pre>
