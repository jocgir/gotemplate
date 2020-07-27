---
bookToC: 2
weight: 2
---
# [Sprig Regex](http://masterminds.github.io/sprig/strings.html)
<!-- markdownlint-disable MD033 MD024 --->

## __regexFind__

```go
func regexFind(regex string, str string) string
```

<pre>
Returns the first (left most) match of the regular expression in the
input string.
</pre>

## __regexFindAll__

```go
func regexFindAll(regex string, str string, n int) []string
```

<pre>
Returns a slice of all matches of the regular expression in the input
string.
</pre>

## __regexMatch__

```go
func regexMatch(regex string, str string) bool
```

<pre>
Returns true if the input string matches the regular expression.
</pre>

## __regexReplaceAll__

```go
func regexReplaceAll(regex string, str string, repl string) string
```

<pre>
Returns a copy of the input string, replacing matches of the Regexp
with the replacement string replacement. Inside string replacement, $
signs are interpreted as in Expand, so for instance $1 represents the
text of the first submatch.
</pre>

## __regexReplaceAllLiteral__

```go
func regexReplaceAllLiteral(regex string, str string, repl string) string
```

<pre>
Returns a copy of the input string, replacing matches of the Regexp
with the replacement string replacement The replacement string is
substituted directly, without using Expand.
</pre>

## __regexSplit__

```go
func regexSplit(regex string, str string, n int) []string
```

<pre>
Slices the input string into substrings separated by the expression
and returns a slice of the substrings between those expression
matches. The last parameter n determines the number of substrings to
return, where -1 means return all matches.
</pre>
