// 基于viper包
package config

import (
	"github.com/levinion/gorush/log"
	"os"
	"github.com/spf13/viper"
)

type Config struct{
	viper.Viper
}

// config初始化
func (cfg *Config)Init(){
	cfg.SetConfigFile("./config.toml")
	err:=cfg.ReadInConfig()
	if err!=nil{
		log.Println("Read config failed:",err)
		return
	}
	log.Println("Read config success!")
}

//新建config
func New(projectName string){
	log.Println("Config file not exist. Creating config.toml...")
	c,err:=os.OpenFile("./"+projectName+"config.toml",os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err!=nil{
		log.WriteError(err)
	}
	c.WriteString(ConfigTemplate)
}

// 配置文件不存在时返回ture，否则返回false
func NotExist() bool{
	_,err:=os.Stat("./config.toml")
	return os.IsNotExist(err)
}

var ConfigTemplate=
`[version]
version = "v0.0.3"

[server]
addr = "localhost:9090"`