package blockchain

import (
	"blockchain_go/block"
	"encoding/binary"
	"log"

	"github.com/boltdb/bolt"
)

var (
	DEFAULT_DIFFICULTY uint32 = binary.BigEndian.Uint32([]byte{33, 0, 0, 100}) // with this difficulty it have 16 bit 0

	SIZE_DIFFICULTY_RESET  = 2016 // reset difficulty after 2016 block
	AVERAGE_TIME_FOR_BLOCK = 10   // 10 minute for one block
	BLOCK_BUCKET           = "blockbucket"
)

type Blockchain struct {
	Tip        []byte
	DB         *bolt.DB
	Difficulty uint32
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Println("Cannot Open Database!")
		return nil
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		if b == nil {
			genesis := block.NewGenesisBlock(DEFAULT_DIFFICULTY)
			b, err := tx.CreateBucket([]byte(BLOCK_BUCKET))
			if err != nil {
				return err
			}
			err = b.Put(genesis.BlockHash[:], genesis.Serialize())
			err = b.Put([]byte("l"), genesis.BlockHash[:])
			if err != nil {
				return err
			}

			tip = genesis.BlockHash[:]
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db, DEFAULT_DIFFICULTY}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
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

	newBlock := block.NewBlock(data, block.Hash(lastHash), bc.Difficulty)

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
