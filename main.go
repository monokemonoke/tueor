package main

import (
	"fmt"
	"os"

	"github.com/monokemonoke/tueor/src/templates"
)

func main() {
	err := templates.HelloworldConfig.Generate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
