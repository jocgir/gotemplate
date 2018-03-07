# Literals protection

## E-Mail protection

The razor convertor is designed to detect email address such as `john.doe@company.com` or `alert@127.0.0.1`.

But it you type something like `@john.doe@(company.com)`, it will try to resolve variable john.doe and company.com.

The result would be `<no value><no value>` unless you have defined:

```go
@john := data("doe = 123.45")
@company := data("com = @Math.Pi")
```

In that case, the result of `@john.doe@(company.com)` will be `123.453.141592653589793`.

## &#64; protection

You can also render the &#64; characters by writing &#64;&#64;.

So this `@@` will render &#64;.