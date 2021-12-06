package core

// BlockChain 链接区块
type BlockChain struct {
	Blocks []*Block
}

// AddBlock 添加区块到链上
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
