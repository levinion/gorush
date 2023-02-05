package main

import (
	"github.com/levinion/gorush/log"

	"github.com/levinion/gorush/cmd"
)

func main() {
	log.InitLogger()
	cmd.Execute()
}
