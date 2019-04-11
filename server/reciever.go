package main

import (
	"io/ioutil"
	"log"
	"os"

	rsocket "github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "file reciever"
	app.Usage = "recieve file"
	app.Action = func(c *cli.Context) error {
		err := rsocket.Receive().
			Acceptor(func(setup payload.SetupPayload, sendingSocket rsocket.RSocket) rsocket.RSocket {
				fileName, _ := setup.MetadataUTF8()
				// bind responder
				return rsocket.NewAbstractSocket(
					rsocket.FireAndForget(
						func(elem payload.Payload) {
							writeFile(fileName, elem)
						}))
			}).
			Transport("127.0.0.1:8081").
			Serve()
		panic(err)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(name string, pl payload.Payload) {
	log.Println(name)
	err := ioutil.WriteFile(name, pl.Data(), 0644)
	if err != nil {
		panic(err)
	}
}
