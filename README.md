# yesno-wtf

This is my first step into writing Go. This is a clone of https://yesno.wtf in which I proxy their api for the actual data.

Build and install:

```
$ go install github.com/dpromanko/yesno-wtf
```
Run:

```
$ $GOPATH/bin/yesno-wtf
```
For the usual https://yesno.wtf experience navigate to:
```
localhost:8080/
```
For the api navigate to:
```
localhost:8080/api
```
If the api isn't pretty enough for you:
```
localhost:8080/api?output=prettyjson
```