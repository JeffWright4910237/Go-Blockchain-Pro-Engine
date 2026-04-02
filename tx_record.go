package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type TxRecord struct {
	TxID    string
	From    string
	To      string
	Value   float64
	Stamp   int64
}

func CreateTx(from, to string, val float64) TxRecord {
	stamp := time.Now().Unix()
	data := from + to + strconv.FormatFloat(val, 'f', 2, 64) + strconv.FormatInt(stamp, 10)
	hash := sha256.Sum256([]byte(data))
	return TxRecord{
		TxID:  hex.EncodeToString(hash[:]),
		From:  from,
		To:    to,
		Value: val,
		Stamp: stamp,
	}
}
