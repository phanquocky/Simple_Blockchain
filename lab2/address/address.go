package address

import (
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func GenerateAddress() (*btcutil.AddressPubKeyHash, string, error) {
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Println("cannot create newPrivateKey ")
		return nil, "", err
	}

	pubkey := privKey.PubKey().SerializeCompressed()

	address, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubkey), &chaincfg.TestNet3Params)
	if err != nil {
		log.Println("cannot create pubkeyhash address from pubkey: ", pubkey)
		return nil, "", err
	}

	return address, hex.EncodeToString(privKey.Serialize()), err
}

func Generate2to2Address() (*btcutil.AddressScriptHash, []*btcec.PrivateKey, []byte, error) {

	privKey1, err := btcec.NewPrivateKey()
	if err != nil {
		log.Println("cannot create newPrivateKey ")
		return nil, nil, nil, err
	}

	pubkey1 := privKey1.PubKey().SerializeCompressed()
	address1, err := btcutil.NewAddressPubKey(pubkey1,
		&chaincfg.TestNet3Params)

	privKey2, err := btcec.NewPrivateKey()
	if err != nil {
		log.Println("cannot create newPrivateKey ")
		return nil, nil, nil, err
	}

	pubkey2 := privKey2.PubKey().SerializeCompressed()
	address2, err := btcutil.NewAddressPubKey(pubkey2,
		&chaincfg.TestNet3Params)

	pkScript, err := txscript.MultiSigScript(
		[]*btcutil.AddressPubKey{address1, address2},
		2)

	address, err := btcutil.NewAddressScriptHash(pkScript, &chaincfg.TestNet3Params)
	if err != nil {
		log.Println("cannot create multisig address")
		return nil, nil, nil, err
	}

	// scriptPkScript, err := txscript.PayToAddrScript(address)
	// if err != nil {

	// }

	return address, []*btcec.PrivateKey{privKey1, privKey2}, pkScript, nil
}
