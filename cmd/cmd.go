package cmd

import (
	"os"
	"os/exec"

	"github.com/levinion/gorush/log"

	"github.com/levinion/gorush/config"
	"github.com/levinion/gorush/test"
	"github.com/urfave/cli"
)

func Run() {
	app := cli.NewApp()
	app.Name = "gorush"
	app.Author = "levinion"
	app.Usage = "a simple and quick blog easy to build and diy"
	app.Version = "1.0"
	app.Action = func(c *cli.Context) {
		if config.NotExist() {
			config.New("")
		}
		e := exec.Command("go", "run", "./main.go")
		e.Run()
	}

	app.Commands = cli.Commands{
		{
			Name:    "serve",
			Usage:   "启动服务",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) {
				if config.NotExist() {
					config.New("")
				}
				e := exec.Command("go", "run", "./main.go")
				e.Run()
			},
		},
		{
			Name:    "new",
			Usage:   "新建项目",
			Aliases: []string{"n", "create"},
			Action: func(c *cli.Context) {
				New()
			},
		},
		{
			Name:    "test",
			Usage:   "供使用源码构建的",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) {
				if config.NotExist() {
					config.New("")
				}
				test.Run()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}
