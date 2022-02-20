package main

import (
	"fmt"

	"github.com/davoodmood/go-vector-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("4th Block")
	chain.AddBlock("5th Block")

	for _, block := range chain.Blocks {
		fmt.Printf("\nPrevious Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
