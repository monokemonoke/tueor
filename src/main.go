package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path"
)

var (
	//go:embed templates/start/templates/*
	helloworld embed.FS

	//go:embed templates/start/params.json
	helloworldParams []byte
)

func main() {
	fileSystem, err := fs.Sub(helloworld, "templates/start/templates")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(fileSystem)

	var params map[string]map[string]interface{}
	if err := json.Unmarshal(helloworldParams, &params); err != nil {
		log.Fatal(err)
		return
	}

	readEntries(fileSystem, ".", params)
}

func readEntries(fsys fs.FS, dirname string, params map[string]map[string]interface{}) error {
	fmt.Println("A", fsys, dirname)
	entries, err := fs.ReadDir(fsys, dirname)
	if err != nil {
		fmt.Println("D", err)
		return err
	}
	fmt.Println("B", len(entries))

	for _, entry := range entries {
		filename := dirname + "/" + entry.Name()
		if filename[:2] == "./" {
			filename = filename[2:]
		}
		fmt.Println("C", filename)

		if entry.IsDir() {
			readEntries(fsys, filename, params)
		} else {
			t := template.Must(template.ParseFS(fsys, filename))
			param := params[filename]
			generate(t, param, filename[:len(filename)-4])
		}
	}
	return nil
}

func generate(t *template.Template, params interface{}, filename string) {
	var buf bytes.Buffer
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.Write(buf.Bytes())

	nowDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.MkdirAll(path.Dir(path.Join(nowDir, filename)), 0777); err != nil {
		log.Fatalln(err)
	}

	if err := os.WriteFile(path.Join(nowDir, filename), out.Bytes(), 0644); err != nil {
		log.Fatalln(err)
	}
}
