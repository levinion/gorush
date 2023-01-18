package page

import (
	"gorush/src/config"
	"log"
	"net/http"
	"text/template"
)

//新建页面的一个示例
//
// func mainPage(w http.ResponseWriter,r *http.Request){
// 	log.Println(r.Method,r.URL)
// 	if r.Method=="GET"{
// 		t,_:=template.ParseFiles("./content/mainPage.html")
// 		t.Execute(w,nil)
// 	}
// }
//
func NewPage(pattern,filename string){
	prefix:=config.Get[string]("page.root")
	suffix:=config.Get[string]("page.ext")
	handler:=func (w http.ResponseWriter,r *http.Request){
		log.Println(r.Method,r.URL)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(prefix+filename+suffix)
			t.Execute(w,nil)
		}
	}
	http.HandleFunc(pattern,handler)
}