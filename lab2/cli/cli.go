package cli

import (
	"flag"

	"log"
	"os"

	"github.com/btcsuite/btcd/rpcclient"
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
		getP2pkhAddressCmd    = flag.NewFlagSet("getp2pkhaddress", flag.ExitOnError)
		spendP2pkhCmd         = flag.NewFlagSet("spendp2pkh", flag.ExitOnError)
		getMultiSigAddressCmd = flag.NewFlagSet("getmultisigaddress", flag.ExitOnError)
		spendMultiSigCmd      = flag.NewFlagSet("spendmultisig", flag.ExitOnError)
	)

	var (
		privKey    = spendP2pkhCmd.String("privkey", "", "The private key")
		prevhash   = spendP2pkhCmd.String("prevhash", "", "The hash transaction you want to spend")
		prevOutIdx = spendP2pkhCmd.Uint("outidx", 0, "The output index of the output you want to spend")

		privKey1      = spendMultiSigCmd.String("privkey1", "", "The private1 key")
		privKey2      = spendMultiSigCmd.String("privkey2", "", "The private2 key")
		mulPrevhash   = spendMultiSigCmd.String("prevhash", "", "The hash transaction you want to spend")
		mulprevOutIdx = spendMultiSigCmd.Uint("outidx", 0, "The output index of the output you want to spend")
		redeemScript  = spendMultiSigCmd.String("redeem", "", "redeem script")
	)

	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "getp2pkhaddress":
		err := getP2pkhAddressCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.getP2pkhAddress()
		if err != nil {
			log.Fatal(err)
		}

	case "spendp2pkh":
		err := spendP2pkhCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.spendP2pkh(*privKey, *prevhash, uint32(*prevOutIdx))
		if err != nil {
			log.Fatal(err)
		}

	case "getmultisigaddress":
		err := getMultiSigAddressCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.getNewMultiSig()
		if err != nil {
			log.Fatal(err)
		}

	case "spendmultisig":
		err := spendMultiSigCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		err = cli.spendMultiSig([2]string{*privKey1, *privKey2}, *mulPrevhash, uint32(*mulprevOutIdx), *redeemScript)
		if err != nil {
			log.Fatal(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
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

func (cli *CLI) printUsage() {
	log.Println("Usage:")
	log.Println("  getp2pkhaddress 			- params() get new p2pkh address and return private key of this address")
	log.Println("  spendp2pkh 					- params(privkey, prevhash, outidx) spend p2pkh outpoint (prevhash:outidx) with private key")
	log.Println("  getmultisigaddress 	- params() get new 2-2 multisig and return 2 private key, redeem script")
	log.Println("  spendmultisig 				- params(privkey1, prevkey2, prevhash, outidx, redeem) using 2 private key, and redeem script to spend outpoint (prevhash:outidx)")
}
