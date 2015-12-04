package main

import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/CraftThatBlock/fddp/server"
)

func Commands() []cli.Command {
	return []cli.Command{
		ConvertCommand(),
		server.ServerCommand(),
		CountCommand(),
		CompareCommand(),
		ListCommand(),
	}
}
