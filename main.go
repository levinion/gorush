package main

import (
	"gorush/src/page"
	"gorush/src/serve"
	"gorush/src/parse"
)


func main(){
	parse.ParseMarkdown()
	page.NewPage("/","mainPage")
	page.NewPage("/posts","posts")
	page.NewPage("/about","about")
	server.Run()
}