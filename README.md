# libdevice42

This is a code-generated library for the [device42](https://api.device42.com/) REST api.
It uses [go-swagger](https://github.com/go-swagger/go-swagger) to generate the associated structures.

## Getting started

```go
package main

import (
	"github.com/go-openapi/strfmt"
    "github.com/mjpitz/libdevice42/client"
)

func main() {
	strfmt.Registry{}
    client := client.NewHTTPClient()
}
```
