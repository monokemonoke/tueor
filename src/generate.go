package src

import (
	"bytes"
	"errors"
	"go/format"
	"html/template"
	"io/fs"
	"os"
	"path"
)

func (c *TemplateConfig) Generate() error {
	fsys, err := fs.Sub(c.Embed, c.Dir)
	if err != nil {
		return err
	}

	return readEntries(fsys, ".", c.Params)
}

func readEntries(fsys fs.FS, dirname string, params Params) error {
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

func generate(t *template.Template, params map[string]string, filename string) error {
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

func generateText(t *template.Template, params map[string]string, filename string) error {
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

func generateGo(t *template.Template, params map[string]string, filename string) error {
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
