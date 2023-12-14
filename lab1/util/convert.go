package util

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/btcsuite/btcd/btcutil/base58"
)

// IntToHex converts an int64 to a byte array
func Int64ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Int32ToHex converts an int32 to a byte array
func Uint32ToHex(num uint32) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func GetPubkeyHash(address string) []byte {
	pubKeyHash := base58.Decode(address)
	return pubKeyHash[1 : len(pubKeyHash)-4]
}
