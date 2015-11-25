package main
import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func convert(c *cli.Context) {
	// Check that a data input is given
	if (c.String("fromHtml") != "" && c.String("fromJson") != "") {
		// Both are given
		fmt.Println("You cannot give 2 sources of input.")
		return
	}else if(c.String("fromHtml") == "" && c.String("fromJson") == ""){
		// None are given
		fmt.Println("You must give a input file using --fromHtml or --fromJson")
		return
	}
	
	// Figure out input
	var data FacebookData
	if (c.String("fromHtml") != "") {
		data = FromHtml(GetFileContent(c.String("fromHtml")))
	}else if (c.String("fromJson") != "") {
		data = FromJson(GetFileContent(c.String("fromJson")))
	}

	if !quiet {
		ShowStats(&data)
	}

	if (c.String("toJson") != "") {
		jsonData, err := json.Marshal(data)
		check(err)
		err = ioutil.WriteFile(c.String("toJson"), jsonData, 0644)
		check(err)
	}
}

func GetFileContent(file string) string {
	// Open the file
	reader, err := ioutil.ReadFile(file)
	check(err)

	return string(reader)
}

func ShowStats(data *FacebookData) {
	fmt.Println("Total messages", data.CountMessages())
}
