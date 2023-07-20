package config

import (
	_ "embed"

	"github.com/levinion/gorush/util"
)

//go:embed config.toml
var config string

func Get() {
	util.CreateFileWithContent(config, "config.toml")
}
