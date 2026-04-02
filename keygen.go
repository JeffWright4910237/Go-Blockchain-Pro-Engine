package main

import (
	"crypto/rand"
	"encoding/hex"
)

func NewWalletAddr() string {
	buf := make([]byte, 20)
	rand.Read(buf)
	return "0x" + hex.EncodeToString(buf)
}
