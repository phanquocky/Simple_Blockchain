package cli

import (
	"blockchain_go/blockchain"
	"blockchain_go/tx"
	"blockchain_go/txscript"
	"blockchain_go/util"
	"blockchain_go/wallet"
	"log"
)

func (cli *CLI) transfer(from, to string, amount uint) {
	var (
		bc                 = blockchain.ReadBlockchain()
		utxos              = bc.FindUTXO(from)
		transferTx         = tx.NewTransaction()
		txAmount           = 0
		changeAmount       = 0
		fromPubkeyHash     = util.GetPubkeyHash(from)
		toPubkeyHash       = util.GetPubkeyHash(to)
		prevOutputFetchers = make([]txscript.PrevOutputFetcher, 0)
	)

	wallets, err := wallet.NewWallets()
	if err != nil {
		log.Println("cannot read wallet from file, ", err)
		return
	}

	wallet := wallets.Wallets[from]

	// add inputs
UTXO:
	for utx, idxs := range utxos {
		for _, idx := range idxs {
			txAmount += utx.Vout[idx].Value
			txInput := tx.NewTxInput(utx.ID, idx, nil, wallet.PublicKey)
			transferTx.AddTxInput(txInput)

			// add prevoutput fetcher
			prevOutputFetchers = append(prevOutputFetchers,
				txscript.PrevOutputFetcher{
					PkHash: utx.Vout[idx].PubKeyHash,
					Amt:    int64(utx.Vout[idx].Value),
				})

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
	//add outputs
	desOutput := tx.NewTxOutput(int(amount), toPubkeyHash)
	changeOutput := tx.NewTxOutput(changeAmount, fromPubkeyHash)
	transferTx.AddTxOutput(desOutput)
	transferTx.AddTxOutput(changeOutput)

	// set TxId
	transferTx.SetID()

	// sign transaction
	txscript.SignRawTransaction(transferTx, wallet.PrivateKey, prevOutputFetchers)
	// mine block
	bc.AddBlock([]*tx.Transaction{transferTx})
	log.Printf("Transfer success, from: %s, to: %s, amount: %d !\n", from, to, amount)

}
