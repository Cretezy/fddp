package main
import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"strings"
	"time"
	"sort"
)
/*
Note to Facebook:
Why do you organize the messages.htm like that?
It could be make simpler and more efficient, just hire me! ;)

Note:
This reads from Facebook's message.htm and converts in a FacebookData.
You can see a sample under samples/sample.htm
*/

func FromHtml(html string) FacebookData {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	check(err)

	threads := make([]Thread, 0)

	whoami := doc.Find("h1").Text()

	// Format needs to equal: Mon Jan 2 15:04:05 MST 2006
	format := "Monday, 2 January 2006 at 15:04 MST"

	doc.Find(".thread").Each(func(threadId int, threadSelector *goquery.Selection) {
		// People in this thread
		persons := strings.Split(strings.TrimSpace(threadSelector.Nodes[0].FirstChild.Data), ", ")
		// List of message in this thread
		messages := make([]Message, 0)

		sender := ""
		sentTime := time.Now()

		threadSelector.Children().Each(func(someId int, someSelector *goquery.Selection) {
			nodeType := someSelector.Nodes[0].Data

			if nodeType == "div" {
				sender = someSelector.Find(".message_header").Find(".user").Text()
				sentTime, err = time.Parse(format, someSelector.Find(".message_header").Find(".meta").Text())
				check(err)
			} else if nodeType == "p" {
				message := Message{Sender: sender, Text: someSelector.Text(), Time: sentTime}
				messages = append(messages, message)
			}
		})
		threads = append(threads, Thread{Persons: persons, Messages: messages})
	})

	threads = FixThreads(threads)

	// Sort by highest messages
	sort.Sort(ByMessage(threads))
	// Reverse (top = more)
	for i, j := 0, len(threads) - 1; i < j; i, j = i + 1, j - 1 {
		threads[i], threads[j] = threads[j], threads[i]
	}

	return FacebookData{WhoAmI: whoami, Threads: threads}
}

// Removes duplicate threads from Html format
func FixThreads(threads []Thread) []Thread {
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