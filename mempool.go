package main

type MemPool struct {
	list []TxRecord
}

func NewMemPool() *MemPool {
	return &MemPool{}
}

func (m *MemPool) Push(tx TxRecord) {
	m.list = append(m.list, tx)
}

func (m *MemPool) FetchAll() []TxRecord {
	return m.list
}

func (m *MemPool) Flush() {
	m.list = nil
}
