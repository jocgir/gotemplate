---
bookToC: 2
weight: 2
---
# [Sprig Strings](http://masterminds.github.io/sprig/strings.html)
<!-- markdownlint-disable MD033 MD024 --->

## __abbrev__

```go
func abbrev(width int, str string) string
```

<pre>
Truncates a string with ellipses (...).
</pre>

## __abbrevboth__

```go
func abbrevboth(left int, right int, str string) string
```

<pre>
Abbreviates both sides with ellipses (...).
</pre>

## __camelcase__

```go
func camelcase(str string) string
```

<pre>
Converts string from snake_case to CamelCase.
</pre>

## __cat__

```go
func cat(args ...interface{}) string
```

<pre>
Concatenates multiple strings together into one, separating them with
spaces.
</pre>

## __containsSprig__

```go
func containsSprig(substr string, str string) bool
```

<pre>
Tests to see if one string is contained inside of another.
</pre>

### Aliases

- _contains_

## __hasPrefix__

```go
func hasPrefix(prefix string, str string) bool
```

<pre>
Tests whether a string has a given prefix.
</pre>

## __hasSuffix__

```go
func hasSuffix(suffix string, str string) bool
```

<pre>
Tests whether a string has a given suffix.
</pre>

## __indentSprig__

```go
func indentSprig(spaces int, str string) string
```

<pre>
Indents every line in a given string to the specified indent width.
This is useful when aligning multi-line strings.
</pre>

### Aliases

- _indent_

## __initials__

```go
func initials(str string) string
```

<pre>
Given multiple words, takes the first letter of each word and combine.
</pre>

## __kebabcase__

```go
func kebabcase(str string) string
```

<pre>
Convert string from camelCase to kebab-case.
</pre>

## __lower__

```go
func lower(str string) string
```

<pre>
Converts the entire string to lowercase.
</pre>

## __nindentSprig__

```go
func nindentSprig(spaces int, str string) string
```

<pre>
Same as the indent function, but prepends a new line to the beginning
of the string.
</pre>

### Aliases

- _nindent_

## __nospace__

```go
func nospace(str string) string
```

<pre>
Removes all whitespace from a string.
</pre>

## __plural__

```go
func plural(one string, many string, count int) string
```

<pre>
Pluralizes a string.
</pre>

## __quote__

```go
func quote(str ...interface{}) string
```

<pre>
Wraps each argument with double quotes.
</pre>

## __randAlpha__

```go
func randAlpha(count int) string
```

<pre>
Generates random string with letters.
</pre>

## __randAlphaNum__

```go
func randAlphaNum(count int) string
```

<pre>
Generates random string with letters and digits.
</pre>

## __randAscii__

```go
func randAscii(count int) string
```

<pre>
Generates random string with ASCII printable characters.
</pre>

## __randNumeric__

```go
func randNumeric(count int) string
```

<pre>
Generates random string with digits.
</pre>

## __repeatSprig__

```go
func repeatSprig(count int, str string) string
```

<pre>
Repeats a string multiple times.
</pre>

### Aliases

- _repeat_

## __replace__

```go
func replace(old string, new string, src string) string
```

<pre>
Performs simple string replacement.
</pre>

## __shuffle__

```go
func shuffle(str string) string
```

<pre>
Shuffle a string.
</pre>

## __snakecase__

```go
func snakecase(str string) string
```

<pre>
Converts string from camelCase to snake_case.
</pre>

## __squote__

```go
func squote(args ...interface{}) string
```

<pre>
Wraps each argument with single quotes.
</pre>

## __substr__

```go
func substr(start int, length int, str string) string
```

<pre>
Get a substring from a string.
</pre>

## __swapcase__

```go
func swapcase(str string) string
```

<pre>
Swaps the uppercase to lowercase and lowercase to uppercase.
</pre>

## __title__

```go
func title(str string) string
```

<pre>
Converts to title case.
</pre>

## __toString__

```go
func toString(value interface{}) string
```

<pre>
Converts any value to string.
</pre>

## __trim__

```go
func trim(str string) string
```

<pre>
Removes space from either side of a string.
</pre>

## __trimAll__

```go
func trimAll(chars string, str string) string
```

<pre>
Removes given characters from the front or back of a string.
</pre>

### Aliases

- _trimall_

## __trimPrefix__

```go
func trimPrefix(prefix string, str string) string
```

<pre>
Trims just the prefix from a string if present.
</pre>

## __trimSuffix__

```go
func trimSuffix(suffix string, str string) string
```

<pre>
Trims just the suffix from a string if present.
</pre>

## __truncSprig__

```go
func truncSprig(length int, str string) string
```

<pre>
Truncates a string (and add no suffix).
</pre>

### Aliases

- _trunc_

## __untitle__

```go
func untitle(str string) string
```

<pre>
Removes title casing.
</pre>

## __upper__

```go
func upper(str string) string
```

<pre>
Converts the entire string to uppercase.
</pre>

## __wrapSprig__

```go
func wrapSprig(length int, str string) string
```

<pre>
Wraps text at a given column count.
</pre>

### Aliases

- _wrap_

## __wrapWith__

```go
func wrapWith(length int, spe string, str string) string
```

<pre>
Works as wrap, but lets you specify the string to wrap with (wrap uses
\n).
</pre>
