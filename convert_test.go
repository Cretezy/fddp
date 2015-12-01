package main

import "testing"

func TestConversion(t *testing.T) {
	dataHTML := FromHTML(GetFileContent("samples/sample.html"))
	dataJSON := FromJSON(GetFileContent("samples/sample.json"))

	if dataHTML.WhoAmI != dataJSON.WhoAmI ||
		len(dataHTML.Threads) != len(dataJSON.Threads) {
		t.Error("Data not matching")
	}
}
