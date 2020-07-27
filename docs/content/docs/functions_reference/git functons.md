---
bookToC: 2
weight: 2
---
# Git functons
<!-- markdownlint-disable MD033 MD024 --->

## __currentBranch__

```go
func currentBranch(path string) string
```

<pre>
Returns the name of the currently checked out git branch at the given
path
</pre>

## __currentCommit__

```go
func currentCommit(path string) string
```

<pre>
Returns the hash of the currently checked out git commit at the given
path
</pre>

## __origin__

```go
func origin(path string) string
```

<pre>
Returns the git origin remote URL at the given path
</pre>
