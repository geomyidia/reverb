# reverb

*A gRPC Server API Convenience Wrapper*

## About

Reverb is cheekily synonymized from the HTTPD package [Echo](https://echo.labstack.com/) and aims to serve a minimally similar purpose, though for gRPC instead of HTTP.

## Example Use

In particular, see the `StartgRPCD` function below:

```go
package app

import (
  "fmt"

  "github.com/labstack/echo/v4"
  "github.com/geomyidia/reverb"
)

// Application ...
type Application struct {
  Config *cfg.Config
  HTTPD  *echo.Echo
  GRPCD  *reverb.Reverb
}

// StartgRPCD ...
func (a *Application) StartgRPCD() {
  serverOpts := fmt.Sprintf("%s:%d", a.Config.GRPCD.Host, a.Config.GRPCD.Port)
  server := a.GRPCD.Start(serverOpts)
  a.SetupgRPCImplementation(server)
  server.Serve()
}

// Setup for other services, e.g., HTTPD
// ...

// Start all services
func (a *Application) Start() {
  a.StartgRPCD()
  a.StartHTTPD()
  ...
}
```
