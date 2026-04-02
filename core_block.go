package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type CoreBlock struct {
	Height      int
	PrevHash    string
	Time        int64
	TxRoot      string
	Nonce       int
	CurrentHash string
}

func ComputeHash(b CoreBlock) string {
	raw := strconv.Itoa(b.Height) + b.PrevHash + strconv.FormatInt(b.Time, 10) + b.TxRoot + strconv.Itoa(b.Nonce)
	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])
}
