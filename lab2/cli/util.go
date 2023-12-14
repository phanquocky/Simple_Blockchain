package cli

import (
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// createOutpoint function create wire.Outpoint from prevhash(string),
// outIdx uint32
func createOutpoint(
	prevHash string, outIdx uint32,
) (*wire.OutPoint, error) {
	prevHashByte, err := hex.DecodeString(prevHash)
	if err != nil {
		log.Println("cannot decode string: err, ", prevHash, err)
		return nil, err
	}
	// reverse prevHashByte
	for i, j := 0, len(prevHashByte)-1; i < j; i, j = i+1, j-1 {
		prevHashByte[i], prevHashByte[j] = prevHashByte[j], prevHashByte[i]
	}

	prevCommitTxHash, err := chainhash.NewHash(prevHashByte)
	if err != nil {
		log.Println("cannot chainhash.NewHash(prevHashByte): err, ", err)
		return nil, err
	}
	outpoint := wire.NewOutPoint(prevCommitTxHash, outIdx)

	return outpoint, nil
}

func stringToPrivKey(privkey string) (*btcec.PrivateKey, error) {
	privkeyBytes, err := hex.DecodeString(privkey)
	if err != nil {
		log.Println("cannot decode string: err, ", privkey, err)
		return nil, err
	}
	privateKey, _ := btcec.PrivKeyFromBytes(privkeyBytes)

	return privateKey, nil
}
