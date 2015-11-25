package main
import (
"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
"strings"
	"fmt"
	"io/ioutil"
)

func convert(c *cli.Context) {
	// figure out input
	var threads []Thread
	if(c.String("fromHtml") != ""){
		threads = FromHtml(GetFileContent(c.String("fromHtml")))
	}else if(c.String("fromJson") != "") {
		threads = FromJson(GetFileContent(c.String("fromJson")))
	}else{
		threads = make([]Thread, 0)
	}

	var words int = 0
	var messages int = 0

	for _, thread := range threads {
		messages += len(thread.Messages)
		for _, message := range thread.Messages {
			words += len(strings.Split(message.Text, " "))
		}
	}
	if(messages == 0){messages++}

	fmt.Println("Words total", words, "Total messages", messages, "Average words/message", words/messages)

}

func GetFileContent(file string) string {
	// Open the file
	reader, err := ioutil.ReadFile(file)
	check(err)

	return string(reader)
}


