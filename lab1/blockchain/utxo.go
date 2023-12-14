package blockchain

import (
	"blockchain_go/block"
	"blockchain_go/tx"
	"blockchain_go/util"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"golang.org/x/exp/slices"
)

const (
	UTXO_BUCKET = "utxobucket"
)

type UTXOSet struct {
	Blockchain *Blockchain
}

func NewUTXOSet(bc *Blockchain) UTXOSet {
	return UTXOSet{Blockchain: bc}
}

func (u UTXOSet) Reindex() {
	db := u.Blockchain.DB
	bucketName := []byte(UTXO_BUCKET)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil {
			log.Printf("cannot delete bucket: %x, %s \n ", bucketName, err.Error())
		}

		_, err = tx.CreateBucket(bucketName)
		if err != nil {
			log.Printf("cannot create bucket: %x, %s \n ", bucketName, err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		log.Printf("cannot update bucket, %x \n", bucketName)
		return
	}

	UTXO := u.Blockchain.FindUTXO()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		for txID, outs := range UTXO {
			key, err := hex.DecodeString(txID)
			fmt.Printf("key: %x \n", key)
			if err != nil {
				log.Println("cannot decode txID, ", txID, err)
				return err
			}

			err = b.Put(key, outs.Serialize())
			if err != nil {
				log.Println("cannot put to database, ", key)
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println("cannot update utxo bucket, ", err)
		return
	}

}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (bc *Blockchain) FindUTXO() map[string]tx.TXOutputs {
	UTXO := make(map[string]tx.TXOutputs)
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		bl := bci.Next()

		for _, transaction := range bl.Transactions {
			txID := hex.EncodeToString(transaction.ID)

		Outputs:
			for outIdx, out := range transaction.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					if slices.Contains(spentTXOs[txID], outIdx) {
						continue Outputs
					}
				}

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs,
					tx.OutpointWithIndex{
						Value:      out.Value,
						PubKeyHash: out.PubKeyHash,
						OutputIdx:  outIdx,
					})
				UTXO[txID] = outs
			}

			if !transaction.IsCoinbase() {
				for _, in := range transaction.Vin {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.OutIdx)
				}
			}
		}

		if block.Hash(block.Hash{}) == bl.PrevBlockHash {
			break
		}
	}

	return UTXO
}

func (u UTXOSet) FindUTXOByAddress(address string) map[string]tx.TXOutputs {
	unspentOutputs := make(map[string]tx.TXOutputs)
	pubkeyHash := util.GetPubkeyHash(address)
	db := u.Blockchain.DB
	fmt.Println("this is correct")
	err := db.View(func(dbtx *bolt.Tx) error {
		b := dbtx.Bucket([]byte(UTXO_BUCKET))
		fmt.Println("bucket: ", b)
		c := b.Cursor()
		fmt.Println("cursor: ", c)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Println(1)
			fmt.Printf("TXID11: %x \n", k)
			txID := hex.EncodeToString(k)
			outs := tx.DeserializeOutputs(v)
			fmt.Println("TXID:", txID)
			for _, out := range outs.Outputs {
				fmt.Println(out.PubKeyHash, out.Value, pubkeyHash)

				if bytes.Equal(out.PubKeyHash, pubkeyHash) {
					fmt.Println("OOKK")
					tem := unspentOutputs[txID]
					tem.Outputs = append(tem.Outputs, out)
					unspentOutputs[txID] = tem
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Println("cannot view bucket: ", UTXO_BUCKET, err)
		return nil
	}

	return unspentOutputs
}

func (u UTXOSet) Update(block *block.Block) {
	db := u.Blockchain.DB

	err := db.Update(func(dbtx *bolt.Tx) error {
		b := dbtx.Bucket([]byte(UTXO_BUCKET))

		for _, btx := range block.Transactions {
			if !btx.IsCoinbase() {
				for _, vin := range btx.Vin {
					updatedOuts := tx.TXOutputs{}
					outsBytes := b.Get(vin.Txid)
					outs := tx.DeserializeOutputs(outsBytes)

					for outIdx, out := range outs.Outputs {
						if outIdx != vin.OutIdx {
							updatedOuts.Outputs = append(updatedOuts.Outputs, out)
						}
					}

					if len(updatedOuts.Outputs) == 0 {
						err := b.Delete(vin.Txid)
						if err != nil {
							log.Println("cannot delete: ", vin.Txid, err)
							return err
						}
					} else {
						err := b.Put(vin.Txid, updatedOuts.Serialize())
						if err != nil {
							log.Println("cannot update: ", vin.Txid, err)
							return err
						}
					}

				}
			}

			newOutputs := tx.TXOutputs{}
			for outIdx, out := range btx.Vout {
				newOutputs.Outputs = append(newOutputs.Outputs,
					tx.OutpointWithIndex{
						PubKeyHash: out.PubKeyHash,
						Value:      out.Value,
						OutputIdx:  outIdx,
					})
			}

			err := b.Put(btx.ID, newOutputs.Serialize())
			if err != nil {
				log.Println("cannot update utxo set, ", err)
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Println("cannot update utxo database, ", err)
		return
	}
}
