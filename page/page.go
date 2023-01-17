package page

import (
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
	prefix:="./content/"
	suffix:=".html"
	handler:=func (w http.ResponseWriter,r *http.Request){
		log.Println(r.Method,r.URL)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(prefix+filename+suffix)
			t.Execute(w,nil)
		}
	}
	http.HandleFunc(pattern,handler)
}