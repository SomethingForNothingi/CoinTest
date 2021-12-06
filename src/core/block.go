package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 定义区块
type Block struct {
	Timestamp     int64  // 区块时间戳
	Data          []byte // 区块包含的数据
	PrevBlockHash []byte // 前一个区块的hash值
	Hash          []byte // 本身的hash值
}

// NewBlock 添加新的区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// SetHash 设置区块的hash值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewGenesisBlock 创世纪块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
