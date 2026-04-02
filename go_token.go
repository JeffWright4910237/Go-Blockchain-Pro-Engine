package main

type GoToken struct {
	Name     string
	Symbol   string
	MaxSupply float64
	Holders  map[string]float64
}

func InitToken() *GoToken {
	return &GoToken{
		Name:     "GoNativeCoin",
		Symbol:   "GNC",
		MaxSupply: 21000000,
		Holders:  make(map[string]float64),
	}
}

func (t *GoToken) Send(from, to string, amt float64) bool {
	if t.Holders[from] < amt {
		return false
	}
	t.Holders[from] -= amt
	t.Holders[to] += amt
	return true
}
