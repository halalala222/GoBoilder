package project

import _ "embed"

//go:embed gitignore.tmpl
var GitIgnoreTemplate []byte

//go:embed README.tmpl
var ReadmeTemplate []byte

//go:embed Makefile.tmpl
var MakefileTemplate []byte
