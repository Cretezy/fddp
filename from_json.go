package main
import (
	"encoding/json"
)

func FromJson(jsonData string) []Thread {
	// Get from json
	var threads []Thread
	err := json.Unmarshal([]byte(jsonData), &threads)
	check(err)

	return threads
}
