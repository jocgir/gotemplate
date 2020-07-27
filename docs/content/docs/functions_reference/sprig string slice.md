---
bookToC: 2
weight: 2
---
# [Sprig String Slice](http://masterminds.github.io/sprig/string_slice.html)
<!-- markdownlint-disable MD033 MD024 --->

## __join__

```go
func join(separator string, list interface{}) string
```

<pre>
Joins a list of strings into a single string, with the given
separator.
</pre>

## __sortAlpha__

```go
func sortAlpha(list interface{}) []string
```

<pre>
Sorts a list of strings into alphabetical (lexicographical) order.
</pre>

## __split__

```go
func split(separator string, str string) map[string]string
```

<pre>
Splits a string into a `dict`. It is designed to make it easy to use
template dot notation for accessing members
</pre>

## __splitList__

```go
func splitList(separator string, str string) []string
```

<pre>
Splits a string into a list of strings.
</pre>

## __splitn__

```go
func splitn(separator string, count int, str string) map[string]string
```

<pre>
Splits a string into a `dict`. It is designed to make it easy to use
template dot notation for accessing members.
</pre>

## __toStrings__

```go
func toStrings(list interface{}) []string
```

<pre>
Given a list-like collection, produce a slice of strings.
</pre>
