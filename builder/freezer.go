package builder

import (
	"os"
	"path/filepath"

	"github.com/levinion/gorush/util"

	"github.com/levinion/gorush/model"

	"github.com/spf13/viper"
)

// 静态化分类界面
func (builder *Builder) FreezeCategory() {
	dir := filepath.Join("docs", "category")
	filename := filepath.Join(dir, "index.html")
	os.MkdirAll(dir, os.ModePerm)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	templateFile := filepath.Join("templates", "pages", "category", "index.html")
	t, err := ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	categoryList := util.ReturnCategoryList(&builder.Posts)
	t.Execute(f, categoryList)
}

// 静态化所有常规页面
func (builder *Builder) FreezeAllContentOnlyPages() {
	pages := viper.GetStringSlice("pages.contentOnlyPages")
	for _, page := range pages {
		builder.FreezeContentOnlyPage(page)
	}
}

// 静态化常规页面
func (builder *Builder) FreezeContentOnlyPage(name string) {
	os.MkdirAll(filepath.Join("docs", name), os.ModePerm)
	filename := filepath.Join("docs", name, "index.html")
	templateFile := filepath.Join("templates", "pages", name, "index.html")
	t, err := ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//传入page对象
	t.Execute(f, builder.Pages[name])
}

// 静态化目录页面
func (builder *Builder) FreezeContents() {
	dir := filepath.Join("docs", "posts")
	filename := filepath.Join(dir, "index.html")
	os.MkdirAll(dir, os.ModePerm)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	templateFile := filepath.Join("templates", "pages", "contents", "index.html")
	t, err := ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	t.Execute(f, builder.Posts)
}

// 静态化每个分类下文章页面
func (builder *Builder) FreezeEachCategoryPosts() {

	//在public/posts中生成静态文件
	templateFile := filepath.Join("templates", "pages", "eachCategoryPosts", "index.html")

	var cMap = make(map[string][]model.Post)

	for _, post := range builder.Posts {
		cMap[post.Category] = append(cMap[post.Category], post)
	}

	for category, posts := range cMap {

		t, err := ParseFiles(templateFile)
		if err != nil {
			panic(err)
		}
		dir := filepath.Join("docs", "posts", category)
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, "index.html")
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		//传入post
		t.Execute(f, posts)
	}
}

// 静态化主页
func (builder *Builder) FreezeIndex() {
	os.MkdirAll(filepath.Join("docs"), os.ModePerm)
	filename := filepath.Join("docs", "index.html")
	templateFile := filepath.Join("templates", "pages", "index", "index.html")
	t, err := ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t.Execute(f, nil)
}

// 静态化文章页面
func (builder *Builder) FreezePosts(tmpl string) {

	//在public/posts中生成静态文件
	templateFile := filepath.Join("templates", "posts", tmpl, "index.html")

	for _, post := range builder.Posts {

		t, err := ParseFiles(templateFile)
		if err != nil {
			panic(err)
		}
		dir := filepath.Join("docs", "posts", post.Category, post.Filename)
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, "index.html")
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		t.Execute(f, post)
	}
}

// 静态化static目录
func (builder *Builder) FreezeStatic() {
	thisDir := "static"
	toDir := filepath.Join("docs")
	util.CopyFilesAll(thisDir, toDir)
}

// 静态化assets目录
func (builder *Builder) FreezeAssets() {
	thisDir := "assets/css"
	toDir := filepath.Join("docs", "assets", "css")
	util.CopyFilesAll(thisDir, toDir)
}
