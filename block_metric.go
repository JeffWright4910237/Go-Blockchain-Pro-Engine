package main

func AvgMineTime(chain ChainRuntime) int64 {
	if len(chain.Blocks) <= 1 {
		return 0
	}
	var total int64
	for i := 1; i < len(chain.Blocks); i++ {
		total += chain.Blocks[i].Time - chain.Blocks[i-1].Time
	}
	return total / int64(len(chain.Blocks)-1)
}
