package main

func SyncChain(chains []ChainRuntime) ChainRuntime {
	best := chains[0]
	for _, c := range chains {
		if len(c.Blocks) > len(best.Blocks) && CheckChain(c) {
			best = c
		}
	}
	return best
}
