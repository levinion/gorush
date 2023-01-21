package cmd

import (
	"github.com/levinion/gorush/log"
	"os"

	"github.com/levinion/gorush/config"
	"github.com/levinion/gorush/workflow"
	"github.com/urfave/cli"
)

func Run(){
	app:=cli.NewApp()
	app.Name="gorush"
	app.Author="levinion"
	app.Usage="a simple and quick blog easy to build and diy"
	app.Version="1.0"
	app.Action=func(c *cli.Context){
		if config.NotExist(){
			config.New("")
		}
		workflow.Run()
	}

	app.Commands=cli.Commands{
		{
			Name: "serve",
			Usage: "Starting serving the net",
			Aliases: []string{"s"},
			Action: func(c *cli.Context){
				if config.NotExist(){
				config.New("")
				}
				workflow.Run()
			},
		},
		{
			Name: "new",
			Usage : "新建项目",
			Aliases: []string{"n","create"},
			Action: func(c *cli.Context){
				New()
			},
		},
	}

	err:=app.Run(os.Args)
	if err!=nil{
		log.Println(err)
	}
}