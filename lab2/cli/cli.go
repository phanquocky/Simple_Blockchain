package cli

import (
	"blockchain_go/lab2/address"
	"flag"
	"fmt"

	"log"
	"os"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

type CLI struct {
	client *rpcclient.Client
}

func New(client *rpcclient.Client) CLI {
	return CLI{
		client: client,
	}
}

func (cli *CLI) Run() {

	var (
		getNewAddresscmd      = flag.NewFlagSet("getnewaddress", flag.ExitOnError)
		spendFundcmd          = flag.NewFlagSet("spendfund", flag.ExitOnError)
		getMultiSigAddresscmd = flag.NewFlagSet("getnewmultisig", flag.ExitOnError)
		spendMultiSigcmd      = flag.NewFlagSet("spendmultisig", flag.ExitOnError)
	)

	var (
		privKey    = spendFundcmd.String("privkey", "", "The private key")
		prevhash   = spendFundcmd.String("prevhash", "", "The hash transaction you want to spend")
		prevOutIdx = spendFundcmd.Uint("outidx", 0, "The output index of the output you want to spend")

		privKey1      = spendMultiSigcmd.String("privkey1", "", "The private1 key")
		privKey2      = spendMultiSigcmd.String("privkey2", "", "The private2 key")
		mulPrevhash   = spendMultiSigcmd.String("prevhash", "", "The hash transaction you want to spend")
		mulprevOutIdx = spendMultiSigcmd.Uint("outidx", 0, "The output index of the output you want to spend")
		redeemScript  = spendMultiSigcmd.String("redeem", "", "redeem script")
	)

	switch os.Args[1] {
	case "getnewaddress":
		err := getNewAddresscmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.getNewAddress()
		if err != nil {
			log.Fatal(err)
		}

	case "spendfund":
		err := spendFundcmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.spendFund(*privKey, *prevhash, uint32(*prevOutIdx))
		if err != nil {
			log.Fatal(err)
		}

	case "getnewmultisig":
		err := getMultiSigAddresscmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.getNewMultiSig()
		if err != nil {
			log.Fatal(err)
		}

	case "spendmultisig":
		err := spendMultiSigcmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.spendMultiSig([2]string{*privKey1, *privKey2}, *mulPrevhash, uint32(*mulprevOutIdx), *redeemScript)
		if err != nil {
			log.Fatal(err)
		}
	default:
		// cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) getNewAddress() error {
	address, privkey, err := address.GenerateAddress()
	if err != nil {
		return err
	}
	log.Println("Your Address: ", address)
	log.Println("Your Private Key: ", privkey)

	return nil
}

func (cli *CLI) spendFund(privkey, prevhash string, prevOutIdx uint32) error {

	privkeybtc, pubkey, err := stringToPrivKey(privkey)
	if err != nil {
		return err
	}
	fmt.Printf("privkey string %x \n", btcutil.Hash160(pubkey.SerializeCompressed()))
	pubKeyHash := btcutil.Hash160(pubkey.SerializeCompressed())

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

func (cli *CLI) sendtx(tx *wire.MsgTx) error {
	hash, err := cli.client.SendRawTransaction(tx, true)
	if err != nil {
		log.Println("cannot send raw tx, err ", err)
		return err
	}
	log.Println("Send transaction success!, txhash: ", hash.String())
	return nil
}
