package blockchain

// BlockChain is a structure that will hold all blocks in an array
type BlockChain struct {
	Blocks []*Block
}

// AddBlock takes a string of data and creates a block and adds it to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

//InitBlockChain creates the blockchain with the Genesis block inside
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
