package tx

type TXInput struct {
	Txid      []byte
	OutIdx    int
	Signature []byte
	PubKey    []byte
}

func NewTxInput(txid []byte, vout int, signature []byte, pubkey []byte) *TXInput {
	return &TXInput{
		Txid:      txid,
		OutIdx:    vout,
		Signature: signature,
		PubKey:    pubkey,
	}
}
