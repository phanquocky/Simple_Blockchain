package main

import (
	"flag"
	"fmt"
	"os"
)

// main function
func main() {
	program := os.Args[0]

	if len(os.Args) < 2 {
		printUsage(program)
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "help":
		printUsage(program)

	case "createchain":
		createChainCommand := flag.NewFlagSet("createchain", flag.ExitOnError)
		err := createChainCommand.Parse(os.Args[2:])
		if err != nil {
			createChainCommand.Usage()
		} else {
			createChain()
		}

	case "printchain":
		printChainCommand := flag.NewFlagSet("printchain", flag.ExitOnError)
		err := printChainCommand.Parse(os.Args[2:])
		if err != nil {
			printChainCommand.Usage()
		} else {
			printChain()
		}

	case "addblock":
		addBlockCommand := flag.NewFlagSet("addblock", flag.ExitOnError)
		var data multiString
		addBlockCommand.Var(&data, "data", "Data for the block")
		err := addBlockCommand.Parse(os.Args[2:])
		if err != nil || len(data) == 0 {
			addBlockCommand.Usage()
		} else {
			addBlock(data)
		}

	case "validtran":
		validTranCommand := flag.NewFlagSet("validtran", flag.ExitOnError)
		data := validTranCommand.String("data", "", "Data of the transaction")
		blockHash := validTranCommand.String("block", "", "Hash value of the block")
		err := validTranCommand.Parse(os.Args[2:])
		if err != nil {
			validTranCommand.Usage()
		} else {
			validTransaction(*data, *blockHash)
		}

	default:
		fmt.Println("Unknown command.")
		printUsage(program)
	}
}
