package parser

import (
	"github.com/levinion/gorush/log"
	"os"
	"path"
	"strings"

    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark-meta"
	"github.com/levinion/gorush/model"
	"bytes"
)


func ParseMarkdown(dir string) []*model.ParseResult{
	files,err:=os.ReadDir(dir)
	if err!=nil{
		log.ParseError(err)
		return nil
	}

	markdown:=goldmark.New(
		goldmark.WithExtensions(meta.Meta,),
	)

	r:=make([]*model.ParseResult,len(files))
	for i,file:=range files{
		if path.Ext(file.Name())!=".md"{
			continue
		}
		in,err:=os.ReadFile(dir+file.Name())
		if err!=nil{
			log.ReadError(err)
			return nil
		}
		fileNameWithoutSuffix:=strings.TrimSuffix(file.Name(),".md")
		context:=parser.NewContext()
		var buf bytes.Buffer
		markdown.Convert(in,&buf,parser.WithContext(context))
		metaData:=meta.Get(context)
		r[i] = &(model.ParseResult{Filename: fileNameWithoutSuffix,Content: buf.String(),MetaData: metaData})
	}
	log.ParseSuccess()
	return r
}