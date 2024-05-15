package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      []byte `json:"data"`
	PrevHash  string `json:"prevHash"`
	Hash      string `json:"hash"`
	Validator string `json:"validator"`
}

func calculateHash(block *Block) string {
	record := fmt.Sprintf("%s%v%v%v%s", strconv.Itoa(block.Index), block.Timestamp, block.Data, block.PrevHash, block.Validator)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createBlock(prevBlock *Block, data []byte, validator string) *Block {
	block := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Validator: validator,
		Hash:      "",
	}
	block.Hash = calculateHash(block)
	return block
}

func createGenesisBlock() *Block {
	return &Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      []byte("Genesis Block"),
		PrevHash:  "",
		Validator: "Genesis",
		Hash:      "",
	}
}
