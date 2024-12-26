package builder

import (
	"html/template"
	"net/http"
	"path/filepath"

	"os"

	"github.com/yuin/goldmark"

	meta "github.com/yuin/goldmark-meta"

	"github.com/levinion/gorush/model"

	"github.com/yuin/goldmark/parser"

	"bytes"
)

type Builder struct {
	Parser goldmark.Markdown
	model.Repo
	Mux *http.ServeMux
	Counter
}

func NewBuilder() *Builder {
	parser := NewParser()
	counter := NewCounter()
	return &Builder{
		Counter: *counter,
		Parser:  parser,
		Repo: model.Repo{
			Posts: make([]model.Post, 0, counter.PostsNum),
			Pages: make(map[string]model.Page, counter.PagesNum),
		},
		Mux: http.NewServeMux(),
	}
}

func (builder *Builder) Dump() {
	builder.DumpPosts()
	builder.DumpAllContentOnlyPages()
}

// 页面静态化
func (builder *Builder) Freeze(tmpl string) {
	//Posts页面静态化
	builder.FreezePosts(tmpl)
	builder.FreezeIndex()
	builder.FreezeContents()
	builder.FreezeAllContentOnlyPages()
	builder.FreezeCategory()
	builder.FreezeEachCategoryPosts()

	builder.FreezeStatic()
	builder.FreezeAssets()
}

func (builder *Builder) Render() {
	builder.RenderAssets()

	builder.RenderPosts()
	builder.RenderIndex()
	builder.RenderContents()
	builder.RenderAllContentOnlyPages()
	builder.RenderCategory()
	builder.RenderEachCategoryPosts()
}

func (builder *Builder) Run(addr string, c chan os.Signal) {

	server := &http.Server{Addr: addr, Handler: builder.Mux}
	go server.ListenAndServe()
	<-c

}

//以下是工具函数：

func (builder *Builder) ParseMarkdown(path string) (bytes.Buffer, map[string]any, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return bytes.Buffer{}, nil, err
	}

	context := parser.NewContext()
	var buf bytes.Buffer //转化结果存储
	builder.Parser.Convert(file, &buf, parser.WithContext(context))
	metaData := meta.Get(context)
	return buf, metaData, nil
}

// 调用http.ParseFiles方法，附加通用模板
func ParseFiles(filename string) (*template.Template, error) {
	var navTmpl = filepath.Join("templates", "common", "nav", "index.html")
	var headerTmpl = filepath.Join("templates", "common", "header", "index.html")
	return template.ParseFiles(filename, navTmpl, headerTmpl)
}
