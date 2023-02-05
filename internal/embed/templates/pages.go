package templates

import (
	"embed"
	"gorush/internal/util"
	"path/filepath"

	"github.com/spf13/viper"
)

//go:embed pages/*
var pages embed.FS

func GetPages() {
	files, err := pages.ReadDir("pages")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.Name() == "contentOnly.html" {
			continue
		}
		content, err := pages.ReadFile(filepath.Join("pages", file.Name()))
		if err != nil {
			panic(err)
		}
		util.MakeTemplateFile("pages", util.TrimExt(file.Name()), string(content))
	}
	GetAllContentOnly()
}

func GetAllContentOnly() {
	allContentOnlyPages := viper.GetStringSlice("pages.contentOnlyPages")
	for _, page := range allContentOnlyPages {
		GetContentOnly(page)
	}
}

func GetContentOnly(name string) {
	content, err := pages.ReadFile("pages/contentOnly.html")
	if err != nil {
		panic(err)
	}
	util.MakeTemplateFile("pages", name, string(content))
}
