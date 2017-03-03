package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
)

func CompareCommand() cli.Command {
	return cli.Command{
		Name:      "compare",
		Usage:     "compare 2 data set",
		ArgsUsage: "input1.json input2.json",
		Action:    CompareAction,
	}
}

func CompareAction(c *cli.Context) {
	// Figure out input
	if len(c.Args()) < 2 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	file1 := c.Args()[0]
	file2 := c.Args()[1]

	data1 := FromJSON(GetFileContent(file1))
	data2 := FromJSON(GetFileContent(file2))

	CheckIndentical(data1, data2, true)
}

func CheckIndentical(data1 FacebookData, data2 FacebookData, print bool) bool {
	indentical := true

	if data1.WhoAmI != data2.WhoAmI {
		indentical = false
		if print {
			fmt.Println("Originator not indentical (", data1.WhoAmI, "vs", data2.WhoAmI, ")")
		}
	}

	if data1.CountThreads() != data2.CountThreads() {
		indentical = false
		if print {
			fmt.Println("Number of threads not indentical(", data1.CountThreads(), "vs", data2.CountThreads(), ")")
		}
	}

	if data1.CountMessages() != data2.CountMessages() {
		indentical = false
		if print {
			fmt.Println("Number of messages not indentical(", data1.CountMessages(), "vs", data2.CountMessages(), ")")
		}
	}

	if data1.CountWords() != data2.CountWords() {
		indentical = false
		if print {
			fmt.Println("Number of words not indentical(", data1.CountWords(), "vs", data2.CountWords(), ")")
		}
	}

	if print {
		if indentical {
			fmt.Println("Data sets indentical!")
		} else {
			fmt.Println("Data sets are not indentical!")
		}
	}

	return indentical
}
