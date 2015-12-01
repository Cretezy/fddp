package main

import "strings"

type FacebookData struct {
	WhoAmI  string   `json:"whoami"`
	Threads []Thread `json:"threads"`
}

func (data FacebookData) CountWords() int {
	var words int = 0

	for _, thread := range data.Threads {
		for _, message := range thread.Messages {
			words += len(strings.Split(message.Text, " "))
		}
	}

	return words
}

func (data FacebookData) CountThreads() int {
	return len(data.Threads)
}

func (data FacebookData) CountMessages() int {
	var messages int = 0

	for _, thread := range data.Threads {
		messages += len(thread.Messages)
	}

	return messages
}
