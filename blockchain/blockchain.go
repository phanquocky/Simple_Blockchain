package blockchain

import (
	"blockchain_go/block"
)

type Blockchain struct {
	Blocks       []*block.Block
	MapBlockHash map[block.Hash]*block.Block
}

func NewBlockchain() *Blockchain {
	blockchain := Blockchain{[]*block.Block{block.NewGenesisBlock()}, map[block.Hash]*block.Block{}}
	genesisBlock := blockchain.Blocks[0]
	blockchain.MapBlockHash[genesisBlock.BlockHash] = genesisBlock

	return &blockchain
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.BlockHash)
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.MapBlockHash[newBlock.BlockHash] = newBlock
}

func (bc *Blockchain) FindBlockByHash(hash block.Hash) *block.Block {
	block, exists := bc.MapBlockHash[hash]

	if exists {
		return block
	} else {
		return nil
	}
}
