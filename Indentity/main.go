package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a single block in the chain
type Block struct {
	Timestamp int64
	Data      map[string]string
	PrevHash  string
	Hash      string
}

// BlockChain is a series of blocks
type BlockChain struct {
	Blocks []*Block
}

// NewBlock creates a new block and adds it to the chain
func (bc *BlockChain) NewBlock(data map[string]string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{time.Now().Unix(), data, prevBlock.Hash, ""}
	newBlock.Hash = newBlock.calculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}

// calculateHash calculates the hash of a block
func (b *Block) calculateHash() string {
	blockData := fmt.Sprintf("%d%v%s", b.Timestamp, b.Data, b.PrevHash)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// NewBlockChain creates a new blockchain with a genesis block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock creates a new genesis block
func NewGenesisBlock() *Block {
	return &Block{time.Now().Unix(), map[string]string{"identity": "genesis"}, "", ""}
}

func main() {
	// Create a new blockchain
	bc := NewBlockChain()

	// Add some blocks to the chain
	bc.NewBlock(map[string]string{"identity": "Alice"})
	bc.NewBlock(map[string]string{"identity": "Bob"})
	bc.NewBlock(map[string]string{"identity": "Charlie"})

	// Print out the blocks in the chain
	for _, block := range bc.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %v\n", block.Data)
		fmt.Printf("Prev. hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}
