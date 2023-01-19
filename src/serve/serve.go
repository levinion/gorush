package serve

import (
	"gorush/src/config"
	"log"
	"net/http"
	"gorush/src/page"
)

func ListenAndServe(){
	addr:=config.Get[string]("server.addr")
	log.Println("Start listening at http://"+addr+"...")
	http.ListenAndServe(addr,nil)
}

func Run(){
	config.Init()
	page.NewPage("/","mainPage")
	page.NewPage("/about/","about")
	page.InitMain()
	ListenAndServe()
}