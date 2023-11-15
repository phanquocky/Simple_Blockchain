package tx

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const (
	SUBSIDY = 10
)

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func NewCoinbaseTX(address, data string) *Transaction {

	txin := TXInput{[]byte{}, -1, nil, nil}
	txout := TXOutput{SUBSIDY, nil}
	txout.Lock(address)

	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	log.Printf("Reward to '%s' \n", address)
	return &tx
}

func NewTransaction() *Transaction {
	return &Transaction{
		ID:   []byte{},
		Vin:  make([]TXInput, 0),
		Vout: make([]TXOutput, 0),
	}
}

// SetID sets ID of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].OutIdx == -1
}

func (tx *Transaction) AddTxInput(txinput *TXInput) {
	tx.Vin = append(tx.Vin, *txinput)
}

func (tx *Transaction) AddTxOutput(txoutput *TXOutput) {
	tx.Vout = append(tx.Vout, *txoutput)
}
