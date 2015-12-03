package main

import (
	"encoding/json"
	"fmt"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
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
	data := FromHTML(GetFileContent(htmlFile))

	// Turn it into Json
	var jsonData []byte
	var err error
	if c.Bool("indent") {
		// Can enable indentation using --indent (or -i),
		// however it uses more space
		jsonData, err = json.MarshalIndent(data, "", "\t")
	} else {
		jsonData, err = json.Marshal(data)
	}
	check(err)

	// Write it out
	err = ioutil.WriteFile(jsonFile, jsonData, 0644)
	check(err)

	fmt.Println("Done converting", htmlFile, "(HTML) to", jsonFile, "(JSON)")
}
