package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const TARGET_BITS = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	var (
		exponent    uint8
		coefficient big.Int
		target      *big.Int
	)
	diffcultyBytes := Uint32ToHex(b.Difficulty)
	exponent = uint8(diffcultyBytes[0])
	coefficient.SetBytes(diffcultyBytes[1:])

	// target = coefficient * (2 ^ (8 * (exponent - 3)))
	target = big.NewInt(1)
	target.Lsh(target, uint(8*(exponent-3)))
	target.Mul(target, &coefficient)

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash[:],
			pow.block.Data,
			Int64ToHex(pow.block.Timestamp),
			Int64ToHex(int64(TARGET_BITS)),
			Int64ToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, Hash) {

	var (
		hashInt big.Int
		hash    Hash
	)
	nonce := 0
	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 { // hashInt < pow.target
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

// IntToHex converts an int64 to a byte array
func Int64ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Int32ToHex converts an int32 to a byte array
func Uint32ToHex(num uint32) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
