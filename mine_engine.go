package main

import "strings"

const MineDifficulty = 4

func RunMine(block CoreBlock) int {
	target := strings.Repeat("0", MineDifficulty)
	nonce := 0
	for {
		block.Nonce = nonce
		if strings.HasPrefix(ComputeHash(block), target) {
			return nonce
		}
		nonce++
	}
}
