package model

import (
	"html/template"
)

type Post struct {
	// base name of the file without Ext
	Category string
	Filename string
	MetaData
	Content template.HTML
}

type Page struct {
	Filename string
	MetaData
	Content template.HTML
}

// category->Post1,Post2,Post3...
type Repo struct {
	Posts []Post
	Pages map[string]Page
}
