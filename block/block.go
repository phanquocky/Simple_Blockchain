package block

import (
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
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
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
