package block

import (
	tree "blockchain_go/merkleTree"
	"blockchain_go/tx"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

const (
	HASH_SIZE = sha256.Size
)

type Hash [HASH_SIZE]byte

type Block struct {
	Nonce         int
	Timestamp     int64
	Transactions  []*tx.Transaction
	PrevBlockHash Hash
	BlockHash     Hash
	Difficulty    uint32
}

func NewBlock(transactions []*tx.Transaction, prevBlockHash Hash, difficulty uint32) *Block {
	block := &Block{0, time.Now().Unix(), transactions, prevBlockHash, Hash{}, difficulty}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.BlockHash = hash
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(coinbase *tx.Transaction, difficulty uint32) *Block {
	return NewBlock([]*tx.Transaction{coinbase}, Hash{}, difficulty)
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte
	for _, transaction := range b.Transactions {
		transactions = append(transactions, transaction.Serialize())
	}
	mTree := tree.NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		return nil
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil
	}

	return &block
}
