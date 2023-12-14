package txscript

import (
	"blockchain_go/tx"
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
)

func SignOneInput(privKey ecdsa.PrivateKey, txsign *tx.Transaction, prevPubkeyHash []byte, idx int, amt int64) {
	if txsign.IsCoinbase() {
		return
	}

	hash, err := calcSignatureHashRaw(prevPubkeyHash, txsign, idx, amt)
	if err != nil {
		log.Println("cannot calc Sign hash raw!, ", err)
		return
	}

	r, s, err := ecdsa.Sign(rand.Reader, &privKey, hash)
	if err != nil {
		log.Println("cannot sign using ecdsa, ", err)
		return
	}

	signature := append(r.Bytes(), s.Bytes()...)
	txsign.Vin[idx].Signature = signature
}

func calcSignatureHashRaw(prevPubkeyHash []byte, txsign *tx.Transaction, idx int, amt int64) ([]byte, error) {

	if idx > len(txsign.Vin)-1 {
		return nil, fmt.Errorf("idx %d but %d txins", idx, len(txsign.Vin))
	}

	var sigHash bytes.Buffer

	txIn := txsign.Vin[idx]

	// write previous outpoint hash
	sigHash.Write(txIn.Txid)

	// write previous outpoint index
	var bIndex [4]byte
	binary.LittleEndian.PutUint32(bIndex[:], uint32(txIn.OutIdx))
	sigHash.Write(bIndex[:])

	// write previous pubkey hash
	sigHash.Write(prevPubkeyHash)

	// write all tx output
	var bOuputs bytes.Buffer
	for _, out := range txsign.Vout {
		bOuputs.Write(out.PubKeyHash)

		var bValue [4]byte
		binary.LittleEndian.PutUint32(bValue[:], uint32(out.Value))
		bOuputs.Write(bIndex[:])
	}

	hashOutputs := sha256.Sum256(bOuputs.Bytes())
	sigHash.Write(hashOutputs[:])

	return sigHash.Bytes(), nil
}

func SignRawTransaction(txsign *tx.Transaction,
	privateKey ecdsa.PrivateKey,
	prevOutputFetchers []PrevOutputFetcher,
) {
	for idx := range txsign.Vin {
		SignOneInput(privateKey, txsign, prevOutputFetchers[idx].PkHash, idx, prevOutputFetchers[idx].Amt)
	}
}
