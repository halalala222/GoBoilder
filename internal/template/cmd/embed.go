package cmd

import _ "embed"

//go:embed app.go.tmpl
var appGoTmpl []byte
