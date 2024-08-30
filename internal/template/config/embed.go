package config

import (
	_ "embed"
)

//go:embed config.go.tmpl
var loaderTemplate []byte

//go:embed config.toml.tmpl
var tomlTemplate []byte

//go:embed config.yaml.tmpl
var yamlTemplate []byte

//go:embed config.json.tmpl
var jsonTemplate []byte

//go:embed config.env.tmpl
var envTemplate []byte
