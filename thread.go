package main

import (
	"strings"
	"time"
)

type Thread struct {
	Messages []Message `json:"messages"`
	Persons  []string  `json:"persons"`
}

func (thread Thread) CountMessages() int {
	return len(thread.Messages)
}
func (thread Thread) CountWords() int {
	words := 0

	for _, message := range thread.Messages {
		words += len(strings.Split(message.Text, " "))
	}

	return words
}

type ByMessage []Thread

func (a ByMessage) Len() int           { return len(a) }
func (a ByMessage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMessage) Less(i, j int) bool { return len(a[i].Messages) < len(a[j].Messages) }

type Message struct {
	Sender string    `json:"sender"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"`
}
