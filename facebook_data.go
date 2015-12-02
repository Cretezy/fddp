package main

type FacebookData struct {
	WhoAmI  string   `json:"whoami"`
	Threads []Thread `json:"threads"`
}

func (data FacebookData) CountWords() int {
	words := 0

	for _, thread := range data.Threads {
		words += thread.CountWords()
	}

	return words
}

func (data FacebookData) CountThreads() int {
	return len(data.Threads)
}

func (data FacebookData) CountMessages() int {
	messages := 0

	for _, thread := range data.Threads {
		messages += thread.CountMessages()
	}

	return messages
}
