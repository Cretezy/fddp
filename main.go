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

func run() {

	reader, err := os.Open("samples/messages.htm")
	doc, err := goquery.NewDocumentFromReader(reader)
	check(err)

	threads := make([]Thread, 0)

	//whoami := doc.Find("h1").Text()

	doc.Find(".thread").Each(func(threadId int, threadSelector *goquery.Selection) {
		persons := strings.Split(strings.TrimSpace(threadSelector.Nodes[0].FirstChild.Data), ", ")
		if (len(persons) < 2) {
			// TODO: Talking to self, may work

		}else if (len(persons) > 2) {
			// TODO: Group convo, may work

		}

		//findOthers(&persons, whoami)
		messages:=make([]Message, 0)
		who := ""
		threadSelector.Children().Each(func(someId int, someSelector *goquery.Selection) {
			nodeType := someSelector.Nodes[0].Data
			class, _ := someSelector.Attr("class");

			if (nodeType == "div" && class == "message") {
				who = someSelector.Find(".message_header").Find(".user").Text()
			}else if (nodeType == "p") {
				message := Message{sender: who, text: someSelector.Text()}
				messages = append(messages, message)
			}
		})
		findThreadAndEdit(persons, &threads, messages)

	})

	threads = qsort(threads)
	// Print amount of messages
	for _, thread := range threads {
		fmt.Println(strings.Join(thread.persons, ", "), len(thread.messages))
	}
}

func findThreadAndEdit(persons []string, threads *[]Thread, messages []Message) {
	threadslol := make([]Thread, 0)
	var threadX *Thread
	for _, thread := range *threads {
		if (matchingPersons(persons, thread.persons)) {

			//fmt.Println("found", persons)
			threadX = &thread
		}else{
			threadslol =append( threadslol, thread)
		}
	}
	if(threadX==nil) {
		//fmt.Println("new", persons)
		threadX = &Thread{persons: persons, messages: make([]Message, 0)}
		*threads = append(*threads, *threadX)
	}

	threadX.messages = append(threadX.messages, messages...)
	threadslol = append( threadslol, *threadX)
	*threads = threadslol
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

func qsort(threads []Thread) []Thread {
	if len(threads) < 2 { return threads }

	left, right := 0, len(threads) - 1

	// Pick a pivot
	pivotIndex := (len(threads) + (len(threads)%2))/2

	// Move the pivot to the right
	threads[pivotIndex], threads[right] = threads[right], threads[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range threads {
		if len(threads[i].messages) < len(threads[right].messages) {
			threads[i], threads[left] = threads[left], threads[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	threads[left], threads[right] = threads[right], threads[left]

	// Go down the rabbit hole
	qsort(threads[:left])
	qsort(threads[left + 1:])


	return threads
}