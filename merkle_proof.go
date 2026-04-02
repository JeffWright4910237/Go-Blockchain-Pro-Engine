package main

import "crypto/sha256"

func MerkleProof(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}
	for len(hashes) > 1 {
		var next [][]byte
		for i := 0; i < len(hashes); i += 2 {
			left := hashes[i]
			var right []byte
			if i+1 < len(hashes) {
				right = hashes[i+1]
			} else {
				right = left
			}
			combined := append(left, right...)
			h := sha256.Sum256(combined)
			next = append(next, h[:])
		}
		hashes = next
	}
	return hashes[0]
}
