package tx

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func NewTxInput(txid []byte, vout int, scriptSig string) *TXInput {
	return &TXInput{
		Txid:      txid,
		Vout:      vout,
		ScriptSig: scriptSig,
	}
}
