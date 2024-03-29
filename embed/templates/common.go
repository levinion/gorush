package templates

import (
	"embed"
	"path/filepath"

	"github.com/levinion/gorush/util"
)

//go:embed common/*
var commons embed.FS

func GetCommon() {
	files, err := commons.ReadDir("common")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		content, err := commons.ReadFile(filepath.Join("common", file.Name()))
		if err != nil {
			panic(err)
		}
		util.MakeTemplateFile("common", util.TrimExt(file.Name()), string(content))
	}
}
