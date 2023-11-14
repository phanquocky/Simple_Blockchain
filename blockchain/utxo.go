package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"encoding/hex"
)

func (bc *Blockchain) FindUTXO(address string) map[*tx.Transaction][]int {
	var (
		utxos     = make(map[*tx.Transaction][]int)
		spentTXOs = make(map[string][]int)
		bci       = bc.Iterator()
	)

	for {
		bl := bci.Next()

		for _, tx := range bl.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
					// if slices.Contains(spentTXOs[txID], outIdx) {
					// 	continue Outputs
					// }
				}

				if out.CanBeUnlockedWith(address) {
					utxos[tx] = append(utxos[tx], outIdx)
				}
			}

			if !tx.IsCoinbase() {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if block.Hash(block.Hash{}) == bl.PrevBlockHash {
			break
		}
	}

	return utxos
}
