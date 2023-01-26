package util

import (
	"log"
	"os/exec"
	"os"
	"strings"
	"path"
)


// 依赖于系统的暂时性的复制方案
func Copy(src string,des string){
	command:=exec.Command("cp","-r",src,des)
	err:=command.Run()
	if err!=nil{
		log.Println("Copy Failed:",err)
	}
}

func IsDirExist(address string) bool{
	fileInfo,err:=os.Stat(address)
	if os.IsExist(err){
		if fileInfo.IsDir(){
			return true
		}
	}
	return false
}

func TrimFilenameSuffix(filename ,suffix string)string{
	return strings.TrimSuffix(path.Base(filename),suffix)
}