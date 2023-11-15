package cli

import (
	"blockchain_go/blockchain"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	bc := blockchain.CreateBlockchain(address)
	bc.DB.Close()

	log.Println("Blockchain create success!")
}
