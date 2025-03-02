package core

import (
	"embed"
)

//go:embed schema.cue
var schemaFS embed.FS

type Template struct {
	Template string `json:"template"`
	Path     string `json:"path"`
}

type Config struct {
	Templates []Template `json:"templates"`
}

type Data map[string]Config
