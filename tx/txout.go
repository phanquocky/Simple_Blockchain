package tx

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewTxOutput(value int, scriptPubkey string) *TXOutput {
	return &TXOutput{
		Value:        value,
		ScriptPubKey: scriptPubkey,
	}
}
