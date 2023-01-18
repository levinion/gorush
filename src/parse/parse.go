package parse

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/russross/blackfriday"
)

func ParseMarkdown(){
	inDir:="./origin/"
	outDir:="./pages/posts/"

	files,err:=os.ReadDir(inDir)
	if err!=nil{
		log.Println("Markdown parse failed:",err.Error())
		return
	}
	for _,file:=range files{
		in,err:=os.ReadFile(inDir+file.Name())
		if err!=nil{
			log.Println("Failed on reading files:",err.Error())
			return
		}
		data:=blackfriday.MarkdownCommon(in)
		fileNameWithoutSuffix:=strings.TrimSuffix(file.Name(),".md")
		out,err:=os.OpenFile(outDir+fileNameWithoutSuffix+".html",os.O_CREATE|os.O_WRONLY,0666)
		if err!=nil{
			log.Println("Failed on writing files:",err.Error())
			return
		}
		defer out.Close()
		fmt.Fprint(out,string(data))
	}
	log.Println("Parse Markdown success!")
}