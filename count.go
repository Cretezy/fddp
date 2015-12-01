package main

import (
	"fmt"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"strings"
)

func CountCommand() cli.Command {
	return cli.Command{
		Name:        "count",
		Description: "count various stats (messages, words, etc)",
		Usage:       "input.json",
		Action:      Count,
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

func Count(c *cli.Context) {
	// Figure out input
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	jsonFile := c.Args()[0]

	data := FromJSON(GetFileContent(jsonFile))
	hasSomeOutput := false

	displayCount("threads", CountThreads(data), c, &hasSomeOutput)
	displayCount("messages", CountMessages(data), c, &hasSomeOutput)
	displayCount("words", CountWords(data), c, &hasSomeOutput)

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

func CountWords(data FacebookData) int {
	var words int = 0

	for _, thread := range data.Threads {
		for _, message := range thread.Messages {
			words += len(strings.Split(message.Text, " "))
		}
	}

	return words
}

func CountThreads(data FacebookData) int {
	return len(data.Threads)
}

func CountMessages(data FacebookData) int {
	var messages int = 0

	for _, thread := range data.Threads {
		messages += len(thread.Messages)
	}

	return messages
}
