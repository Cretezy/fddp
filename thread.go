package main

type Thread struct {
	messages []Message
	persons  []string
}

type ByMessage []Thread
func (a ByMessage) Len() int           { return len(a) }
func (a ByMessage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMessage) Less(i, j int) bool { return len(a[i].messages) < len(a[j].messages)}

type Message struct {
	sender string
	text   string
}