package main

import (
	"fmt"
	"log"
	"os"

	rsocket "github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "file sender"
	app.Usage = "send file"

	app.Commands = []cli.Command{
		{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "select file",
			Action: func(c *cli.Context) error {
				fmt.Println("transfer file: ", c.Args().First())
				upload(c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func upload(filename string) error {

	client, err := rsocket.Connect().
		SetupPayload(payload.NewString("file", filename)).
		Transport("127.0.0.1:8081").
		Start()
	defer client.Close()
	filepayload, err := payload.NewFile(filename, nil)
	if err != nil {
		panic(err)
	}
	// Send request
	client.FireAndForget(filepayload)
	return nil
}
