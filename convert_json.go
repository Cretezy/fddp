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

func ToJSON(fbData FacebookData, indent bool) string {
	var jsonData []byte
	var err error
	if indent {
		// Can enable indentation using --indent (or -i),
		// however it uses more space
		jsonData, err = json.MarshalIndent(fbData, "", "\t")
	} else {
		jsonData, err = json.Marshal(fbData)
	}
	check(err)
	return string(jsonData)
}
