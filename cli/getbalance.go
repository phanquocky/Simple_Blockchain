package cli

import (
	"blockchain_go/blockchain"
	"fmt"
)

func (cli *CLI) getBalance(address string) {
	bc := blockchain.ReadBlockchain()
	defer bc.DB.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for tx, outs := range UTXOs {
		for _, out := range outs {
			balance += tx.Vout[out].Value
			fmt.Printf("UTXO: txId: %x, outputindx: %d , value: %d \n", tx.ID, out, tx.Vout[out].Value)
		}
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
