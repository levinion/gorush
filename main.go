package main

import (
	"gorush/cmd"
	"gorush/log"
)

func main() {
	log.InitLogger()
	cmd.Execute()
}
