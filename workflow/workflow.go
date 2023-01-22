package workflow

import (
	"github.com/levinion/gorush/config"
	"github.com/levinion/gorush/log"
	"net/http"
	"github.com/levinion/gorush/render"
)

func ListenAndServe(){
	addr:=config.Get[string]("server.addr")
	log.Listen(addr)
	http.ListenAndServe(addr,nil)
}

func Run(){
	config.Init()
	render.RenderHTML("/","./templates/mainPage.html")
	render.RenderHTML("/about/","./templates/about.html")
	render.GroupRenderMarkdown("/posts/","./resources/","./templates/posts.html","./templates/default.html")
	// render.RenderMarkdown("/test/","./resources/parts.md","./templates/default.html")
	ListenAndServe()
}