package project

import _ "embed"

//go:embed gitignore.tmpl
var gitIgnoreTemplate []byte

//go:embed README.tmpl
var readmeTemplate []byte

//go:embed Makefile.tmpl
var makefileTemplate []byte
