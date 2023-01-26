package render

import (
	"os"
	"path"
	"strings"

	"github.com/levinion/gorush/log"
	"github.com/levinion/gorush/util"

	"bytes"

	"github.com/levinion/gorush/model"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

//遍历解析目录下的Markdown文件
func (r *Renderer)GroupParseMarkdown(dir string){
	files,err:=os.ReadDir(dir)
	if err!=nil{
		log.ParseError(err)
		return
	}

	markdown:=goldmark.New(
		goldmark.WithExtensions(meta.Meta,),
	)

	for _,file:=range files{
		if path.Ext(file.Name())!=".md"{
			continue
		}
		in,err:=os.ReadFile(dir+file.Name())
		if err!=nil{
			log.ReadError(err)
			return
		}
		filenameWithoutSuffix:=strings.TrimSuffix(file.Name(),".md")
		context:=parser.NewContext()
		var buf bytes.Buffer
		markdown.Convert(in,&buf,parser.WithContext(context))
		metaData:=meta.Get(context)
		r.Articles[filenameWithoutSuffix] = &(model.Article{Title: filenameWithoutSuffix,Content: buf.String(),MetaData: metaData})
	}
	log.ParseSuccess()
}

//解析单个Markdown文件
func (r *Renderer)ParseMarkdown(filename string){
	if path.Ext(filename)!=".md"{
		log.Println("操作非法：非markdown文件")
		return
	}
	f,err:=os.ReadFile(filename)
	if err!=nil{
		log.ReadError(err)
	}
	markdown:=goldmark.New(
		goldmark.WithExtensions(meta.Meta,),
	)
	filenameWithoutSuffix:=util.TrimFilenameSuffix(filename,".md")
	context:=parser.NewContext()
	var buf bytes.Buffer
	markdown.Convert(f,&buf,parser.WithContext(context))
	metaData:=meta.Get(context)
	r.MdPages[filenameWithoutSuffix]=&(model.Article{Title: filenameWithoutSuffix,Content: buf.String(),MetaData: metaData})
}