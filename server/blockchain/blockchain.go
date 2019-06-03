package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block is...
type Block struct {
	Hash              string
	PreviousBlockHash string
	Data              string
}

// Blockchain ...
type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PreviousBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data, PreviousBlockHash string) *Block {
	block := &Block{Data: data, PreviousBlockHash: PreviousBlockHash}
	return block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	//newBlock.setHash()
	return newBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("genesis", "")
}
