package cli

import (
	"blockchain_go/lab2/address"
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// getP2pkhAddress function create p2pkh address and return private key for this address
func (cli *CLI) getP2pkhAddress() error {
	address, privkey, err := address.GenerateP2pkhAddress()
	if err != nil {
		return err
	}
	log.Println("Your Address: ", address)
	log.Println("Your Private Key: ", privkey)

	return nil
}

// spendP2pkh function using privkey to sign transaction
// using prevhash:prevOutIdx output as input of this transaction
func (cli *CLI) spendP2pkh(privkey, prevhash string, prevOutIdx uint32) error {

	privkeybtc, err := stringToPrivKey(privkey)
	if err != nil {
		return err
	}
	pubKeyHash := btcutil.Hash160(privkeybtc.PubKey().SerializeCompressed())

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

	signature, err := txscript.SignatureScript(
		tx, 0, commitTx.MsgTx().TxOut[prevOutIdx].PkScript,
		txscript.SigHashAll, privkeybtc, true,
	)
	if err != nil {
		return err
	}

	tx.TxIn[0].SignatureScript = signature

	cli.sendtx(tx)
	return nil
}
