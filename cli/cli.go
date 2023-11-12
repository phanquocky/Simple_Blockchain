package cli

import (
	"blockchain_go/block"
	"blockchain_go/blockchain"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	BC *blockchain.Blockchain
}

func (cli *CLI) Run() {

	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(0)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.BC.Iterator()

	for {
		bl := bci.Next()

		fmt.Printf("Prev. hash: %x\n", bl.PrevBlockHash)
		fmt.Printf("Data: %s\n", bl.Data)
		fmt.Printf("Hash: %x\n", bl.BlockHash)
		pow := block.NewProofOfWork(bl)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if block.Hash(block.Hash{}) == bl.PrevBlockHash {
			break
		}
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}
