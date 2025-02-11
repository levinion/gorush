package builder

import (
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/levinion/gorush/util"

	"github.com/spf13/viper"
)

var wg sync.WaitGroup

// 并发在posts目录下处理事务
func WalkPosts(f func(dirName string, file fs.DirEntry)) {
	dir := filepath.Join("content", "posts")
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(entries)
	}
	excludeDirs := viper.GetStringSlice("posts.excludeDirs")

	for _, entry := range entries {
		if !entry.IsDir() || !util.IsNotIn(excludeDirs, entry.Name()) {
			continue
		}
		dirName := filepath.Join(dir, entry.Name())
		wg.Add(1)
		// 并发读取文章
		go func() {
			files, err := os.ReadDir(dirName)
			if err != nil {
				panic(err)
			}
			for _, file := range files {
				f(dirName, file)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
