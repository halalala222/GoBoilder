package config

import (
	_ "embed"
)

//go:embed config.go.tmpl
var Template []byte

//go:embed config.toml.tmpl
var TOMLTemplate []byte

//go:embed config.yaml.tmpl
var YAMLTemplate []byte

//go:embed config.json.tmpl
var JSONTemplate []byte

//go:embed config.env.tmpl
var ENVTemplate []byte
