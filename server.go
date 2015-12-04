package main

import (
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/CraftThatBlock/fddp/Godeps/_workspace/src/github.com/go-martini/martini"
	"net/http"
	"fmt"
	"io/ioutil"

)

func ServerCommand() cli.Command {
	return cli.Command{
		Name:        "server",
		Usage:       "run web app",
		Description: "Use enviroment variable PORT to set which port to listen to.",
		ArgsUsage:   "",
		Action:      ServerAction,
		Flags:       []cli.Flag{},
	}
}

func ServerAction(c *cli.Context) {
	m := martini.Classic()
	//m.Get("/", func() string {
	//	return "Hello  world!"
	//})


	m.Post("/convert", func(r *http.Request) string {
		file, _, err := r.FormFile("messages")
		check(err)

		b, err := ioutil.ReadAll(file)
		check(err)
		fbData := FromHTML(string(b))

		fmt.Println(fbData)

		defer file.Close()
		return ToJSON(fbData, true)
	})
	m.Run()
}
