# Problem

`Problem` is a simple Go library implementing [RFC 7807](https://datatracker.ietf.org/doc/html/rfc7807): Problem Details
for HTTP APIs.

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
