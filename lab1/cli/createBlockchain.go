package cli

import (
	"blockchain_go/blockchain"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	bc := blockchain.CreateBlockchain(address)
	defer bc.DB.Close()

	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	log.Println("Blockchain create success!")
}
