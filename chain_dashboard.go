package main

import "fmt"

func ShowDashboard(chain ChainRuntime) {
	fmt.Println("=== Chain Dashboard ===")
	fmt.Println("Height:", len(chain.Blocks))
	fmt.Println("Last Hash:", chain.Blocks[len(chain.Blocks)-1].CurrentHash)
}

func CheckFork(x, y ChainRuntime) bool {
	min := len(x.Blocks)
	if len(y.Blocks) < min {
		min = len(y.Blocks)
	}
	for i := 0; i < min; i++ {
		if x.Blocks[i].CurrentHash != y.Blocks[i].CurrentHash {
			return true
		}
	}
	return false
}
