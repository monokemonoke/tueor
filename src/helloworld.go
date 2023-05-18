package src

import "embed"

var (
	//go:embed templates/helloworld/*
	Helloworld embed.FS

	HelloworldDir = "templates/helloworld"

	HelloworldParams Params = Params{
		"go.mod.txt": {
			"moduleName": "example.com/monokemonoke/hoge",
		},
		"main.go.txt": {
			"message": "example.com/monokemonoke/hoge",
		},
	}
)
