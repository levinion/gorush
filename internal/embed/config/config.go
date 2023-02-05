package config

import (
	_ "embed"
	"gorush/internal/util"
)

//go:embed config.toml
var config string

func Get() {
	util.CreateFileWithContent(config, "config.toml")
}
