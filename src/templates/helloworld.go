package templates

import (
	"embed"

	"github.com/monokemonoke/tueor/src"
)

var (
	//go:embed helloworld/*
	Helloworld embed.FS

	HelloworldConfig = src.TemplateConfig{
		Embed: Helloworld,
		Dir:   "helloworld",
		Params: src.Params{
			"go.mod.txt": {
				"moduleName": "example.com/monokemonoke/hoge",
			},
			"main.go.txt": {
				"message": "example.com/monokemonoke/hoge",
			},
		},
	}
)
