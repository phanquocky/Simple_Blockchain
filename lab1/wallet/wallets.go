package wallet

import (
	"crypto/elliptic"
	"encoding/json"
	"log"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet `json:"wallets"`
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewWallets() (*Wallets, error) {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.ReadFromFile()

	return wallets, err
}

// CreateWallet adds a Wallet to Wallets
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.GetAddress()

	ws.Wallets[address] = wallet

	return address
}

// GetWallet returns a Wallet by its address
func (ws *Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) SaveToFile() {
	walletsBytes, err := json.Marshal(ws)

	if err != nil {
		log.Println("Cannot encode wallets to bytes")
		return
	}

	os.WriteFile(WALLET_FILE, walletsBytes, 0666)
}

func (ws *Wallets) ReadFromFile() error {
	curve := elliptic.P256()
	walletsBytes, err := os.ReadFile(WALLET_FILE)
	if err != nil {
		log.Println("Cannot read  wallets file")
		return err
	}

	json.Unmarshal(walletsBytes, ws)

	// TODO: ??? because when save to file it lose curve
	for _, v := range ws.Wallets {
		v.PrivateKey.Curve = curve
	}
	return nil
}
