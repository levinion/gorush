//对viper包进行重新封装，使用泛型改写Get函数
package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init(){
	viper.SetConfigFile("./config.toml")
	err:=viper.ReadInConfig()
	if err!=nil{
		log.Println("Read config failed:",err)
		return
	}
	log.Println("Read config success!")
}

func Get[T any](key string)T{
	return viper.Get(key).(T)
}