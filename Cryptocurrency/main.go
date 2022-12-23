package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
)

const targetBits = 24
const maxNonce = math.MaxInt64

// Block represents a single block in the chain
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash string
	Hash          string
	Nonce         int
}

// Transaction represents a cryptocurrency transaction
type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

// BlockChain is a series of blocks
type BlockChain struct {
	Blocks []*Block
	// A map of unspent transactions
	UTXOs map[string]Transaction
	// A list of pending transactions
	PendingTransactions []*Transaction
}

// NewBlock creates a new block and adds it to the chain
func (bc *BlockChain) NewBlock(transactions []*Transaction) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{time.Now().Unix(), transactions, prevBlock.Hash, "", 0}
	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.Run()

	newBlock.Hash = hash
	newBlock.Nonce = nonce

	bc.Blocks = append(bc.Blocks, newBlock)

	// Update the UTXOs with the new transactions
	for _, t := range transactions {
		bc.UTXOs[t.Sender] = Transaction{t.Sender, t.Recipient, bc.UTXOs[t.Sender].Amount - t.Amount}
		bc.UTXOs[t.Recipient] = Transaction{t.Recipient, "", bc.UTXOs[t.Recipient].Amount + t.Amount}
	}

	return newBlock
}

// NewTransaction creates a new transaction and adds it to the next block in the chain
func (bc *BlockChain) NewTransaction(sender string, recipient string, amount int) int {
	transaction := &Transaction{sender, recipient, amount}
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
	return bc.Blocks[len(bc.Blocks)-1].Index + 1
}

// NewBlockChain creates a new blockchain with a genesis block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}, make(map[string]Transaction)}
}

// NewGenesisBlock creates a new genesis block
func NewGenesisBlock() *Block {
	return &Block{time.Now().Unix(), []*Transaction{}, "", "", 0}
}

// HashTransactions returns the hash of the transactions in a block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, t := range b.Transactions {
		transactions = append(transactions, t.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// Serialize serializes a transaction
func (t *Transaction) Serialize() []byte {
	return []byte(fmt.Sprintf("%s%s%d", t.Sender, t.Recipient, t.Amount))
}

// ProofOfWork represents the proof-of-work for mining a block
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork creates a new proof-of-work
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// Run performs the proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%v\"\n", pow.Block.Transactions)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates the proof-of-work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.Target) == -1

	return isValid
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransactions(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// IntToHex converts an int to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// MerkleTree represents a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode represents a node in a Merkle tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a list of transactions
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, dat := range data {
		node := MerkleNode{nil, nil, dat}
		nodes = append(nodes, node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := MerkleNode{&nodes[j], &nodes[j+1], nil}
			newLevel = append(newLevel, node)
		}

		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

func main() {
	// Create a new blockchain
	bc := NewBlockChain()

	// Add some transactions to the blockchain
	bc.NewTransaction("Alice", "Bob", 10)
	bc.NewTransaction("Bob", "Charlie", 5)
	bc.NewTransaction("Charlie", "Alice", 2)

	// Mine a new block
	bc.NewBlock(bc.PendingTransactions)

	// Print out the blocks in the chain
	for _, block := range bc.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	// Print out the unspent transactions
	fmt.Println("Unspent transactions:")
	for k, v := range bc.UTXOs {
		fmt.Printf("%s: %d\n", k, v.Amount)
	}
}
