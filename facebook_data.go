package main

type FacebookData struct {
	WhoAmI  string   `json:"whoami"`
	Threads []Thread `json:"threads"`
}
