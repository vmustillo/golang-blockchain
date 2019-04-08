package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain is a structure that will hold all blocks in an array
type BlockChain struct {
	blocks []*Block
}

// Block is a structure that each block in the chain will be modeled after
// contains a hash, data, and hash of the previous block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash creates the hash for that block and assigns it to the hash field of said block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a block using the data from the parameters
// These parameters include a string of data and the hash of the previous block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
	}
	block.DeriveHash()
	return block
}

// AddBlock takes a string of data and creates a block and adds it to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis creates the first block in the chain (Genesis Block)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//InitBlockChain creates the blockchain with the Genesis block inside
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
