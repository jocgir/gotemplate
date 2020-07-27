---
bookToC: 2
weight: 2
---
# Data Conversion
<!-- markdownlint-disable MD033 MD024 --->

## __data__

```go
func data(data interface{}, context ...interface{}) interface{}
```

<pre>
Tries to convert the supplied data string into data structure (Go
spec). It will try to convert HCL, YAML and JSON format. If context is
omitted, default context is used.
</pre>

### Aliases

- _DATA_
- _fromData_
- _fromDATA_

## __hcl__

```go
func hcl(hcl interface{}, context ...interface{}) interface{}
```

<pre>
Converts the supplied hcl string into data structure (Go spec). If
context is omitted, default context is used.
</pre>

### Aliases

- _HCL_
- _fromHcl_
- _fromHCL_
- _tfvars_
- _fromTFVars_
- _TFVARS_
- _fromTFVARS_

## __json__

```go
func json(json interface{}, context ...interface{}) interface{}
```

<pre>
Converts the supplied json string into data structure (Go spec). If
context is omitted, default context is used.
</pre>

### Aliases

- _JSON_
- _fromJson_
- _fromJSON_

## __toBash__

```go
func toBash(value interface{}) string
```

<pre>
Converts the supplied value to bash compatible representation.
</pre>

## __toHcl__

```go
func toHcl(value interface{}) string
```

<pre>
Converts the supplied value to compact HCL representation.
</pre>

### Aliases

- _toHCL_

## __toInternalHcl__

```go
func toInternalHcl(value interface{}) string
```

<pre>
Converts the supplied value to compact HCL representation used inside
outer HCL definition.
</pre>

### Aliases

- _toInternalHCL_
- _toIHCL_
- _toIHcl_

## __toJson__

```go
func toJson(value interface{}) string
```

<pre>
Converts the supplied value to compact JSON representation.
</pre>

### Aliases

- _toJSON_

## __toPrettyHcl__

```go
func toPrettyHcl(value interface{}) string
```

<pre>
Converts the supplied value to pretty HCL representation.
</pre>

### Aliases

- _toPrettyHCL_

## __toPrettyJson__

```go
func toPrettyJson(value interface{}) string
```

<pre>
Converts the supplied value to pretty JSON representation.
</pre>

### Aliases

- _toPrettyJSON_

## __toPrettyTFVars__

```go
func toPrettyTFVars(value interface{}) string
```

<pre>
Converts the supplied value to pretty HCL representation (without
multiple map declarations).
</pre>

## __toQuotedHcl__

```go
func toQuotedHcl(value interface{}) string
```

<pre>
Converts the supplied value to compact quoted HCL representation.
</pre>

### Aliases

- _toQuotedHCL_

## __toQuotedJson__

```go
func toQuotedJson(value interface{}) string
```

<pre>
Converts the supplied value to compact quoted JSON representation.
</pre>

### Aliases

- _toQuotedJSON_

## __toQuotedTFVars__

```go
func toQuotedTFVars(value interface{}) string
```

<pre>
Converts the supplied value to compact HCL representation (without
multiple map declarations).
</pre>

## __toTFVars__

```go
func toTFVars(value interface{}) string
```

<pre>
Converts the supplied value to compact HCL representation (without
multiple map declarations).
</pre>

## __toYaml__

```go
func toYaml(value interface{}) string
```

<pre>
Converts the supplied value to YAML representation.
</pre>

### Aliases

- _toYAML_

## __yaml__

```go
func yaml(yaml interface{}, context ...interface{}) interface{}
```

<pre>
Converts the supplied yaml string into data structure (Go spec). If
context is omitted, default context is used.
</pre>

### Aliases

- _YAML_
- _fromYaml_
- _fromYAML_
