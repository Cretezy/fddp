package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"sort"
	"strings"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:      "list",
		Usage:     "list various stats",
		ArgsUsage: "input.json",
		Action:    ListAction,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "count, c",
				Value: 50,
				Usage: "how much to list",
			},
		},
	}
}

func ListAction(c *cli.Context) {
	// Figure out input
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	jsonFile := c.Args()[0]

	data := FromJSON(GetFileContent(jsonFile))

	count := c.Int("count")

	threads := data.Threads

	if count > len(threads) {
		count = len(threads)
	}

	// Sort by message count
	sort.Sort(ByMessage(threads))

	// Reverse (top = more)
	for i, j := 0, len(threads)-1; i < j; i, j = i+1, j-1 {
		threads[i], threads[j] = threads[j], threads[i]
	}

	threads = threads[:count]

	for _, thread := range threads {
		fmt.Println(strings.Join(DeleteElementFromSlice(thread.Persons, data.WhoAmI), ", ")+" > Messages:", thread.CountMessages(), "Words:", thread.CountWords())
	}
}
