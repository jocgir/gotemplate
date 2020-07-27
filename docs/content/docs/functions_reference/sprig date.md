---
bookToC: 2
weight: 2
---
# [Sprig Date](http://masterminds.github.io/sprig/date.html)
<!-- markdownlint-disable MD033 MD024 --->

## __ago__

```go
func ago(date interface{}) string
```

<pre>
The ago function returns duration from time.Now in seconds resolution.
</pre>

## __date__

```go
func date(fmt string, date interface{}) string
```

<pre>
The date function formats a date
[format](https://golang.org/pkg/time/#Time.Format).
</pre>

## __dateInZone__

```go
func dateInZone(fmt string, date interface{}, zone string) string
```

<pre>
Same as date, but with a timezone.
</pre>

### Aliases

- _date_in_zone_

## __dateModify__

```go
func dateModify(fmt string, date time.Time) time.Time
```

<pre>
The dateModify takes a modification and a date and returns the
timestamp.
</pre>

## __date_modify__

```go
func date_modify(fmt string, date time.Time) time.Time
```

<pre>
The dateModify takes a modification and a date and returns the
timestamp.
</pre>

## __duration__

```go
func duration(second interface{}) string
```

<pre>
Formats a given amount of seconds as a time.Duration.
</pre>

## __durationRound__

```go
func durationRound(duration interface{}) string
```

<pre>
Rounds a given duration to the most significant unit.
</pre>

## __htmlDate__

```go
func htmlDate(date interface{}) string
```

<pre>
The htmlDate function formates a date for inserting into an HTML date
picker input field.
</pre>

## __htmlDateInZone__

```go
func htmlDateInZone(date interface{}, zone string) string
```

<pre>
Same as htmlDate, but with a timezone.
</pre>

## __now__

```go
func now() time.Time
```

<pre>
The current date/time. Use this in conjunction with other date
functions.
</pre>

## __toDate__

```go
func toDate(fmt string, str string) time.Time
```

<pre>
Converts a string to a date. The first argument is the date layout and
the second the date string. If the string can’t be convert it
returns the zero value.
</pre>

## __unixEpoch__

```go
func unixEpoch(date time.Time) string
```

<pre>
Returns the seconds since the unix epoch for a 'time.Time'.
</pre>
