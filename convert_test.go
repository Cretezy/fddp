package main

import "testing"

func TestConversion(t *testing.T) {
	dataHtml := FromHTML(GetFileContent("samples/sample.html"))
	dataJson := FromJSON(GetFileContent("samples/sample.json"))

	if dataHtml.WhoAmI != dataJson.WhoAmI {
		t.Error("WhoAmI not matching")
	}
	if dataHtml.CountThreads() != dataJson.CountThreads() {
		t.Error("thread number not matching")
	}

	if dataHtml.CountMessages() != dataJson.CountMessages() {
		t.Error("message number not matching")
	}
}
