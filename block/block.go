package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

const (
	HASH_SIZE = sha256.Size
)

type Hash [HASH_SIZE]byte

type Block struct {
	Nonce         int
	Timestamp     int64
	Data          []byte
	PrevBlockHash Hash
	BlockHash     Hash
	Difficulty    uint32
}

func NewBlock(data string, prevBlockHash Hash, difficulty uint32) *Block {
	block := &Block{0, time.Now().Unix(), []byte(data), prevBlockHash, Hash{}, difficulty}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.BlockHash = hash
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(difficulty uint32) *Block {
	return NewBlock("Genesis Block", Hash{}, difficulty)
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash[:], b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.BlockHash = hash
}
