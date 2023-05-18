package templates

import (
	"embed"

	"github.com/monokemonoke/tueor/src"
)

var (
	//go:embed start/*
	Start embed.FS

	StartConfig = src.TemplateConfig{
		Embed: Start,
		Dir:   "start",
		Params: src.Params{
			"go.mod.txt": {
				"moduleName": "example.com/monokemonoke/hoge",
			},
			"src/main.go.txt": {
				"message": "example.com/monokemonoke/hoge",
			},
		},
	}
)
