package util

import (
	"log"
	"os/exec"
	"os"
)

// func Copy(src string,des string){
// 	input,err:=os.OpenFile(src,os.O_RDONLY,0666)
// 	if err!=nil{
// 		log.Println("Read File Failed:",err)
// 	}
// 	output,err:=os.OpenFile(des,os.O_CREATE|os.O_WRONLY,0666)
// 	if err!=nil{
// 		log.Println("Write File Failed:",err)
// 	}
// 	io.Copy(output,input)
// }

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