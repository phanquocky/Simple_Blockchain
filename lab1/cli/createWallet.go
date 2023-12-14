package cli

import (
	"blockchain_go/wallet"
	"log"
)

func (cli *CLI) createWallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	log.Printf("Your new address: %s\n", address)
}
