package templates

import (
	_ "embed"

	"github.com/levinion/gorush/util"
)

//go:embed post/basic.html
var basic string

func GetPosts() {
	GetBasic()
}

func GetBasic() {
	util.MakeTemplateFile("posts", "Basic", basic)
}
