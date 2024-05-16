package model

import (
	"bytes"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	block := Block{
		Index:     1,
		Timestamp: "2022-01-01 00:00:00",
		Data:      []byte{},
		PrevHash:  "0000000000000000",
		Validator: "",
	}

	expectedHash := calculateHash(&block)
	block.Hash = expectedHash

	if block.Hash != expectedHash {
		t.Errorf("Hash wrong expected %s but got %s", expectedHash, block.Hash)
	}
}

func TestGenerateBlock(t *testing.T) {
	oldBlock := Block{
		Index:     0,
		Timestamp: "2022-01-01 00:00:00",
		Data:      []byte("hello world 1"),
		PrevHash:  "",
		Hash:      calculateHash(&Block{0, "2022-01-01 00:00:00", []byte("hello world 1"), "", "", "genesis"}),
		Validator: "genesis",
	}

	data := []byte("Test Block")
	newBlock := createBlock(&oldBlock, data, "validator 01")

	if newBlock.Index != oldBlock.Index+1 {
		t.Errorf("Index mismatch! Expected %d but got %d", oldBlock.Index+1, newBlock.Index)
	}

	if newBlock.PrevHash != oldBlock.Hash {
		t.Errorf("PrevHash wrong! Expected %s but got %s", oldBlock.Hash, newBlock.PrevHash)
	}

	if bytes.Equal(newBlock.Data, data) {
		t.Errorf("Data mismatch! Expected %s but got %s", data, newBlock.Data)
	}

	if newBlock.Hash != calculateHash(newBlock) {
		t.Errorf("Hash mismatch! Expected %s but got %s", calculateHash(newBlock), newBlock.Hash)
	}
}
