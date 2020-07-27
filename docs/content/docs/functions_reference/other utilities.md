---
bookToC: 2
weight: 2
---
# Other utilities
<!-- markdownlint-disable MD033 MD024 --->

## __center__

```go
func center(width interface{}, args ...interface{}) string
```

<pre>
Returns the concatenation of supplied arguments centered within width.
</pre>

### Aliases

- _centered_

## __color__

```go
func color(args ...interface{}) string
```

<pre>
Colors the rendered string.

The first arguments are interpretated as color attributes until the
first non color attribute. Attributes are case insensitive.

Valid attributes are:
Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid,
    ReverseVideo, Concealed, CrossedOut

Valid color are:
    Black, Red, Green, Yellow, Blue, Magenta, Cyan, White

Color can be prefixed by:
    Fg: Meaning foreground (Fg is assumed if not specified)
    FgHi: Meaning high intensity forgeround
    Bg: Meaning background"
    BgHi: Meaning high intensity background
</pre>

### Aliases

- _colored_
- _enhanced_

## __colorln__

```go
func colorln(args ...interface{}) string
```

<pre>
Same as color, but using sprintln instead of sprint to format
arguments
</pre>

## __concat__

```go
func concat(args ...interface{}) string
```

<pre>
Returns the concatenation (without separator) of the string
representation of objects.
</pre>

## __formatList__

```go
func formatList(format string, list ...interface{}) IGenericList
```

<pre>
Return a list of strings by applying the format to each element of the
supplied list.

You can also use autoWrap as Razor expression if you don't want to
specify the format.
The format is then automatically induced by the context around the
declaration).
Valid aliases for autoWrap are: aWrap, awrap.

Ex:
    Hello @&lt;autoWrap(to(10)) World!
</pre>

### Aliases

- _autoWrap_
- _aWrap_
- _awrap_

## __id__

```go
func id(identifier string, replaceChar ...interface{}) string
```

<pre>
Returns a valid go identifier from the supplied string (replacing any
non compliant character by replacement, default _ ).
</pre>

### Aliases

- _identifier_

## __iif__

```go
func iif(testValue interface{}, valueTrue interface{}, valueFalse
interface{}) interface{}
```

<pre>
If testValue is empty, returns falseValue, otherwise returns
trueValue.
    WARNING: All arguments are evaluated and must by valid.
</pre>

### Aliases

- _ternary_

## __indent__

```go
func indent(nbSpace int, args ...interface{}) string
```

<pre>
Indents every line in a given string to the specified indent width.
This is useful when aligning multi-line strings.
</pre>

## __joinLines__

```go
func joinLines(format ...interface{}) string
```

<pre>
Merge the supplied objects into a newline separated string.
</pre>

## __lorem__

```go
func lorem(loremType interface{}, params ...int) string
```

<pre>
Returns a random string. Valid types are be word, words, sentence,
para, paragraph, host, email, url.
</pre>

### Aliases

- _loremIpsum_

## __mergeList__

```go
func mergeList(lists ...IGenericList) IGenericList
```

<pre>
Return a single list containing all elements from the lists supplied.
</pre>

## __nIndent__

```go
func nIndent(nbSpace int, args ...interface{}) string
```

<pre>
Work as indent but add a newline before.
</pre>

### Aliases

- _nindent_

## __raw__

```go
func raw(args ...interface{}) interface{}
```

<pre>
Print the arguments outside of their enclosing quotes
</pre>

### Aliases

- _printRaw_

## __reCompile__

```go
func reCompile(arg1 string) *regexp.Regexp
```

<pre>
Parses a regular expression and returns Regexp object that can be used
to match against text.
</pre>

## __repeat__

```go
func repeat(n int, element interface{}) IGenericList
```

<pre>
Returns an array with the item repeated n times.
</pre>

## __sIndent__

```go
func sIndent(spacer string, args ...interface{}) string
```

<pre>
Indents the elements using the provided spacer.

You can also use autoIndent as Razor expression if you don't want to
specify the spacer.
Spacer will then be auto determined by the spaces that precede the
expression.
Valid aliases for autoIndent are: aIndent, aindent.
</pre>

### Aliases

- _sindent_
- _spaceIndent_
- _autoIndent_
- _aindent_
- _aIndent_

## __splitLines__

```go
func splitLines(content interface{}) []interface{}
```

<pre>
Returns a list of strings from the supplied object with newline as the
separator.
</pre>

## __stripColor__

```go
func stripColor(arg1 interface{}) string
```

<pre>
Remove all ANSI colors & attributes from a string.
</pre>

### Aliases

- _stripansi_
- _stripANSI_
- _striptcolor_

## __wrap__

```go
func wrap(width interface{}, args ...interface{}) string
```

<pre>
Wraps the rendered arguments within width.
</pre>

### Aliases

- _wrapped_
