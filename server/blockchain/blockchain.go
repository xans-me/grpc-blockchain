package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block keeps block headers
type Block struct {
	Data          string
	PrevBlockHash string
	Hash          string
}

// BlockChain struct to keeps a sequence of Blocks
type BlockChain struct {
	Blocks []*Block
}

// setHash method of block to calculates and sets block hash
func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{data, prevBlockHash, ""}
	block.setHash()

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

// AddBlock saves provided data as a block in the BlockChain
func (bc *BlockChain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// NewBlockChain creates a new BlockChain with genesis Block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
