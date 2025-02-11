package templates

import (
	_ "embed"
	"html/template"
	"os"
	"path/filepath"

	"github.com/levinion/gorush/model"
)

func Get() {
	GetPages()
	GetPosts()
	GetCommon()
}

//go:embed new/default.tmpl
var defaultPost string

// 新建一个包含元数据的默认文章模板
func MakeStdPostTemplateFile(file string, meta model.MetaData) {
	tmpl := defaultPost
	dirname := filepath.Dir(file)
	filename := filepath.Base(file)
	dirAddr := filepath.Join("content", "posts", dirname)
	fileAddr := filepath.Join(dirAddr, filename+".md")
	os.MkdirAll(dirAddr, os.ModePerm)
	f, err := os.OpenFile(fileAddr, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t, err := template.New("default").Delims("{{{", "}}}").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	t.Execute(f, meta)
}

// 新建一个普通页面模板（空文件）
func MakeCommonPage(name string) {
	filename := filepath.Join("content", "pages", name+".md")
	os.Create(filename)
}
