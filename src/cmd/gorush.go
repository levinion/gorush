package cmd

import(
	"github.com/urfave/cli"
	"github.com/levinion/gorush/src/serve"
	"os"
	"log"
)

func Run(){
	app:=cli.NewApp()
	app.Name="gorush"
	app.Usage="a simple and quick blog easy to build and diy"
	app.Action=func(c *cli.Context){
		serve.Run()
	}

	app.Commands=cli.Commands{
		{
			Name: "serve",
			Usage: "Starting serving the net",
			Aliases: []string{"s"},
			Action: func(c *cli.Context){
				serve.Run()
			},
		},
	}

	err:=app.Run(os.Args)
	if err!=nil{
		log.Println(err)
	}
}