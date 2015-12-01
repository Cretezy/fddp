package main

import "github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"

func Commands() []cli.Command {
	return []cli.Command{
		ConvertCommand(),
		CountCommand(),
		CompareCommand(),
	}
}
