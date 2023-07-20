package util

import (
	"os"
	"path/filepath"
)

// 创建基础模板文件，分别传入模板名称（可自定义）。
// 其中dir有三个个可选值，分别是“posts”和“pages”以及“common”，对应文章、页面模板以及共用的模板
func MakeTemplateFile(dir, name, tmpl string) {
	templateDir := filepath.Join("templates", dir, name)
	os.MkdirAll(templateDir, os.ModePerm)
	// if error 则跳过

	filename := filepath.Join(templateDir, "index.html")
	if IsExist(filename) {
		return
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(tmpl)
}
