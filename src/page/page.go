package page

import (
	"gorush/src/config"
	"log"
	"net/http"
	"text/template"
	"gorush/src/parse"
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
		log.Println(r.Method,r.URL.Path)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(prefix+filename+suffix)
			t.Execute(w,nil)
		}
	}
	http.HandleFunc(pattern,handler)
	log.Println("Start handling:",pattern)
}

func InitPosts(){
	ParseResults:=parse.ParseMarkdown()
	for _,parseResult:=range ParseResults{
		handler:=handleString(parseResult.Content)
		pattern:="/posts/"+parseResult.Filename
		http.HandleFunc(pattern,handler)
	}
}

func handleString(content string)func(w http.ResponseWriter,r *http.Request){
	return	func(w http.ResponseWriter,r *http.Request){
		log.Println(r.Method,r.URL.Path)
		if r.Method=="GET"{
			t,_:=template.ParseFiles("./assets/templates/default.html")
			t.Execute(w,content)
		}
	}
}