package main

type ContractState struct {
	Addr  string
	Store map[string]string
}

type ContractEngine struct {
	contracts map[string]*ContractState
}

func NewContractEngine() *ContractEngine {
	return &ContractEngine{contracts: make(map[string]*ContractState)}
}

func (c *ContractEngine) Deploy() *ContractState {
	addr := NewWalletAddr()
	cs := &ContractState{Addr: addr, Store: make(map[string]string)}
	c.contracts[addr] = cs
	return cs
}
