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
	Timestamp     int64
	Data          []byte
	PrevBlockHash Hash
	BlockHash     Hash
}

func NewBlock(data string, prevBlockHash Hash) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, Hash{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", Hash{})
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash[:], b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.BlockHash = hash
}
