package builder

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"sort"
	"time"

	"github.com/levinion/gorush/internal/model"
	"github.com/levinion/gorush/internal/util"

	"github.com/spf13/viper"
)

func (builder *Builder) DumpAllContentOnlyPages() {
	pages := viper.GetStringSlice("pages.contentOnlyPages")
	for _, page := range pages {
		builder.DumpContentOnlyPage(page)
	}
}

func (builder *Builder) DumpContentOnlyPage(name string) {
	path := filepath.Join("content", "pages", name+".md")
	//由于没有元数据，所以舍弃
	buf, _, err := builder.ParseMarkdown(path)
	if err != nil {
		panic(err)
	}
	page := model.Page{
		Filename: name,
		Content:  template.HTML(buf.String()),
	}
	builder.Repo.Pages[name] = page
}

// 读取content/posts下目录和文章，存取到builder.repo
func (builder *Builder) DumpPosts() {
	WalkPosts(func(dirName string, file fs.DirEntry) {
		builder.readPostsOfCategory(dirName, file)
	})
	//按时间对文章进行排序
	builder.sortPostsByTime()
}

func (builder *Builder) readPostsOfCategory(dirName string, file fs.DirEntry) {

	filename := filepath.Join(dirName, file.Name())
	if filepath.Ext(filename) != ".md" {
		return
	}
	buf, metaData, err := builder.ParseMarkdown(filename)
	if err != nil {
		panic(err)
	}

	post := model.Post{
		Category: filepath.Base(dirName),
		Filename: util.TrimExt(file.Name()),
		MetaData: model.MetaData{
			Title:   metaData["title"].(string), //断言为string类型
			Created: metaData["created"].(string),
		},
		Content: template.HTML(buf.String()),
	}
	builder.Repo.Posts = append(builder.Repo.Posts, post)

}

// 按时间对文章进行排序
func (builder *Builder) sortPostsByTime() {
	sort.Slice(builder.Repo.Posts, func(i, j int) bool {
		timeI, _ := time.ParseInLocation(
			"2006-01-02 15:04:05",
			builder.Repo.Posts[i].MetaData.Created,
			time.Local,
		)
		timeJ, _ := time.ParseInLocation(
			"2006-01-02 15:04:05",
			builder.Repo.Posts[j].MetaData.Created,
			time.Local,
		)
		return timeI.After(timeJ)
	})
}
