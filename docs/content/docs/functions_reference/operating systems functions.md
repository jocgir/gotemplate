---
bookToC: 2
weight: 2
---
# Operating systems functions
<!-- markdownlint-disable MD033 MD024 --->

## __diff__

```go
func diff(text1 interface{}, text2 interface{}) interface{}
```

<pre>
Returns a colored string that highlight differences between supplied
texts.
</pre>

### Aliases

- _difference_

## __exists__

```go
func exists(filename interface{}) bool
```

<pre>
Determines if a file exists or not.
</pre>

### Aliases

- _fileExists_
- _isExist_

## __glob__

```go
func glob(args ...interface{}) IGenericList
```

<pre>
Returns the expanded list of supplied arguments (expand *[]? on
filename).
</pre>

### Aliases

- _expand_

## __group__

```go
func group() *user.Group
```

<pre>
Returns the current user group information (user.Group object).
</pre>

### Aliases

- _userGroup_

## __home__

```go
func home() string
```

<pre>
Returns the home directory of the current user.
</pre>

### Aliases

- _homeDir_
- _homeFolder_

## __isDir__

```go
func isDir(filename interface{}) bool
```

<pre>
Determines if the file is a directory.
</pre>

### Aliases

- _isDirectory_
- _isFolder_

## __isExecutable__

```go
func isExecutable(filename interface{}) bool
```

<pre>
Determines if the file is executable by the current user.
</pre>

## __isFile__

```go
func isFile(filename interface{}) bool
```

<pre>
Determines if the file is a file (i.e. not a directory).
</pre>

## __isReadable__

```go
func isReadable(filename interface{}) bool
```

<pre>
Determines if the file is readable by the current user.
</pre>

## __isWriteable__

```go
func isWriteable(filename interface{}) bool
```

<pre>
Determines if the file is writeable by the current user.
</pre>

## __joinPath__

```go
func joinPath(args ...string) string
```

<pre>
Joins any number of path elements into a single path, adding a
separating slash if necessary. The result is Cleaned; in particular
all empty strings are ignored.
</pre>

## __lastMod__

```go
func lastMod(filename interface{}) time.Time
```

<pre>
Returns the last modification time of the file.
</pre>

### Aliases

- _lastModification_
- _lastModificationTime_

## __lookPath__

```go
func lookPath(arg1 interface{}) string
```

<pre>
Returns the location of the specified executable (returns empty string
if not found).
</pre>

### Aliases

- _whereIs_
- _look_
- _which_
- _type_

## __mode__

```go
func mode(filename interface{}) os.FileMode
```

<pre>
Returns the file mode.
</pre>

### Aliases

- _fileMode_

## __pwd__

```go
func pwd() string
```

<pre>
Returns the current working directory.
</pre>

### Aliases

- _currentDir_

## __save__

```go
func save(filename string, object interface{}) string
```

<pre>
Save object to file.
</pre>

### Aliases

- _write_
- _writeTo_

## __size__

```go
func size(filename interface{}) int64
```

<pre>
Returns the file size.
</pre>

### Aliases

- _fileSize_

## __stat__

```go
func stat(arg1 string) os.FileInfo
```

<pre>
Returns the file Stat information (os.Stat object).
</pre>

### Aliases

- _fileStat_

## __user__

```go
func user() *user.User
```

<pre>
Returns the current user information (user.User object).
</pre>

### Aliases

- _currentUser_

## __username__

```go
func username() string
```

<pre>
Returns the current user name.
</pre>
