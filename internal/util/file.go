package util

import (
	"fmt"
	"gorush/internal/model"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// 去除文件扩展名，可传入完整路径或文件BaseName
func TrimExt(filename string) string {
	baseName := filepath.Base(filename)
	ext := filepath.Ext(filename)
	r := strings.TrimSuffix(baseName, ext)
	return r
}

func CreateFileWithContent(content string, name ...string) {
	filename := filepath.Join(name...)
	if IsExist(filename) {
		return
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprint(f, content)
}

// 只支持单级目录的拷贝
func CopyFilesAll(thisDir, toDir string) {
	os.MkdirAll(toDir, os.ModePerm)
	filepath.Walk(thisDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			toPath := filepath.Join(toDir, info.Name())
			from, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(toPath, from, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
}

func SimpleMkdir(em ...string) {
	os.MkdirAll(filepath.Join(em...), os.ModePerm)
}

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func IsNotIn[T string | int | float32 | float64](arr []T, target T) bool {
	for _, t := range arr {
		if t == target {
			return false
		}
	}
	return true
}

func ReturnCategoryList(posts *[]model.Post) []string {
	var arr = make([]string, 0)
	for _, post := range *posts {
		if IsNotIn(arr, post.Category) {
			arr = append(arr, post.Category)
		}
	}
	return arr
}
