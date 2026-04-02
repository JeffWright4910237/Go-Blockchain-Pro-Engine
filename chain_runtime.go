package main

import "time"

type ChainRuntime struct {
	Blocks []CoreBlock
}

func NewChain() *ChainRuntime {
	genesis := CoreBlock{
		Height:   0,
		PrevHash: "0",
		Time:     time.Now().Unix(),
		TxRoot:   "genesis_root",
		Nonce:    0,
	}
	genesis.CurrentHash = ComputeHash(genesis)
	return &ChainRuntime{Blocks: []CoreBlock{genesis}}
}
