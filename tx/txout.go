package tx

import (
	"blockchain_go/util"
	"bytes"
)

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

func (out *TXOutput) Lock(address string) {
	out.PubKeyHash = util.GetPubkeyHash(address)
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

func NewTxOutput(value int, pubKeyHash []byte) *TXOutput {
	return &TXOutput{
		Value:      value,
		PubKeyHash: pubKeyHash,
	}
}
