package cli

import (
	"blockchain_go/blockchain"
	"blockchain_go/tx"
	"fmt"
	"log"
)

func (cli *CLI) transfer(from, to string, amount uint) {
	var (
		bc           = blockchain.ReadBlockchain()
		utxos        = bc.FindUTXO(from)
		transferTx   = tx.NewTransaction()
		txAmount     = 0
		changeAmount = 0
	)

	// add inputs
UTXO:
	for utx, idxs := range utxos {
		for _, idx := range idxs {
			txAmount += utx.Vout[idx].Value
			txInput := tx.NewTxInput(utx.ID, idx, from)
			transferTx.AddTxInput(txInput)

			if txAmount >= int(amount) {
				break UTXO
			}

		}
	}

	if txAmount < int(amount) {
		log.Println("You don't have enough coin to transfer")
		return
	}

	changeAmount = txAmount - int(amount)
	fmt.Printf("Change: %d, amount: %d \n", changeAmount, amount)
	//add outputs
	desOutput := tx.NewTxOutput(int(amount), to)
	changeOutput := tx.NewTxOutput(changeAmount, from)
	transferTx.AddTxOutput(desOutput)
	transferTx.AddTxOutput(changeOutput)

	// set TxId
	transferTx.SetID()
	// mine block
	bc.AddBlock([]*tx.Transaction{transferTx})
	log.Printf("Transfer success, from: %s, to: %s, amount: %d !\n", from, to, amount)

}
