package domain

import _ "embed"

//go:embed user.go.tmpl
var UserTemplate []byte

//go:embed errors.go.tmpl
var ErrorsTemplate []byte
