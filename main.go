package main

import (
	"gorush/page"
	"log"
	"net/http"
)


func main(){
	page.NewPage("/","mainPage")
	page.NewPage("/posts","posts")
	log.Println("Start listening at http://localhost:9090/")
	http.ListenAndServe(":9090",nil)
}