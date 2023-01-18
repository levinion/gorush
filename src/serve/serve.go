package serve

import (
	"gorush/src/config"
	"log"
	"net/http"
)

func Run(){
	addr:=config.Get[string]("server.addr")
	log.Println("Start listening at http://"+addr+"...")
	http.ListenAndServe(addr,nil)
}