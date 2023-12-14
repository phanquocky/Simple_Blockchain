package cli

import (
	"blockchain_go/block"
	"blockchain_go/blockchain"
	"fmt"
)

func (cli *CLI) printChain() {

	bc := blockchain.ReadBlockchain()
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		bl := bci.Next()

		fmt.Printf("Prev. hash: %x\n", bl.PrevBlockHash)
		fmt.Printf("Hash: %x\n", bl.BlockHash)
		fmt.Printf("Txs: %x\n", bl.HashTransactions())
		fmt.Printf("Nonce: %x\n", bl.Nonce)

		if block.Hash(block.Hash{}) == bl.PrevBlockHash {
			break
		}
	}
}
