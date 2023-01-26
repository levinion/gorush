package render

import (
	"net/http"
	"text/template"

	"github.com/levinion/gorush/config"
	"github.com/levinion/gorush/log"
	"github.com/levinion/gorush/model"
	"github.com/levinion/gorush/util"
)


type Renderer struct{
	Articles map[string]*model.Article
	MdPages map[string]*model.Article
	Handlers map[string]*func(w http.ResponseWriter,r *http.Request)
	*config.Config
}

//新建一个Render对象
func NewRenderer(parseDir ...string)*Renderer{
	cfg:=config.Config{}
	cfg.Init()
	if len(parseDir)==0{
		return 	&Renderer{
			Articles: make(map[string]*model.Article),
			MdPages: make(map[string]*model.Article),
			Handlers: make(map[string]*func(w http.ResponseWriter, r *http.Request)),
			Config: &cfg,
		}
	}
	r:=&Renderer{
		Articles: make(map[string]*model.Article, 0),
		MdPages: make(map[string]*model.Article),
		Handlers: make(map[string]*func(w http.ResponseWriter, r *http.Request)),
		Config: &cfg,
	}
	r.GroupParseMarkdown(parseDir[0])
	return r
}

// 直接使用HTML渲染页面，需传入路由路径和文件路径
func (r *Renderer)RenderHTML(pattern,filename string){
	handler:=func (w http.ResponseWriter,r *http.Request){
		log.Info(r)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(filename)
			t.Execute(w,nil)
		}
	}
	r.Handlers[pattern]=&handler
}

/*
 GroupParseMarkdown的简写
*/
func (r *Renderer)Parse(dir string){
	r.GroupParseMarkdown(dir)
}

/* 
将文件夹下所有markdown文件渲染成目录和文章，
需分别传入路由路径、目录页面模板、文章页面模板，
若未在初始化Renderer时传入参数，请先调用Parse方法完成对目录的解析
*/
func (r *Renderer)GroupRenderMarkdown(renderRoot,categoryTemplateName,markdownTemplateName string){
	for _,article:=range r.Articles{
		handler:=makeFuncHandler(article, markdownTemplateName)
		pattern:=renderRoot+article.Title
		r.Handlers[pattern]=&handler
	}
	r.categoryRenderMarkdown(renderRoot,categoryTemplateName)
}

//渲染目录页面，由 GroupRenderMarkdown 函数调用
func (r *Renderer) categoryRenderMarkdown(renderRoot,categoryTemplateName string){
	handler:=func(w http.ResponseWriter,req *http.Request){
		log.Info(req)
		if req.Method=="GET"{
			t,_:=template.ParseFiles(categoryTemplateName)
			t.Execute(w,r.Articles)
		}
	}
	pattern:=renderRoot
	r.Handlers[pattern]=&handler
}

//FuncHandler工厂
func makeFuncHandler(article *model.Article,templateName string)func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		log.Info(r)
		if r.Method=="GET"{
			t,_:=template.ParseFiles(templateName)
			t.Execute(w,article)
		}
	}
}

//使用单个Markdown文件渲染页面，需分别传入路由路径、文件路径、模板路径
func (r *Renderer)RenderMarkdown(pattern,filename,templateName string){
	r.ParseMarkdown(filename)
	filenameWithoutSuffix:=util.TrimFilenameSuffix(filename,".md")
	handler:=makeFuncHandler(r.MdPages[filenameWithoutSuffix],templateName)
	r.Handlers[pattern]=&handler
}

//监听并运行端口
func (r *Renderer)Run(addr ...string){
	var newAddr string
	if len(addr)==0{
		newAddr=r.Config.GetString("server.addr")
	}else if len(addr)==1{
		newAddr=addr[0]
	}else{
		panic("错误，传入多个端口！")
	}
	mux:=http.NewServeMux()
	for k,v:=range r.Handlers{
		mux.HandleFunc(k,*v)
	}
	log.Listen(newAddr)
	http.ListenAndServe(newAddr,mux)
}
