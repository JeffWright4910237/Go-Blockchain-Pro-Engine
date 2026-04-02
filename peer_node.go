package main

import "fmt"

type PeerNode struct {
	NID    string
	Neigh  map[string]bool
	Cache  map[string]bool
}

func NewPeerNode(id string) *PeerNode {
	return &PeerNode{
		NID:   id,
		Neigh: make(map[string]bool),
		Cache: make(map[string]bool),
	}
}

func (p *PeerNode) SendBlock(hash string) {
	for n := range p.Neigh {
		fmt.Printf("Node %s send to %s | %s\n", p.NID, n, hash)
	}
}
