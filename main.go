package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/cli"
	"os"
	"strings"
	"time"
	"sort"
)
func main() {
	// Calculate running time, useful to know perfomance
	start := time.Now()

	app := cli.NewApp()
	app.Name = "fddp"
	app.Usage = "Facebook Downloaded Data Processor"
	app.Action = run

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

	reader, err := os.Open(c.Args().First())
	doc, err := goquery.NewDocumentFromReader(reader)
	check(err)


	threads := make([]Thread, 0)

	whoami := doc.Find("h1").Text()

	doc.Find(".thread").Each(func(threadId int, threadSelector *goquery.Selection) {
		// People in this thread
		persons := strings.Split(strings.TrimSpace(threadSelector.Nodes[0].FirstChild.Data), ", ")

		// List of message in this thread
		messages := make([]Message, 0)

		sender := ""

		threadSelector.Children().Each(func(someId int, someSelector *goquery.Selection) {
			nodeType := someSelector.Nodes[0].Data

			if nodeType == "div" {
				sender = someSelector.Find(".message_header").Find(".user").Text()
			} else if nodeType == "p" {
				message := Message{sender: sender, text: someSelector.Text()}
				messages = append(messages, message)
			}
		})
		addToThread(&threads, persons, messages)
	})

	// Sort by highest messages
	sort.Sort(ByMessage(threads))
	// Reverse (top = more)
	for i, j := 0, len(threads) - 1; i < j; i, j = i + 1, j - 1 {
		threads[i], threads[j] = threads[j], threads[i]
	}

	// Print amount of messages
	fmt.Println("You are", whoami)

	// Print list of conversations and how much messages
	for _, thread := range threads {
		fmt.Println("The conversation between", strings.Join(thread.persons, " and "), "has", len(thread.messages), "messages")
	}
}

func addToThread(threads *[]Thread, persons []string, messages []Message) {
	newThreads := make([]Thread, 0)
	var newThread Thread
	for _, thread := range *threads {
		if matchingPersons(persons, thread.persons) {
			newThread = thread
		} else {
			newThreads = append(newThreads, thread)
		}
	}
	if len(newThread.persons) < 1 {
		newThread = Thread{persons: persons, messages: make([]Message, 0)}
	}

	newThread.messages = append(newThread.messages, messages...)
	*threads = append(newThreads, newThread)
}

// Get if 2 slices of people's name match, order doesn't matter
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
