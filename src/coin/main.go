package main

import (
	core "CoinTest/src/core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()

	bc.AddBlock("Send 1 BTC to my")
	bc.AddBlock("Send 2 BTC to my")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash： %x\n", block.PrevBlockHash)
		fmt.Printf("Data： %s\n", block.Data)
		fmt.Printf("Hash： %x\n", block.Hash)
		fmt.Println()
	}

}
