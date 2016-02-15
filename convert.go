package main

import (
	"fmt"
	"github.com/Cretezy/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"io/ioutil"
)

func ConvertCommand() cli.Command {
	return cli.Command{
		Name:      "convert",
		Usage:     "convert messages from HTML to JSON",
		ArgsUsage: "input.html output.json",
		Action:    ConvertAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "indent, i",
				Usage: "indent output of JSON file",
			},
		},
	}
}

func ConvertAction(c *cli.Context) {
	// Figure out input
	if len(c.Args()) < 2 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	htmlFile := c.Args()[0]
	jsonFile := c.Args()[1]

	// Get FacebookData from Html file
	fbData := FromHTML(GetFileContent(htmlFile))

	jsonData := ToJSON(fbData, c.Bool("indent"))

	// Write it out
	err := ioutil.WriteFile(jsonFile, []byte(jsonData), 0644)
	check(err)

	fmt.Println("Done converting", htmlFile, "(HTML) to", jsonFile, "(JSON)")
}
