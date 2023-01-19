package parse

import (
	"log"
	"os"
	"path"
	"strings"

    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark-meta"
	"bytes"
)

type ParseResult struct{
	Filename string
	Content string
	MetaData map[string]any
}

func ParseMarkdown() []*ParseResult{
	dir:="./posts/"
	files,err:=os.ReadDir(dir)
	if err!=nil{
		log.Println("Markdown parse failed:",err.Error())
		return nil
	}

	markdown:=goldmark.New(
		goldmark.WithExtensions(meta.Meta,),
	)


	r:=make([]*ParseResult,len(files))
	for i,file:=range files{
		if path.Ext(file.Name())!=".md"{
			continue
		}
		in,err:=os.ReadFile(dir+file.Name())
		if err!=nil{
			log.Println("Read file failed:",err.Error())
			return nil
		}
		fileNameWithoutSuffix:=strings.TrimSuffix(file.Name(),".md")
		context:=parser.NewContext()
		var buf bytes.Buffer
		markdown.Convert(in,&buf,parser.WithContext(context))
		metaData:=meta.Get(context)
		r[i] = &ParseResult{fileNameWithoutSuffix,buf.String(),metaData}
	}
	log.Println("Parse Markdown success!")
	return r
}