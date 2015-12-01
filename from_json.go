package main

import (
	"encoding/json"
)

// Converts data from JSON to FacebookData
func FromJSON(jsonData string) FacebookData {
	// Get from JSON
	var data FacebookData
	err := json.Unmarshal([]byte(jsonData), &data)
	check(err)

	return data
}
