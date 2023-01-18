package server

import(
	"log"
	"net/http"
)

func Run(){
	log.Println("Start listening at http://localhost:9090/")
	http.ListenAndServe(":9090",nil)
}