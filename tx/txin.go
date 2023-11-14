package tx

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
}

func NewTxInput(txid []byte, vout int, signature []byte) *TXInput {
	return &TXInput{
		Txid:      txid,
		Vout:      vout,
		Signature: signature,
	}
}
