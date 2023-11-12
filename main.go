package main

import (
	"blockchain_go/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()

	for i := 0; i < 3000; i++ {
		bc.AddBlock("Send 1 BTC to Ivan")
		bl := bc.Blocks[len(bc.Blocks)-1]
		fmt.Printf("Prev. hash: %x\n", bl.PrevBlockHash)
		fmt.Printf("Data: %s\n", bl.Data)
		fmt.Printf("Hash: %x\n", bl.BlockHash)
		fmt.Printf("Difficulty: %d\n", bl.Difficulty)
		fmt.Printf("Nonce: %d\n", bl.Nonce)
		fmt.Printf("Timestamp: %d\n", bl.Timestamp)
		fmt.Println()
	}
	// bc.AddBlock("Send 2 more BTC to Ivan")
	// bc.AddBlock("Send 3 BTC to Ivan")
	// bc.AddBlock("Send 4 more BTC to Ivan")

	// for _, bl := range bc.Blocks {
	// 	fmt.Printf("Prev. hash: %x\n", bl.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", bl.Data)
	// 	fmt.Printf("Hash: %x\n", bl.BlockHash)
	// 	fmt.Printf("Difficulty: %d\n", bl.Difficulty)
	// 	fmt.Printf("Nonce: %d\n", bl.Nonce)
	// 	fmt.Printf("Timestamp: %d\n", bl.Timestamp)

	// 	fmt.Println()
	// 	pow := block.NewProofOfWork(bl)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

}
