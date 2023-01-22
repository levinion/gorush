package log

import (
	"log"
	"net/http"
)
//自定义操作失败提示
func Error(content string,e error){
	log.Println(content+":",e.Error())
}
//读取文件失败
func ReadError(e error){
	log.Println("Error happened on reading file(s):",e.Error())
}
//写入文件失败
func WriteError(e error){
	log.Println("Error happened on writing file(s):",e.Error())
}
//转化文件失败
func ParseError(e error){
	log.Println("Error happened on parsing file(s):",e.Error())
}
//自定义操作成功提示
func Success(content string){
	log.Println(content)
}
//读取文件成功
func ReadSuccess(){
	log.Println("Read file(s) Success!")
}
//写入文件成功
func WriteSuccess(){
	log.Println("Write file(s) Success!")
}
//转化文件成功
func ParseSuccess(){
	log.Println("Parse file(s) Success!")
}
//日志输出请求基本信息
func Info(r *http.Request){
	log.Println(r.Method,r.URL.Path)
}
//日志输出路由地址
func Handle(pattern string){
	log.Println("Start handling",pattern)
}
//日志输出监听端口提示
func Listen(addr string){
	log.Println("Start listening at http://"+addr+"...")
}
//日志基本输出功能
func Println(a ...any){
	log.Println(a...)
}