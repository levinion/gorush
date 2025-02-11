package builder

import (
	"net/http"
	"path/filepath"

	"github.com/levinion/gorush/util"

	"github.com/spf13/viper"
)

func (builder *Builder) RenderAssets() {
	fs := http.FileServer(http.Dir("docs/assets/css/"))
	builder.Mux.Handle("/assets/css/", http.StripPrefix("/assets/css/", fs))
}

func (builder *Builder) RenderCategory() {
	pattern := "/category/"
	filename := filepath.Join("docs", "category", "index.html")
	handler := util.RenderHandle(filename)
	builder.Mux.HandleFunc(pattern, handler)
}

func (builder *Builder) RenderAllContentOnlyPages() {
	pages := viper.GetStringSlice("pages.contentOnlyPages")
	for _, page := range pages {
		builder.RenderContentOnlyPage(page)
	}
}

func (builder *Builder) RenderContentOnlyPage(name string) {
	pattern := "/" + name + "/"
	filename := filepath.Join("docs", name, "index.html")
	handler := util.RenderHandle(filename)
	builder.Mux.HandleFunc(pattern, handler)
}

func (builder *Builder) RenderContents() {
	pattern := "/posts/"
	filename := filepath.Join("docs", "posts", "index.html")
	handler := util.RenderHandle(filename)
	builder.Mux.HandleFunc(pattern, handler)
}

func (builder *Builder) RenderEachCategoryPosts() {
	// 分别完成pattern和handler
	categoryList := util.ReturnCategoryList(&builder.Posts)
	for _, category := range categoryList {
		// pattern从/posts开始，到meta中定义的Title结束，与public结构保持一致
		pattern := filepath.Join("/posts", category) + "/"
		filename := filepath.Join("docs", "posts", category, "index.html")

		handler := util.RenderHandle(filename)
		builder.Mux.HandleFunc(pattern, handler)
	}
}

func (builder *Builder) RenderIndex() {
	pattern := "/"
	filename := filepath.Join("docs", "index.html")
	handler := util.RenderHandle(filename)
	builder.Mux.HandleFunc(pattern, handler)
}

func (builder *Builder) RenderPosts() {
	for _, post := range builder.Posts {
		// pattern从/posts开始，到meta中定义的Title结束，与public结构保持一致
		pattern := filepath.Join("/posts", post.Category, post.Filename+"/")
		filename := filepath.Join("docs", "posts", post.Category, post.Filename, "index.html")

		handler := util.RenderHandle(filename)
		builder.Mux.HandleFunc(pattern, handler)
	}
}
