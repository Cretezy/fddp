package main

type FacebookData struct {
	WhoAmI  string   `json:"whoami"`
	Threads []Thread `json:"threads"`
}

func (fbData FacebookData) CountWords() int {
	words := 0

	for _, thread := range fbData.Threads {
		words += thread.CountWords()
	}

	return words
}

func (fbData FacebookData) CountThreads() int {
	return len(fbData.Threads)
}

func (fbData FacebookData) CountMessages() int {
	messages := 0

	for _, thread := range fbData.Threads {
		messages += thread.CountMessages()
	}

	return messages
}

func (fbData FacebookData) CountWordsPerMessage() int {
	return fbData.CountWords() / fbData.CountMessages()
}
