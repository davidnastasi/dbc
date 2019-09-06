package blockchain

// BlockChain model
type BlockChain struct {
	Blocks []*Block
}

// AddBlock add new Block
func (chain *BlockChain) AddBlock(data string) {
	prev := chain.Blocks[len(chain.Blocks)-1]
	newBlock := createBlock(data, prev.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// CreateBlockChain generate the inital
func CreateBlockChain() *BlockChain {
	return &BlockChain{[]*Block{genesis()}}
}
