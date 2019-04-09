package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Take the data from the block

// Create a councer (nonce) which starts at 0

// Create a hash of the data plus the counter

// Check the hash to see if it meets a set of requirements

/*
	Requirements:
		The first few bytes must contain 0s
*/

// Difficulty is an arbitrary value (staying static for now)
// Represents the amount of zero bits that start the hash
// Would increase in a real blockchain to account for an increase in the number of miners and to make it harder to mine blocks
const Difficulty = 12

// ProofOfWork is a structure that contains a block and the target. When the target is hit, the block can be signed and added the chain
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof returns a ProofOfWork with a block and target
// Left shifts the target 244 bits
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData derives the hash for the block inside the Proof of Work
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// Run outputs an integer and slice of bytes inside of a tuple
// Run is virtually an infinite loop that joins the nonce, difficuty, previous hash and data and then hashes it using sha256
// This hash is then converted to a big int and compared to the target of the block, and if it is less, than the function increments nonce and runs again
// If the hash is equal to the target, the hash and the nonce (incremented) are returned
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate validates that the block met the target by calling InitData on the blocks nonce and comparing it to its target
// This can be uised to quickly validate the proof of work of any block
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex is a helper function that takes an integer and writes it to a slice of bytes in Big Endian format
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
