package main
import (
	"encoding/json"
)

func FromJson(jsonData string) FacebookData {
	// Get from json
	var data FacebookData
	err := json.Unmarshal([]byte(jsonData), &data)
	check(err)

	return data
}
