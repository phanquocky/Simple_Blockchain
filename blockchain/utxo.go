package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"blockchain_go/util"
	"encoding/hex"

	"golang.org/x/exp/slices"
)

func (bc *Blockchain) FindUTXO(address string) map[*tx.Transaction][]int {
	var (
		utxos      = make(map[*tx.Transaction][]int)
		spentTXOs  = make(map[string][]int)
		bci        = bc.Iterator()
		pubkeyHash = util.GetPubkeyHash(address)
	)

	for {
		bl := bci.Next()

		for _, tx := range bl.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					if slices.Contains(spentTXOs[txID], outIdx) {
						continue Outputs
					}
				}

				if out.IsLockedWithKey(pubkeyHash) {
					utxos[tx] = append(utxos[tx], outIdx)
				}
			}

			if !tx.IsCoinbase() {
				for _, in := range tx.Vin {
					//TODO if in.CanUnlockOutputWith(address) {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.OutIdx)
					// }
				}
			}
		}

		if block.Hash(block.Hash{}) == bl.PrevBlockHash {
			break
		}
	}

	return utxos
}
