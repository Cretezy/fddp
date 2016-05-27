package main

import (
	"github.com/codegangsta/cli"
)

func Commands() []cli.Command {
	return []cli.Command{
		ConvertCommand(),
		ServerCommand(),
		CountCommand(),
		CompareCommand(),
		ListCommand(),
	}
}
