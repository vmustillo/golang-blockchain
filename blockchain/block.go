package blockchain

// BlockChain is a structure that will hold all blocks in an array
type BlockChain struct {
	Blocks []*Block
}

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

// AddBlock takes a string of data and creates a block and adds it to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis creates the first block in the chain (Genesis Block)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//InitBlockChain creates the blockchain with the Genesis block inside
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
