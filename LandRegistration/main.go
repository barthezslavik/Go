package main

import (
	"time"
)

// Block represents a block in the land registration blockchain
type Block struct {
	Timestamp    time.Time
	Land         *Land
	Owner        string
	PreviousHash []byte
	Hash         []byte
}

// Blockchain represents the land registration blockchain
type Blockchain struct {
	Blocks []*Block
}

// Land represents a piece of land in the land registration system
type Land struct {
	ID       string
	Location string
	Size     int
}

// AddBlock adds a new block to the blockchain
func (b *Blockchain) AddBlock(land *Land, owner string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := &Block{time.Now(), land, owner, prevBlock.Hash, []byte{}}
	newBlock.Hash = newBlock.CalculateHash()
	b.Blocks = append(b.Blocks, newBlock)
}

// CalculateHash calculates the hash of a block
func (b *Block) CalculateHash() []byte {
	// Calculate the hash of the block (omitted for simplicity)
	return []byte{}
}

func main() {
	// Create the genesis block
	genesisBlock := &Block{time.Now(), &Land{}, "", []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.CalculateHash()

	// Create the blockchain with the genesis block
	blockchain := &Blockchain{[]*Block{genesisBlock}}

	// Add a new block to the blockchain
	blockchain.AddBlock(&Land{ID: "123", Location: "New York", Size: 1000}, "Alice")

	// Add another new block to the blockchain
	blockchain.AddBlock(&Land{ID: "456", Location: "Chicago", Size: 500}, "Bob")
}
