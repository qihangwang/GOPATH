package main

import (
	"fmt"
	"project/study base usage"
)

func main(){
	bc := _021最新最系统区块链详解.NewblockChain()
	bc.AddBlock(data:"A sent B 1BTC")
	bc.AddBlock(data:"A sent B 1BTC")
	for _,block := range bc.blocks{
		fmt.Printf(format:"Version: %d\n",block.Version)
		fmt.Printf(format:"PrevBlockHash: %x\n",block.PrevBlockHash)
		fmt.Printf(format:"Hash: %x\n",block.Hash)
		fmt.Printf(format:"Timestamp: %d\n",block.TimeStamp)
		fmt.Printf(format:"Bits: %d\n",block.Bits)
		fmt.Printf(format:"Nonce: %d\n",block.Nonce)
		fmt.Printf(format:"Data: %s\n",block.Data)
	}
}
