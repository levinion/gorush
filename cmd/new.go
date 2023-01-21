package cmd

import (
	"os"

	"github.com/levinion/gorush/config"
	"github.com/levinion/gorush/log"
	"github.com/levinion/gorush/util"
)

func New(){
	projectName:=os.Args[2]+"/"
	projectAddress:="./"+os.Args[2]

	if util.IsDirExist(projectAddress){
		log.Println("文件夹已存在，操作失败")
	}
	
	os.MkdirAll(projectAddress+"/templates",os.ModePerm)
	os.MkdirAll(projectAddress+"/resources",os.ModePerm)

	config.New(projectName)
}