package main

func CheckChain(chain ChainRuntime) bool {
	for i := 1; i < len(chain.Blocks); i++ {
		curr := chain.Blocks[i]
		prev := chain.Blocks[i-1]
		if curr.CurrentHash != ComputeHash(curr) {
			return false
		}
		if curr.PrevHash != prev.CurrentHash {
			return false
		}
	}
	return true
}
