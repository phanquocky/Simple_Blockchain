package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	VERSION                 = byte(0x00)
	WALLET_FILE             = "wallet.dat"
	ADDRESS_CHECKSUM_LENGTH = 4
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	if public == nil || private == ecdsa.PrivateKey(ecdsa.PrivateKey{}) {
		return nil
	}

	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Println("Cannot generate random key!, ", err)
		return ecdsa.PrivateKey{}, nil
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

func (w Wallet) GetAddress() string {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{VERSION}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)

	return base58.Encode(fullPayload)
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Println("cannot ripemd160 pubkey")
		return nil
	}

	return RIPEMD160Hasher.Sum(nil)
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:ADDRESS_CHECKSUM_LENGTH]
}
