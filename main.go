package main
import (
	"encoding/json"
	"fmt"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"sort"
	"strings"
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
	app.Action = func(c *cli.Context) {
		quiet = c.Bool("quiet")
		jsonFile = c.String("json")
		run(c)
	}

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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "json",
			Value: "",
			Usage: "where to output json",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "show no output",
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
	if len(c.Args()) < 1 {
		fmt.Println("Must supply message file. Example: samples/sample.html")
		return
	}

	threads := FromHtml(c.Args().First())

	// Sort by highest messages
	sort.Sort(ByMessage(threads))
	// Reverse (top = more)
	for i, j := 0, len(threads)-1; i < j; i, j = i+1, j-1 {
		threads[i], threads[j] = threads[j], threads[i]
	}

	// Print amount of messages
	if !quiet {
		//fmt.Println("You are", whoami)
		fmt.Println("You have messaged", len(threads), "people")

		// Print list of conversations and how much messages
		for _, thread := range threads {
			fmt.Println("The conversation between", strings.Join(thread.Persons, " and "), "has", len(thread.Messages), "messages")
		}
	}
	if jsonFile != "" {
		json, err := json.Marshal(threads)
		check(err)
		err = ioutil.WriteFile(jsonFile, json, 0644)
		check(err)
	}


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

func fixThreads(threads []Thread) []Thread {
	newThreads := make([]Thread, 0)
	persons := make([][]string, 0)

	for _, thread := range threads {
		skip := false
		for _, personCheck := range persons {
			if (matchingPersons(personCheck, thread.Persons)) {
				skip = true
				break
			}
		}
		if (skip) {
			continue
		}
		persons = append(persons, thread.Persons)

		newThread := Thread{Persons:thread.Persons, Messages:make([]Message, 0)}

		for _, otherThread := range threads {
			if (matchingPersons(thread.Persons, otherThread.Persons)) {
				newThread.Messages = append(newThread.Messages, otherThread.Messages...)
			}
		}
		newThreads = append(newThreads, newThread)
	}

	return newThreads
}