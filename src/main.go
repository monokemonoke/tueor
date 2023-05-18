package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
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
	entries, err := fs.ReadDir(fsys, dirname)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		filename := dirname + "/" + entry.Name()
		if filename[:2] == "./" {
			filename = filename[2:]
		}

		if entry.IsDir() {
			readEntries(fsys, filename, params)
		} else {
			t := template.Must(template.ParseFS(fsys, filename))
			param := params[filename]
			generate(t, param, filename)
		}
	}
	return nil
}

func generate(t *template.Template, params interface{}, filename string) error {
	if len(filename) < 4 {
		return errors.New("filename is too short")
	}
	generateFile := filename[:len(filename)-4]
	ext := path.Ext(generateFile)

	if ext == "go" {
		return generateGo(t, params, generateFile)
	}
	return generateText(t, params, generateFile)
}

func generateText(t *template.Template, params interface{}, filename string) error {
	var buf bytes.Buffer
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.Write(buf.Bytes())

	nowDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path.Dir(path.Join(nowDir, filename)), 0777); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(nowDir, filename), out.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func generateGo(t *template.Template, params interface{}, filename string) error {
	var buf bytes.Buffer
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.Write(buf.Bytes())

	body, err := format.Source(out.Bytes())
	if err != nil {
		return err
	}

	nowDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path.Dir(path.Join(nowDir, filename)), 0777); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(nowDir, filename), body, 0644); err != nil {
		return err
	}
	return nil
}
