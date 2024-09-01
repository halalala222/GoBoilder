package rest

import _ "embed"

//go:embed rest_chi.go.tmpl
var restChiTemplate []byte

//go:embed rest_echo.go.tmpl
var restEchoTemplate []byte

//go:embed rest_gin.go.tmpl
var restGinTemplate []byte

//go:embed rest_fiber.go.tmpl
var restFiberTemplate []byte
