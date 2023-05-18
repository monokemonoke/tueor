package src

import "embed"

var (
	//go:embed templates/start/*
	Start embed.FS

	StartDir = "templates/start"

	StartParams Params = Params{
		"go.mod.txt": {
			"moduleName": "example.com/monokemonoke/hoge",
		},
		"src/main.go.txt": {
			"message": "example.com/monokemonoke/hoge",
		},
	}
)
