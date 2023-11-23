package cli

import (
	"blockchain_go/blockchain"
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address string) {
	bc := blockchain.ReadBlockchain()
	defer bc.DB.Close()

	balance := 0
	utxoSet := blockchain.NewUTXOSet(bc)
	UTXOs := utxoSet.FindUTXOByAddress(address)
	fmt.Println("utxoSet: ", UTXOs)
	fmt.Println("Get Balance")

	for txID, outs := range UTXOs {
		for _, out := range outs.Outputs {
			balance += out.Value
			log.Printf("UTXO: txId: %x, outputindx: %d , value: %d \n", txID, out.OutputIdx, out.Value)
		}
	}

	log.Printf("Balance of '%s': %d\n", address, balance)
}
