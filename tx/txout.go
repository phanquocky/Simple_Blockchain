package tx

import (
	"blockchain_go/util"
	"bytes"
	"encoding/json"
	"log"
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

type OutpointWithIndex struct {
	Value      int
	PubKeyHash []byte
	OutputIdx  int
}

// TXOutputs collects TXOutput
type TXOutputs struct {
	Outputs []OutpointWithIndex
}

func (outs TXOutputs) AddOutput(out OutpointWithIndex) {
	outs.Outputs = append(outs.Outputs, out)
}

func (outs TXOutputs) Serialize() []byte {
	outsBytes, err := json.Marshal(outs)
	if err != nil {
		log.Println("cannot json.Marchal outs.serialize()!, ", err)
		return nil
	}
	return outsBytes
}

// DeserializeOutputs deserializes TXOutputs
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	json.Unmarshal(data, &outputs)
	return outputs
}
