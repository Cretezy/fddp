package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"time"
)

// Runs fddp
var start time.Time

func main() {
	// Calculate running time from start to bottom, useful to know perfomance

	app := cli.NewApp()

	app.Name = "fddp"
	app.Usage = "Facebook Downloaded Data Processor"
	app.Version = "0.0.2"

	app.Commands = Commands()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.Bool("verbose") {
			start = time.Now()
		}

		return nil
	}

	app.After = func(c *cli.Context) error {
		if c.Bool("verbose") {
			elapsed := time.Since(start)
			fmt.Println("\nTook", elapsed)
		}

		return nil
	}

	app.Run(os.Args)
}
