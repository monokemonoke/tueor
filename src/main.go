package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"os"
)

var (
	//go:embed templates/helloworld/templates/*
	helloworld embed.FS

	//go:embed templates/helloworld/params.json
	helloworldParams []byte
)

func readEntries(fsys fs.FS, dirname string, params map[string]map[string]interface{}) error {
	entries, err := fs.ReadDir(fsys, dirname)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			readEntries(fsys, dirname+"/"+entry.Name(), params)
			return nil
		}

		filename := dirname + "/" + entry.Name()
		filename = filename[2:]
		t := template.Must(template.ParseFS(fsys, filename))
		param := params[filename]
		generate(t, param, filename[:len(filename)-4])
	}
	return nil
}

func main() {
	fileSystem, err := fs.Sub(helloworld, "templates/helloworld/templates")
	if err != nil {
		log.Fatal(err)
		return
	}

	var params map[string]map[string]interface{}
	if err := json.Unmarshal(helloworldParams, &params); err != nil {
		log.Fatal(err)
		return
	}

	readEntries(fileSystem, ".", params)
}

func generate(t *template.Template, params interface{}, filename string) {
	var buf bytes.Buffer
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.Write(buf.Bytes())

	if err := os.WriteFile(filename, out.Bytes(), 0644); err != nil {
		log.Fatalln(err)
	}
}
