package cli

import (
	"blockchain_go/lab2/address"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func (cli *CLI) getNewMultiSig() error {
	address, privKeys, redeem, err := address.Generate2to2Address()
	if err != nil {
		return err
	}
	log.Println("Your multisig address: ", address.String())
	log.Println("First Private Key: ", hex.EncodeToString(privKeys[0].Serialize()))

	log.Println("Second Private Key: ", hex.EncodeToString(privKeys[1].Serialize()))
	log.Println("Redeem script: ", hex.EncodeToString(redeem))

	return nil
}

func (cli *CLI) spendMultiSig(privkeys [2]string, prevhash string, prevOutIdx uint32, redeem string) error {

	var privKeysbtc [2]*btcec.PrivateKey
	for i, v := range privkeys {
		privkey, _, err := stringToPrivKey(v)
		if err != nil {
			return err
		}
		privKeysbtc[i] = privkey
	}
	pk1 := privKeysbtc[0].PubKey().SerializeCompressed()
	address1, err := btcutil.NewAddressPubKey(pk1,
		&chaincfg.TestNet3Params)
	if err != nil {
		return err
	}

	pk2 := privKeysbtc[1].PubKey().SerializeCompressed()
	address2, err := btcutil.NewAddressPubKey(pk2,
		&chaincfg.TestNet3Params)
	if err != nil {
		return err
	}

	outpoint, err := createOutpoint(prevhash, prevOutIdx)
	if err != nil {
		return err
	}

	commitTx, err := cli.client.GetRawTransaction(&outpoint.Hash)
	if err != nil {
		return err
	}

	tx := wire.NewMsgTx(1)
	tx.AddTxIn(&wire.TxIn{
		PreviousOutPoint: *outpoint,
	})

	pubKeyHash := btcutil.Hash160(privKeysbtc[0].PubKey().SerializeCompressed())
	pkScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_DUP).AddOp(txscript.OP_HASH160).
		AddData(pubKeyHash).AddOp(txscript.OP_EQUALVERIFY).AddOp(txscript.OP_CHECKSIG).
		Script()
	if err != nil {
		return err
	}

	tx.AddTxOut(&wire.TxOut{
		Value:    1000,
		PkScript: pkScript,
	})

	redeemBytes, err := hex.DecodeString(redeem)
	if err != nil {
		return err
	}
	scriptAddr, err := btcutil.NewAddressScriptHash(redeemBytes, &chaincfg.TestNet3Params)
	fmt.Println("scriptAddr: ", scriptAddr.String())

	signature, err := txscript.SignTxOutput(&chaincfg.TestNet3Params,
		tx, 0, commitTx.MsgTx().TxOut[prevOutIdx].PkScript, txscript.SigHashAll,
		mkGetKey(map[string]addressToKey{
			address1.EncodeAddress(): {privKeysbtc[0], true},
		}), mkGetScript(map[string][]byte{
			scriptAddr.EncodeAddress(): redeemBytes,
		}), nil)
	if err != nil {
		return err
	}

	signature, err = txscript.SignTxOutput(
		&chaincfg.TestNet3Params, tx, 0,
		commitTx.MsgTx().TxOut[prevOutIdx].PkScript, txscript.SigHashAll,
		mkGetKey(map[string]addressToKey{
			address1.EncodeAddress(): {privKeysbtc[0], true},
			address2.EncodeAddress(): {privKeysbtc[1], true},
		}), mkGetScript(map[string][]byte{
			scriptAddr.EncodeAddress(): redeemBytes,
		}), signature,
	)
	if err != nil {
		return err
	}

	tx.TxIn[0].SignatureScript = signature

	cli.sendtx(tx)
	return nil

}

type addressToKey struct {
	key        *btcec.PrivateKey
	compressed bool
}

func mkGetKey(keys map[string]addressToKey) txscript.KeyDB {
	if keys == nil {
		return txscript.KeyClosure(func(addr btcutil.Address) (*btcec.PrivateKey,
			bool, error) {
			return nil, false, errors.New("nope")
		})
	}
	return txscript.KeyClosure(func(addr btcutil.Address) (*btcec.PrivateKey,
		bool, error) {
		a2k, ok := keys[addr.EncodeAddress()]
		if !ok {
			return nil, false, errors.New("nope")
		}
		return a2k.key, a2k.compressed, nil
	})
}

func mkGetScript(scripts map[string][]byte) txscript.ScriptDB {
	if scripts == nil {
		return txscript.ScriptClosure(func(addr btcutil.Address) ([]byte, error) {
			return nil, errors.New("nope")
		})
	}
	return txscript.ScriptClosure(func(addr btcutil.Address) ([]byte, error) {
		script, ok := scripts[addr.EncodeAddress()]
		if !ok {
			return nil, errors.New("nope")
		}
		return script, nil
	})
}
