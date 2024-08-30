package http

import _ "embed"

//go:embed chi.go.tmpl
var ChiConfigTemplate []byte

//go:embed echo.go.tmpl
var EchoConfigTemplate []byte

//go:embed gin.go.tmpl
var GinConfigTemplate []byte

//go:embed fiber.go.tmpl
var FiberConfigTemplate []byte
