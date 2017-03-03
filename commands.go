package main

import (
	"gopkg.in/urfave/cli.v1"
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
