package server
import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/go-martini/martini"
)

func ServerCommand() cli.Command{
	return cli.Command{
		Name: "server",
		Usage: "run web app",
		Description: "Use enviroment variable PORT to set which port to listen to.",
		ArgsUsage: "",
		Action: ServerAction,
		Flags: []cli.Flag{

		},
	}
}

func ServerAction(c *cli.Context){
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}