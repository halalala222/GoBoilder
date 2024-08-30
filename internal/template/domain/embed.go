package domain

import _ "embed"

//go:embed user.go.tmpl
var userTemplate []byte

//go:embed errors.go.tmpl
var errorsTemplate []byte
