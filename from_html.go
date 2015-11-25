package main
import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

func FromHtml(html string) []Thread {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	check(err)

	threads := make([]Thread, 0)

	// whoami := doc.Find("h1").Text()

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
	return fixThreads(threads)
}