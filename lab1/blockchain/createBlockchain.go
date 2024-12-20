package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

func CreateBlockchain(address string) *Blockchain {
	var tip []byte

	// remove "my.db" file
	err := os.Remove(DB_FILE)
	if err != nil {
		log.Println("Cannot remove file, ", DB_FILE, err)
	}

	db, err := bolt.Open(DB_FILE, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Println("Cannot Open Database!")
		return nil
	}

	err = db.Update(func(dbtx *bolt.Tx) error {
		bucket := dbtx.Bucket([]byte(BLOCK_BUCKET))

		if bucket == nil {
			coinbaseTx := tx.NewCoinbaseTX(address, GENESIS_COINBASE_DATA)
			genesisBlock := block.NewGenesisBlock(coinbaseTx, DEFAULT_DIFFICULTY)

			// update db
			b, err := dbtx.CreateBucket([]byte(BLOCK_BUCKET))
			if err != nil {
				log.Println("Cannot create Bucket ", BLOCK_BUCKET, err)
				return err
			}

			err = b.Put(genesisBlock.BlockHash[:], genesisBlock.Serialize())
			if err != nil {
				log.Println("Cannot put genesisblock to database, ", err)
				return err
			}

			err = b.Put([]byte("l"), genesisBlock.BlockHash[:])
			if err != nil {
				return err
			}

			tip = genesisBlock.BlockHash[:]
		} else {
			tip = bucket.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Println("Cannot create blockchain, ", err)
		return nil
	}

	return &Blockchain{
		Tip:        tip,
		DB:         db,
		Difficulty: DEFAULT_DIFFICULTY,
	}
}
