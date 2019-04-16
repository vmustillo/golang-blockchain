package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block is a structure that each block in the chain will be modeled after
// contains a hash, data, and hash of the previous block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates a block using the data from the parameters
// These parameters include a string of data and the hash of the previous block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
		0,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis creates the first block in the chain (Genesis Block)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize calls an encoding package "gob" to encode a block into slices of bytes
// This is because BadgerDB is a key-value store database that only allows arrays of bytes to be stored
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

// Deserialize allows a slice of bytes to be decoded using the same "gob" package that was used to encode, to return a block from a slice of bytes.
// This will be used to return the slices of bytes in the database back into blocks
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block
}
