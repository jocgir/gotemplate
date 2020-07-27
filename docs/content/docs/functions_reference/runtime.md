---
bookToC: 2
weight: 2
---
# Runtime
<!-- markdownlint-disable MD033 MD024 --->

## __alias__

```go
func alias(name string, function string, source interface{}, args
...interface{}) string
```

<pre>
Defines an alias (go template function) using the function (exec, run,
include, template). Executed in the context of the caller.
</pre>

## __aliases__

```go
func aliases() []string
```

<pre>
Returns the list of all functions that are simply an alias of another
function.
</pre>

## __allFunctions__

```go
func allFunctions() []string
```

<pre>
Returns the list of all available functions.
</pre>

## __assert__

```go
func assert(test interface{}, message ...interface{}) string
```

<pre>
Raises a formated error if the test condition is false.
</pre>

### Aliases

- _assertion_

## __assertWarning__

```go
func assertWarning(test interface{}, message ...interface{}) string
```

<pre>
Issues a formated warning if the test condition is false.
</pre>

### Aliases

- _assertw_

## __categories__

```go
func categories() []template.FuncCategory
```

<pre>
Returns all functions group by categories.

The returned value has the following properties:
    Name string
    Functions []string
</pre>

## __completeExamples__

```go
func completeExamples() string
```

<pre>
Complete the examples that are not fully generated.
</pre>

## __current__

```go
func current() string
```

<pre>
Returns the current folder (like pwd, but returns the folder of the
currently running folder).
</pre>

## __ellipsis__

```go
func ellipsis(function string, args ...interface{}) interface{}
```

<pre>
Returns the result of the function by expanding its last argument that
must be an array into values. It's like calling function(arg1, arg2,
otherArgs...).
</pre>

## __exec__

```go
func exec(command interface{}, args ...interface{}) interface{}
```

<pre>
Returns the result of the shell command as structured data (as string
if no other conversion is possible).
</pre>

### Aliases

- _execute_

## __exit__

```go
func exit(exitValue int) int
```

<pre>
Exits the current program execution.
</pre>

## __func__

```go
func func(name string, function string, source interface{}, config
interface{}) string
```

<pre>
Defines a function with the current context using the function (exec,
run, include, template). Executed in the context of the caller.
</pre>

## __function__

```go
func function(name string) *template.FuncInfo
```

<pre>
Returns the information relative to a specific function.

The returned value has the following properties:
    Name string
    Description string
    Signature string
    Group string
    Aliases []string
    Arguments string
    Result string
</pre>

## __functions__

```go
func functions() []string
```

<pre>
Returns the list of all available functions (excluding aliases).
</pre>

## __getAttributes__

```go
func getAttributes(arg1 interface{}) string
```

<pre>
List all attributes accessible from the supplied object.
</pre>

### Aliases

- _attr_
- _attributes_

## __getMethods__

```go
func getMethods(arg1 interface{}) string
```

<pre>
List all methods signatures accessible from the supplied object.
</pre>

### Aliases

- _methods_

## __getSignature__

```go
func getSignature(arg1 interface{}) string
```

<pre>
List all attributes and methods signatures accessible from the
supplied object.
</pre>

### Aliases

- _sign_
- _signature_

## __include__

```go
func include(source interface{}, context ...interface{}) interface{}
```

<pre>
Returns the result of the named template rendering (like template but
it is possible to capture the output).
</pre>

## __localAlias__

```go
func localAlias(name string, function string, source interface{}, args
...interface{}) string
```

<pre>
Defines an alias (go template function) using the function (exec, run,
include, template). Executed in the context of the function it maps
to.
</pre>

## __raise__

```go
func raise(args ...interface{}) string
```

<pre>
Raise a formated error.
</pre>

### Aliases

- _raiseError_

## __run__

```go
func run(command interface{}, args ...interface{}) interface{}
```

<pre>
Returns the result of the shell command as string.
</pre>

## __substitute__

```go
func substitute(content string) string
```

<pre>
Applies the supplied regex substitute specified on the command line on
the supplied string (see --substitute).
</pre>

## __templateNames__

```go
func templateNames() []string
```

<pre>
Returns the list of available templates names.
</pre>

## __templates__

```go
func templates() []*template.Template
```

<pre>
Returns the list of available templates.
</pre>

## __userContext__

```go
func userContext() interface{}
```

<pre>
Returns the user context (i.e. all global variables except the
injected constant).
</pre>

### Aliases

- _c_
- _context_
