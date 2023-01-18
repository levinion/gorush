package main

import (
	"gorush/src/page"
	"gorush/src/parse"
	"gorush/src/serve"
	"gorush/src/config"
)

func main(){
	config.Init()
	parse.ParseMarkdown()
	page.NewPage("/","mainPage")
	page.NewPage("/posts","posts")
	page.NewPage("/about","about")
	serve.Run()
}