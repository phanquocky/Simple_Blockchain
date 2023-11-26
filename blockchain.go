package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

var BLOCKCHAIN_SAVER = "blockchain.dat"

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) Print() {
	for index, block := range bc.Blocks {
		fmt.Printf("Block index %d\n", index)
		block.Print()
	}
}

func BuildBlockchain() *Blockchain {
	// Create the initial blockchain with genesis block
	genesisBlock := CreateGenesisBlock()
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

func CreateGenesisBlock() *Block {
	// Create the initial block (genesis block)
	transaction := &Transaction{Data: []byte("Genesis Transaction")}
	block := NewBlock([]*Transaction{transaction}, []byte{})
	return &block
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	// Create a new block and add it to the chain
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, &newBlock)
}

func (bc *Blockchain) Validate() bool {
	for i := len(bc.Blocks) - 1; i > 0; i-- {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Verify the hash of the current block
		if !bytes.Equal(currentBlock.Hash, currentBlock.HashABLock()) {
			return false
		}

		// Verify the previous block's hash
		if !bytes.Equal(previousBlock.Hash, previousBlock.HashABLock()) {
			return false
		}
	}
	return true
}

func (bc *Blockchain) SaveBlockchain() error {
	file, err := os.Create(BLOCKCHAIN_SAVER)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(bc)
	return err
}

func LoadBlockchain() (*Blockchain, error) {
	var bc Blockchain

	file, err := os.Open(BLOCKCHAIN_SAVER)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&bc)
	return &bc, err
}
