package main

import "fmt"

type Command struct {
	Name     string
	HelpText string
	Params   []string
}

var commands = []Command{
	{
		Name:     "createchain",
		HelpText: "Create a new chain with Genesis block",
		Params:   []string{},
	},
	{
		Name:     "printchain",
		HelpText: "Display all blocks in the blockchain",
		Params:   []string{},
	},
	{
		Name:     "addblock",
		HelpText: "Create a block and add into the blockchain",
		Params:   []string{"data"},
	},
	{
		Name:     "validtran",
		HelpText: "Validate a transaction is in a block",
		Params:   []string{"data", "block"},
	},
}

func (cmd *Command) printCommandUsage() {
	fmt.Printf("  %s\t%s\n", cmd.Name, cmd.HelpText)
	if len(cmd.Params) != 0 {
		fmt.Println("\tArguments:")
		for _, param := range cmd.Params {
			fmt.Printf("  \t-%s=[VALUE]\n", param)
		}
	}
}

func printUsage(program string) {
	fmt.Printf("Usage: %s <command> [options]\n", program)
	fmt.Println("Commands:")
	for _, cmd := range commands {
		cmd.printCommandUsage()
	}
}

func createChain() {
	bc := BuildBlockchain()
	bc.SaveBlockchain()
}

func printChain() {
	bc, err := LoadBlockchain()
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		fmt.Println("Create new blockchain with createchain command")
		return
	}
	bc.Print()
}

func addBlock(datas []string) {
	bc, err := LoadBlockchain()
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		fmt.Println("Create new blockchain with createchain command")
		return
	}
	transactions := make([]*Transaction, 0)
	for _, data := range datas {
		transactions = append(transactions, NewTransaction([]byte(data)))
	}
	bc.AddBlock(transactions)
	bc.SaveBlockchain()
}

func validTransaction(data []string, blockHash string) {
	bc, err := LoadBlockchain()
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		fmt.Println("Create new blockchain with createchain command")
		return
	}
	var foundBlock = false
	for _, block := range bc.Blocks {
		if fmt.Sprintf("%x", block.Hash) == blockHash {
			foundBlock = true
			tree := NewMerkleTree(block.Transactions)

			for _, d := range data {

				valid := MerkleVerify(tree, []byte(d))
				if valid {
					fmt.Printf("Transaction data %v is stored in the block\n", d)
				} else {
					fmt.Printf("Transaction data %v is NOT stored in the block\n", d)
				}
			}
			break
		}
	}
	if foundBlock == false {
		fmt.Printf("Block %v was not found in blockchain", blockHash)
	}
}
