package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"bytes"
	"errors"
	"fmt"
)

func (bc *Blockchain) GetRawTransaction(ID []byte) (*tx.Transaction, error) {
	bci := bc.Iterator()

	for {
		bl := bci.Next()

		for _, tx := range bl.Transactions {
			fmt.Printf("%x %x \n", tx.ID, ID)
			if bytes.Equal(tx.ID, ID) {
				return tx, nil
			}
		}

		if bl.PrevBlockHash == block.Hash(block.Hash{}) {
			break
		}
	}

	return nil, errors.New("transaction is not found")
}
