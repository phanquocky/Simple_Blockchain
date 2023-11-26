package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

type Transaction struct {
	Data []byte
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{Data: data}
}

type Block struct {
	Timestamp     int64
	MerkleTree    *MerkleTree
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) Print() {
	fmt.Printf("\tTimestamp:\t%v\n", time.Unix(b.Timestamp, 0))
	fmt.Printf("\tMerkle Value:\t%x\n", b.MerkleTree.Root.Data)
	fmt.Printf("\tPrev Block:\t%x\n", b.PrevBlockHash)
	fmt.Printf("\tHash Value:\t%x\n", b.Hash)
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) Block {
	newBlock := Block{
		Timestamp:     time.Now().Unix(),
		MerkleTree:    NewMerkleTree(transactions),
		PrevBlockHash: prevBlockHash,
	}
	newBlock.SetHash()
	return newBlock
}

func (b *Block) HashABLock() []byte {
	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, uint64(b.Timestamp))
	headers := append(timestampBytes, b.MerkleTree.Root.Data...)
	headers = append(headers, b.PrevBlockHash...)
	hash := sha256.Sum256(headers)
	return hash[:]
}

func (b *Block) SetHash() {
	b.Hash = b.HashABLock()
}
