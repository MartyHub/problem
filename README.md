# Problem

`Problem` is a simple Go library implementing [RFC 7807](https://datatracker.ietf.org/doc/html/rfc7807): Problem Details
for HTTP APIs.

![build](https://github.com/MartyHub/problem/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/MartyHub/problem)](https://goreportcard.com/report/github.com/MartyHub/problem)

## Usage

```go
package sample

import (
	"errors"
	"net/http"
	
	"github.com/MartyHub/problem"
)

func ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	problem.Write(writer, problem.New(http.StatusForbidden, "Your current balance is 30, but that costs 50.").
		Error(errors.New("an optional error")).
		Request(req),
	)
}
```
