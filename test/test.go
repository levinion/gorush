package test

import (
	"github.com/levinion/gorush/render"
)

func Run(){
	r:=render.NewRenderer()		//新建Renderer对象，配置文件将在这一步初始化
	r.Parse("./resources/")		//解析Markdown文件目录，若在上一步传入目录则可略去

	//下面是页面添加流程示例：
	r.RenderHTML("/","./templates/mainPage.html")
	// r.RenderHTML("/about/","./templates/about.html")
	r.GroupRenderMarkdown("/posts/","./templates/posts.html","./templates/default.html")
	// render.RenderMarkdown("/test/","./resources/parts.md","./templates/default.html")

	r.Run()		//调用Run以运行服务，可写入端口，否则将使用配置文件中定义的端口
}