package pages

import (
	"github.com/levinion/gorush/util"
)

func Get() {
	util.CreateFileWithContent("", "content", "pages", "about.md")
}
