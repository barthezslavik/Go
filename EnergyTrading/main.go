package main

import (
	"time"
)

// Block represents a block in the energy trading blockchain
type Block struct {
	Timestamp    time.Time
	Energy       int
	Sender       string
	Recipient    string
	PreviousHash []byte
	Hash         []byte
}

// Blockchain represents the energy trading blockchain
type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block to the blockchain
func (b *Blockchain) AddBlock(energy int, sender string, recipient string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := &Block{time.Now(), energy, sender, recipient, prevBlock.Hash, []byte{}}
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
	genesisBlock := &Block{time.Now(), 0, "", "", []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.CalculateHash()

	// Create the blockchain with the genesis block
	blockchain := &Blockchain{[]*Block{genesisBlock}}

	// Add a new block to the blockchain
	blockchain.AddBlock(100, "Alice", "Bob")

	// Add another new block to the blockchain
	blockchain.AddBlock(50, "Bob", "Charlie")
}
