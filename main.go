package main

import (
	"gorush/src/page"
	"gorush/src/serve"
	"gorush/src/config"
)

func main(){
	config.Init()
	page.NewPage("/","mainPage")
	page.NewPage("/posts/","posts")
	page.NewPage("/about/","about")
	page.InitPosts()
	serve.Run()
}