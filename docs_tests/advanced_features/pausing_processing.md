# Pausing and resuming

You can pause gotemplate or razor processing with the following comments:

```go
{{ add 4 5 }}
@(4+5)

# gotemplate-pause!
{{ add 4 5 }}
@(4+5)
# gotemplate-resume!

# razor-pause!
{{ add 4 5 }}
@(4+5)
# razor-resume!

{{ add 4 5 }}
@(4+5)
```

Gives the following output:

```go
9
9

# gotemplate-pause!
{{ add 4 5 }}
@(4+5)
# gotemplate-resume!

# razor-pause!
9
@(4+5)
# razor-resume!

9
9
```
