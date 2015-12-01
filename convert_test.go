package main

import "testing"

func TestConversion(t *testing.T) {
	dataHTML := FromHTML(GetFileContent("samples/sample.html"))
	dataJSON := FromJSON(GetFileContent("samples/sample.json"))

	if !CheckIndentical(dataHTML, dataJSON, true) {
		t.Error("Data not matching")
	}
}
