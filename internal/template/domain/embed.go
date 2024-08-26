package domain

import _ "embed"

//go:embed article.go.tmpl
var ArticleTemplate []byte

//go:embed errors.go.tmpl
var ErrorsTemplate []byte
