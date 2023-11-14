package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) Run() {

	cli.validateArgs()

	var (
		printChainCmd       = flag.NewFlagSet("printchain", flag.ExitOnError)
		createBlockchainCmd = flag.NewFlagSet("createblockchain", flag.ExitOnError)
		getBalanceCmd       = flag.NewFlagSet("getbalance", flag.ExitOnError)
		transferCmd         = flag.NewFlagSet("transfer", flag.ExitOnError)
	)

	var (
		createBlockchainAddr = createBlockchainCmd.String("address", "", "The address coinbase transaction in genesis block")
		getBalanceAddr       = getBalanceCmd.String("address", "", "The address balance")
		transferFrom         = transferCmd.String("from", "", "Source address")
		transferTo           = transferCmd.String("to", "", "Destination address")
		transferAmount       = transferCmd.Uint("amount", 0, "Amount to send")
	)

	switch os.Args[1] {
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		cli.printChain()

	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *createBlockchainAddr == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddr)

	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *getBalanceAddr == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}

		cli.getBalance(*getBalanceAddr)

	case "transfer":
		err := transferCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *transferFrom == "" || *transferTo == "" || *transferAmount == 0 {
			transferCmd.Usage()
			os.Exit(1)
		}

		cli.transfer(*transferFrom, *transferTo, *transferAmount)
	default:
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(0)
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  printchain 			- params() print all the blocks of the blockchain")
	fmt.Println("  createblockchain - params(-address) create a new block  chain")
	fmt.Println("  getbalance 			- params(-address) get balance of address")
	fmt.Println("  transfer 				- params(-from, -to, -amount) transfer coin")
}
