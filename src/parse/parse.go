package parse

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/russross/blackfriday"
)

type ParseResult struct{
	Filename string
	Content string
}

func ParseMarkdown() []*ParseResult{
	dir:="./posts/"
	files,err:=os.ReadDir(dir)
	if err!=nil{
		log.Println("Markdown parse failed:",err.Error())
		return nil
	}
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
		r[i] = &ParseResult{fileNameWithoutSuffix,string(blackfriday.MarkdownCommon(in))}
	}
	log.Println("Parse Markdown success!")
	return r
}