package builder

import (
	"io/fs"
	"os"
	"path/filepath"
)

// counter用以统计当前文章和页面总数，在Builder初始化前调用，以减少内存分配
type Counter struct {
	PostsNum int
	PagesNum int
}

func NewCounter() *Counter {
	postsNum := countPosts()
	pagesNum := countPages()
	return &Counter{PostsNum: postsNum, PagesNum: pagesNum}
}

func countPosts() (postsNum int) {
	WalkPosts(func(dirName string, file fs.DirEntry) {
		postsNum++
	})
	return
}

func countPages() (pagesNum int) {
	files, err := os.ReadDir(filepath.Join("content", "posts"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			pagesNum++
		}
	}
	return
}
