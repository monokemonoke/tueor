package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"html/template"
	"log"
	"os"
)

var (
	//go:embed templates/test.txt
	testTxt string
)

func main() {
	t := template.Must(template.New("test").Parse(testTxt))
	params := map[string]interface{}{
		"message": "message from generator",
	}

	var buf bytes.Buffer
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.WriteString("// Code generated by main.go\n")
	out.Write(buf.Bytes())
	body, err := format.Source(out.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("test.go", body, 0644); err != nil {
		log.Fatalln(err)
	}
}
