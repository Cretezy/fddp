package main

import (
	"fmt"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func CountCommand() cli.Command {
	return cli.Command{
		Name:        "count",
		Description: "count various stats (messages, words, etc)",
		Usage:       "input.json",
		Action:      CountAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "threads, t",
				Usage: "count threads",
			},
			cli.BoolFlag{
				Name:  "messages, m",
				Usage: "count messages",
			},
			cli.BoolFlag{
				Name:  "words, w",
				Usage: "count words",
			},
		},
	}
}

func CountAction(c *cli.Context) {
	// Figure out input
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	jsonFile := c.Args()[0]

	data := FromJSON(GetFileContent(jsonFile))
	hasSomeOutput := false

	displayCount("threads", data.CountMessages(), c, &hasSomeOutput)
	displayCount("messages", data.CountMessages(), c, &hasSomeOutput)
	displayCount("words", data.CountWords(), c, &hasSomeOutput)

	if !hasSomeOutput {
		fmt.Println("You must include a flag to display an output")
		fmt.Println()
		cli.ShowCommandHelp(c, c.Command.Name)
	}
}

func displayCount(stat string, statCount int, c *cli.Context, hasSomeOutput *bool) {
	if c.Bool(stat) {
		*hasSomeOutput = true
		fmt.Println("Data set has", statCount, stat)
	}
}

