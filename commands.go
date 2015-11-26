package main
import "github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"

func Commands() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "convert",
			Description: "convert messages from Html to Json",
			Usage: "input.html output.json",
			Action: Convert,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "indent, i",
					Usage: "indent output of Json file",
				},
			},
		},
	}
}
