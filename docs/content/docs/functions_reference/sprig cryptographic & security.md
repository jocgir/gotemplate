---
bookToC: 2
weight: 2
---
# [Sprig Cryptographic & Security](http://masterminds.github.io/sprig/crypto.html)
<!-- markdownlint-disable MD033 MD024 --->

## __adler32sum__

```go
func adler32sum(input string) string
```

<pre>
Computes Adler-32 checksum.
</pre>

## __buildCustomCert__

```go
func buildCustomCert(base64Cert string, base64PrivateKey string)
sprig.certificate
```

<pre>
Allows customizing the certificate.
</pre>

## __decryptAES__

```go
func decryptAES(encoded string, key string) string
```

<pre>
Receives a base64 string encoded by the AES-256 CBC algorithm and
returns the decoded text.
</pre>

## __derivePassword__

```go
func derivePassword(counter uint32, passwordType string, password string,
user string, site string) string
```

<pre>
Derive a specific password based on some shared 'master password'
constraints.
</pre>

## __encryptAES__

```go
func encryptAES(secret string, key string) string
```

<pre>
Encrypts text with AES-256 CBC and returns a base64 encoded string.
</pre>

## __genCA__

```go
func genCA(cn string, nbDays int) sprig.certificate
```

<pre>
Generates a new, self-signed x509 certificate authority.
</pre>

## __genPrivateKey__

```go
func genPrivateKey(type string) string
```

<pre>
Generates a new private key encoded into a PEM block. Type should be:
ecdsa, dsa or rsa.
</pre>

## __genSelfSignedCert__

```go
func genSelfSignedCert(cn string, ipList []interface{}, dnsList
[]interface{}, nbDays int) sprig.certificate
```

<pre>
Generates a new, self-signed x509 certificate.
</pre>

## __genSignedCert__

```go
func genSignedCert(cn string, ipList []interface{}, dnsList []interface{},
nbDays int, ca sprig.certificate) sprig.certificate
```

<pre>
Generates a new, x509 certificate signed by the specified CA.
</pre>

## __htpasswd__

```go
func htpasswd(user string, password string) string
```

<pre>
Takes a username and password and generates a bcrypt hash of the
password. The result can be used for basic authentication on an Apache
HTTP Server.
</pre>

## __sha1sum__

```go
func sha1sum(input string) string
```

<pre>
Computes SHA1 digest.
</pre>

## __sha256sum__

```go
func sha256sum(input string) string
```

<pre>
Computes SHA256 digest.
</pre>
