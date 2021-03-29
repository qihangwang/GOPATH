package go_project

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	//版本
	Version int64
	//前区块的哈希值
	PrevBlockHash []byte
	//当前区块的哈希值
	Hash []byte
	//梅克尔根
	MerkelRoot []byte
	//时间戳
	TimeStamp int64
	//难度值
	Bits int64
	//随机值
	Nonce int64
	//交易信息
	Data []byte

}

func NewBlock(data string,PrevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version: 1,
		PrevBlockHash: PrevBlockHash,
		//Hash TODO
		MerkelRoot: []byte{},
		TimeStamp: time.Now().Unix(),
		Bits:  1,
		Nonce: 1,
		Data: []byte(data)}
	block.SetHash()
	return &block
}

func (block *Block)SetHash(){
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerkelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Nonce),
		block.Data}
	//func Join(s [][]byte, sep []byte) []byte
	data := bytes.Join(tmp,[]byte{})
	//func Sum256(data []byte) [Size]byte
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

//创世区块
func NewGenesisBlock() *Block {
	return NewBlock(data:"Genesis Block", []byte{} )
}