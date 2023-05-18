package main

import (
	"fmt"
	"os"

	"github.com/monokemonoke/tueor/src"
	"github.com/monokemonoke/tueor/src/templates"
)

func main() {
	configMap := GetConfigMap()
	cfg, ok := configMap[os.Args[1]]
	if !ok {
		fmt.Fprintln(os.Stderr, "Not found")
	}

	err := cfg.Generate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func GetConfigMap() map[string]src.TemplateConfig {
	return map[string]src.TemplateConfig{
		"helloworld": templates.HelloworldConfig,
		"start":      templates.StartConfig,
	}
}
