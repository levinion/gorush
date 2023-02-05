package assets

import (
	"embed"
	"os"
	"path/filepath"
)

//go:embed css/*
var assetsDir embed.FS

func Get() {
	files, err := assetsDir.ReadDir("css")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		content, err := assetsDir.ReadFile(filepath.Join("css", file.Name()))
		if err != nil {
			panic(err)
		}
		GetAssets("css", file.Name(), string(content))
	}
}

func GetAssets(t, name, content string) {
	dir := filepath.Join("assets", t)
	os.MkdirAll(dir, os.ModePerm)
	filename := filepath.Join(dir, name)
	err := os.WriteFile(filename, []byte(content), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
