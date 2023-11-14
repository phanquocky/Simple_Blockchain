package cli

import (
	"blockchain_go/blockchain"
	"fmt"
)

func (cli *CLI) createBlockchain(address string) {
	fmt.Println("Create blockchain")
	bc := blockchain.CreateBlockchain(address)
	fmt.Println("BLockchain, ", bc)
	bc.DB.Close()

	fmt.Println("Blockchain create success!")
}
