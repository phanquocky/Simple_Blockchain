package blockchain

import (
	"blockchain_go/block"
	"encoding/binary"
)

var (
	DEFAULT_DIFFICULTY uint32 = binary.BigEndian.Uint32([]byte{33, 0, 0, 100}) // with this difficulty it have 16 bit 0

	SIZE_DIFFICULTY_RESET  = 2016 // reset difficulty after 2016 block
	AVERAGE_TIME_FOR_BLOCK = 10   // 10 minute for one block
)

type Blockchain struct {
	Blocks       []*block.Block
	MapBlockHash map[block.Hash]*block.Block
	Difficulty   uint32
}

func NewBlockchain() *Blockchain {
	blockchain :=
		Blockchain{
			[]*block.Block{block.NewGenesisBlock(DEFAULT_DIFFICULTY)},
			map[block.Hash]*block.Block{},
			DEFAULT_DIFFICULTY,
		}

	genesisBlock := blockchain.Blocks[0]
	blockchain.MapBlockHash[genesisBlock.BlockHash] = genesisBlock

	return &blockchain
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.BlockHash, bc.Difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.MapBlockHash[newBlock.BlockHash] = newBlock

	// bc.ResetDifficulty()
}

func (bc *Blockchain) FindBlockByHash(hash block.Hash) *block.Block {
	block, exists := bc.MapBlockHash[hash]

	if exists {
		return block
	} else {
		return nil
	}
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
