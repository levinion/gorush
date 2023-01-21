package render

import (
	"github.com/levinion/gorush/log"
	"net/http"
	"text/template"
	"github.com/levinion/gorush/parser"
	"github.com/levinion/gorush/model"
)

func RenderPage(pattern,filename string){
	handler:=func (w http.ResponseWriter,r *http.Request){
		log.Info(r)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(filename)
			t.Execute(w,nil)
		}
	}
	http.HandleFunc(pattern,handler)
	log.Handle(pattern)
}

func GroupRenderMarkdown(renderRoot,markdownRoot,categoryTemplateName,markdownTemplateName string){
	ParseResults:=parser.ParseMarkdown(markdownRoot)
	for _,parseResult:=range ParseResults{
		handler:=handleResult(parseResult, markdownTemplateName)
		pattern:=renderRoot+parseResult.Filename
		http.HandleFunc(pattern,handler)
	}
	categoryRenderMarkdown(renderRoot,categoryTemplateName,ParseResults)
}

func categoryRenderMarkdown(renderRoot,categoryTemplateName string,parseResult []*model.ParseResult){
	handler:=func(w http.ResponseWriter,r *http.Request){
		log.Info(r)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(categoryTemplateName)
			t.Execute(w,parseResult)
		}
	}
	pattern:=renderRoot
	http.HandleFunc(pattern,handler)
}

func handleResult(parseResult *model.ParseResult,templateName string)func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		log.Info(r)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(templateName)
			t.Execute(w,parseResult)
		}
	}
}
