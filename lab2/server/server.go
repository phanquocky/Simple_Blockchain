package server

import (
	"os"
	"path/filepath"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
)

const (
	RPC_USER     = "admin"
	RPC_PASS     = "admin123"
	RPC_HOST     = "localhost:18332"
	RPC_ENDPOINT = "ws"
)

func NewClient() (*rpcclient.Client, error) {
	// Connect to local btcwallet RPC server using websockets.
	certHomeDir := btcutil.AppDataDir("btcwallet", false)
	certs, err := os.ReadFile(filepath.Join(certHomeDir, "rpc.cert"))
	if err != nil {
		return nil, err
	}
	connCfg := &rpcclient.ConnConfig{
		Host:         RPC_HOST,
		Endpoint:     RPC_ENDPOINT,
		User:         RPC_USER,
		Pass:         RPC_PASS,
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
