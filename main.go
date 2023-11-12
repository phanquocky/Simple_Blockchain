package main

import (
	"blockchain_go/blockchain"
	"blockchain_go/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.DB.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
