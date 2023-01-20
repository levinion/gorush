package page

import (
	"github.com/levinion/gorush/src/config"
	"log"
	"net/http"
	"text/template"
	"github.com/levinion/gorush/src/parse"
)

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

func InitMain(){
	ParseResults:=parse.ParseMarkdown()
	for _,parseResult:=range ParseResults{
		handler:=handleResult(parseResult,"./assets/templates/default.html")
		pattern:="/posts/"+parseResult.Filename
		http.HandleFunc(pattern,handler)
	}
	initPosts(ParseResults)
}

func initPosts(parseResult []*parse.ParseResult){
	handler:=func(w http.ResponseWriter,r *http.Request){
		log.Println(r.Method,r.URL.Path)
		if r.Method=="GET"{
			t,_:=template.ParseFiles("./pages/posts.html")
			t.Execute(w,parseResult)
		}
	}
	pattern:="/posts/"
	http.HandleFunc(pattern,handler)
}

func handleResult(parseResult *parse.ParseResult,templateName string)func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		log.Println(r.Method,r.URL.Path)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(templateName)
			t.Execute(w,parseResult)
		}
	}
}
