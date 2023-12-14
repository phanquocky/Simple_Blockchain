package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"blockchain_go/txscript"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	DEFAULT_DIFFICULTY uint32 = binary.BigEndian.Uint32([]byte{33, 0, 0, 100}) // with this difficulty it have 16 bit 0

	SIZE_DIFFICULTY_RESET  = 2016 // reset difficulty after 2016 block
	AVERAGE_TIME_FOR_BLOCK = 10   // 10 minute for one block
	BLOCK_BUCKET           = "blockbucket"
	GENESIS_COINBASE_DATA  = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	DB_FILE                = "my.db"
)

type Blockchain struct {
	// Tip: the hash of last block in blockchain, it like identify of a blockchain
	Tip        []byte
	DB         *bolt.DB
	Difficulty uint32
}

func ReadBlockchain() *Blockchain {
	var (
		tip []byte
	)

	db, err := bolt.Open(DB_FILE, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Println("Cannot open database, ", err)
		return nil
	}

	err = db.Update(func(dbtx *bolt.Tx) error {
		bucket := dbtx.Bucket([]byte(BLOCK_BUCKET))
		if bucket == nil {
			log.Println("Bucket doesn't exists!, ", err)
			return err
		}

		tip = bucket.Get([]byte("l"))

		return nil
	})
	if err != nil {
		return nil
	}

	return &Blockchain{
		Tip:        tip,
		DB:         db,
		Difficulty: DEFAULT_DIFFICULTY,
	}
}

func (bc *Blockchain) AddBlock(transactions []*tx.Transaction) *block.Block {
	// verify transactions
	for _, transaction := range transactions {
		var prevOutputFetcher = make([]txscript.PrevOutputFetcher, len(transaction.Vin))

		for idx, in := range transaction.Vin {
			fmt.Printf("txVin: %x \n", in.Txid)
			prevTx, err := bc.GetRawTransaction(in.Txid)
			if err != nil {
				log.Println("cannot get rawtransaction, ", err)
				return nil
			}

			prevOutputFetcher[idx] = txscript.PrevOutputFetcher{
				PkHash: prevTx.Vout[in.OutIdx].PubKeyHash,
				Amt:    int64(prevTx.Vout[in.OutIdx].Value),
			}
		}

		if !txscript.VerifyTransaction(transaction, prevOutputFetcher) {
			log.Println("Verify transaction faile, ", transaction.ID)
			return nil
		}
	}

	log.Println("Verify all of transactions success!")

	var lastHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_BUCKET))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Println("Error cannot reading from database, ", err)
		return nil
	}

	newBlock := block.NewBlock(transactions, block.Hash(lastHash), bc.Difficulty)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_BUCKET))
		err := b.Put(newBlock.BlockHash[:], newBlock.Serialize())
		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), newBlock.BlockHash[:])
		if err != nil {
			return err
		}
		bc.Tip = newBlock.BlockHash[:]

		return nil
	})
	if err != nil {
		log.Println("Error cannot update database when add block, ", err)
		return nil
	}

	return newBlock

}
