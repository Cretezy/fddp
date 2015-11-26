package main
import (
	"fmt"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"os"
	"time"
)

func main() {
	// Calculate running time, useful to know perfomance
	start := time.Now()

	app := cli.NewApp()
	app.Name = "fddp"
	app.Usage = "Facebook Downloaded Data Processor"

	app.Commands = []cli.Command{
		cli.Command{
			Name: "convert",
			Description: "convert messages from Html to Json",
			Usage: "{0} aaa",
			Action: Convert,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "indent, i",
					Usage: "indent output of Json file",
				},
			},
		},
	}

	app.Run(os.Args)

	// Calculate time
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}