package main
import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()

	run()

	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Message struct {
	sender string
	text   string
}
type Thread struct {
	messages []Message
	persons  []string
}

func AddMessage(messages Message, thread *Thread) {
	thread.messages = append(thread.messages, messages)
}

func run() {

	reader, err := os.Open("samples/messages.htm")
	doc, err := goquery.NewDocumentFromReader(reader)
	check(err)

	threads := make([]Thread, 0)

	//whoami := doc.Find("h1").Text()

	doc.Find(".thread").Each(func(threadId int, threadSelector *goquery.Selection) {
		persons := strings.Split(threadSelector.Nodes[0].FirstChild.Data, ", ")
		if (len(persons) < 2) {
			// TODO: Talking to self, may work

		}else if (len(persons) > 2) {
			// TODO: Group convo, may work

		}

		//findOthers(&persons, whoami)
		thread := findThread(persons, &threads)

		threadSelector.Children().Each(func(someId int, someSelector *goquery.Selection) {
			nodeType := someSelector.Nodes[0].Data
			class, _ := someSelector.Attr("class");
			who := ""
			if (nodeType == "div" && class == "message") {
				who = someSelector.Find(".message_header").Find(".user").Text()
			}else if (nodeType == "p") {
				message := Message{sender: who, text: someSelector.Text()}
				thread.messages = append(thread.messages, message)
			}
		})
	})
	// Print amount of messages
	for _, thread := range threads {
		fmt.Println(thread.persons, len(thread.messages))
	}
}
func findThread(persons []string, threads *[]Thread) *Thread {
	for _, thread := range *threads {
		if (matchingPersons(persons, thread.persons)) {

			fmt.Println("found", persons)
			return &thread
		}
	}

	fmt.Println("new", persons)
	thread := Thread{persons: persons, messages: make([]Message, 0)}
	*threads = append(*threads, thread)

	fmt.Println(len(*threads))

	return &thread
}

func matchingPersons(persons1 []string, persons2 []string) bool {

	return containsAll(persons1, persons2) && containsAll(persons2, persons1)
}
func containsAll(slice1 []string, slice2 []string) bool {
	for _, parent := range slice1 {
		has := false
		for _, child := range slice2 {
			if (parent == child) {
				has = true
				break
			}
		}
		if (!has) {
			return false
		}
	}
	return true;
}