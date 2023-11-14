package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"encoding/binary"
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

func (bc *Blockchain) AddBlock(transactions []*tx.Transaction) {
	var lastHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_BUCKET))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Println("Error cannot reading from database, ", err)
		return
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
		return
	}

	// bc.ResetDifficulty()
}

// // this func will reset diff after 2016 block
// func (bc *Blockchain) ResetDifficulty() {
// 	if len(bc.Blocks)%SIZE_DIFFICULTY_RESET == 0 && len(bc.Blocks) > 0 {
// 		actualTime := bc.Blocks[len(bc.Blocks)-1].Timestamp - bc.Blocks[len(bc.Blocks)-SIZE_DIFFICULTY_RESET].Timestamp
// 		fmt.Println("Before reset: ", bc.Difficulty)
// 		fmt.Println("Value: ", float64((int64(SIZE_DIFFICULTY_RESET))*int64(time.Duration(AVERAGE_TIME_FOR_BLOCK*int(time.Second)).Seconds())))

// 		coeff := float64(actualTime / ((int64(SIZE_DIFFICULTY_RESET)) * int64(time.Duration(AVERAGE_TIME_FOR_BLOCK*int(time.Second)).Seconds())))
// 		fmt.Println("Coeeff: ", coeff)
// 		bc.Difficulty = uint32(float64(bc.Difficulty) * coeff)
// 		fmt.Println("After reset: ", bc.Difficulty)
// 	}
// }
