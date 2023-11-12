package blockchain

import (
	"blockchain_go/block"
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.DB}

	return bci
}

func (i *BlockchainIterator) Next() *block.Block {
	var bl *block.Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_BUCKET))
		encodedBlock := b.Get(i.currentHash)
		bl = block.DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Println("Error cannot get prev block, ", err)
		return nil
	}

	i.currentHash = bl.PrevBlockHash[:]

	return bl
}
