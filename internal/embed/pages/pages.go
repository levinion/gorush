package pages

import (
	"github.com/levinion/gorush/internal/util"
)

func Get(){
	util.CreateFileWithContent("","content","pages","about.md")
}