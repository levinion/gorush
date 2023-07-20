package root

import (
	"embed"
	"path/filepath"

	"github.com/levinion/gorush/util"
)

//go:embed resources/*
var resources embed.FS

func Get() {
	files, err := resources.ReadDir("resources")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		content, err := resources.ReadFile(filepath.Join("resources", file.Name()))
		if err != nil {
			panic(err)
		}
		util.CreateFileWithContent(string(content), file.Name())
	}
}
