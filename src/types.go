package src

import "embed"

type Params map[string]map[string]interface{}

type TemplateConfig struct {
	Embed  embed.FS
	Dir    string
	Params Params
}
