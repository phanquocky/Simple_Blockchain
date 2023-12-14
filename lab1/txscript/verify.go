package txscript

import (
	"blockchain_go/tx"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"math/big"
)

func VerifyTransaction(
	txsign *tx.Transaction,
	prevOutputFetchers []PrevOutputFetcher,
) bool {

	for idx := range txsign.Vin {
		if !VerifyOneInput(txsign, prevOutputFetchers[idx].PkHash, idx, prevOutputFetchers[idx].Amt) {
			return false
		}
	}

	return true
}

func VerifyOneInput(
	txsign *tx.Transaction,
	prevPubkeyHash []byte, idx int, amt int64,
) bool {

	curve := elliptic.P256()

	hash, err := calcSignatureHashRaw(prevPubkeyHash, txsign, idx, amt)
	if err != nil {
		log.Println("cannot calc Sign hash raw!, ", err)
		return false
	}

	in := txsign.Vin[idx]

	r := big.Int{}
	s := big.Int{}
	sigLen := len(in.Signature)
	r.SetBytes(in.Signature[:(sigLen / 2)])
	s.SetBytes(in.Signature[(sigLen / 2):])

	x := big.Int{}
	y := big.Int{}
	keyLen := len(in.PubKey)
	x.SetBytes(in.PubKey[:(keyLen / 2)])
	y.SetBytes(in.PubKey[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}

	return ecdsa.Verify(&rawPubKey, hash, &r, &s)
}
