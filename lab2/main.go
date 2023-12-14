package main

import (
	"blockchain_go/lab2/cli"
	"blockchain_go/lab2/server"
	"log"
)

func main() {
	client, err := server.NewClient()

	if err != nil {
		log.Fatal("cannot create new Client, err ", err)
	}
	cli := cli.New(client)
	cli.Run()

}
