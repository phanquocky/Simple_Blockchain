package block

import (
	"blockchain_go/util"
	"bytes"
	"crypto/sha256"
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
	diffcultyBytes := util.Uint32ToHex(b.Difficulty)
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
			pow.block.HashTransactions(),
			util.Int64ToHex(pow.block.Timestamp),
			util.Int64ToHex(int64(TARGET_BITS)),
			util.Int64ToHex(int64(nonce)),
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

	log.Printf("Mining the block containing \"%x\"\n", pow.block.HashTransactions())
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
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
