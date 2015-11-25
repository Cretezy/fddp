package main
import (
	"fmt"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"os"
	"time"
)

var quiet bool
var jsonFile string

func main() {
	// Calculate running time, useful to know perfomance
	start := time.Now()

	app := cli.NewApp()
	app.Name = "fddp"
	app.Usage = "Facebook Downloaded Data Processor"

	app.Commands = []cli.Command{
		cli.Command{
			Name: "convert",
			Description: "convert messages from one data holder to another (i.e.: html to json)",
			// fddp convert --fromHtml messages.htm --toJson messages.json
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "fromHtml, fH",
				},
				cli.StringFlag{
					Name: "fromJson, fJ",
				},
				cli.StringFlag{
					Name: "toJson, tJ",
				},
			},
			Action: convert,
		},
	}

	app.Run(os.Args)

	// Calculate time
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func run(c *cli.Context) {

}

// Get if 2 slices of people's name match, order doesn't matter
// Not by me
func matchingPersons(persons1 []string, persons2 []string) bool {
	diffStr := []string{}
	m := map[string]int{}

	for _, s1Val := range persons1 {
		m[s1Val] = 1
	}
	for _, s2Val := range persons2 {
		m[s2Val] = m[s2Val] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return len(diffStr) == 0
}

func contains(slice []string, contain string) bool {
	for _, element := range slice {
		if element == contain {
			return true
		}
	}
	return false
}
