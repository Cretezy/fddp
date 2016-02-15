package main

import (
	"fmt"
	"github.com/Cretezy/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"os"
	"time"
)

// Runs fddp
func main() {
	// Calculate running time from start to bottom, useful to know perfomance
	start := time.Now()

	app := cli.NewApp()

	app.Name = "fddp"
	app.Usage = "Facebook Downloaded Data Processor"
	app.Version = "0.0.2"

	app.Commands = Commands()

	app.Run(os.Args)

	// Calculate time
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println("Took", elapsed)
}
