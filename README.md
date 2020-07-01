# libdevice42

This is a code-generated library for the [device42](https://api.device42.com/) REST api.
It uses [go-swagger](https://github.com/go-swagger/go-swagger) to generate the associated structures.

## Getting started

```go
package main

import (
	"fmt"
	"github.com/mjpitz/libdevice42/client"
	"github.com/mjpitz/libdevice42/client/devices"
	"log"
)

func main() {
	d42 := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: "api.device42.com",
		BasePath: "/",
		Schemes: []string{"https"},
	})

	results, err := d42.Devices.GetDevices(&devices.GetDevicesParams{})

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, device := range results.GetPayload().Devices {
		fmt.Println(fmt.Sprintf("%v", device))
	}
}
```
