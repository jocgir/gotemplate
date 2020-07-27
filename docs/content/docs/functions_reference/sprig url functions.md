---
bookToC: 2
weight: 2
---
# [Sprig URL functions](http://masterminds.github.io/sprig/url.html)
<!-- markdownlint-disable MD033 MD024 --->

## __urlJoin__

```go
func urlJoin(dictionary map[string]interface{}) string
```

<pre>
Joins map (produced by `urlParse`) to produce URL string.
</pre>

## __urlParse__

```go
func urlParse(uri string) map[string]interface{}
```

<pre>
Parses string for URL and produces dict with URL parts.
</pre>
