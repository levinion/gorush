package render

import (
	"github.com/levinion/gorush/log"
	"net/http"
	"text/template"
	"github.com/levinion/gorush/parser"
	"github.com/levinion/gorush/model"
	"github.com/levinion/gorush/config"
)

// 将HTML渲染成页面，需传入路由路径和文件路径
func RenderHTML(pattern,filename string){
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

/* 
将文件夹下所有markdown文件渲染成目录和文章，
需分别传入路由路径、markdown文件所在文件夹路径、目录页面模板、文章页面模板
*/
func GroupRenderMarkdown(renderRoot,markdownRoot,categoryTemplateName,markdownTemplateName string){
	parseResults:=parser.GroupParseMarkdown(markdownRoot)
	for _,parseResult:=range parseResults{
		handler:=handleParseResult(parseResult, markdownTemplateName)
		pattern:=renderRoot+parseResult.Filename
		http.HandleFunc(pattern,handler)
	}
	categoryRenderMarkdown(renderRoot,categoryTemplateName,parseResults)
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

func handleParseResult(parseResult *model.ParseResult,templateName string)func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		log.Info(r)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(templateName)
			t.Execute(w,parseResult)
		}
	}
}

func RenderMarkdown(pattern,filename,templateName string){
	parseResult:=parser.ParseMarkdown(filename)
	handler:=handleParseResult(parseResult,templateName)
	http.HandleFunc(pattern,handler)
}

func Run(){
	addr:=config.Get[string]("server.addr")
	log.Listen(addr)
	http.ListenAndServe(addr,nil)
}
