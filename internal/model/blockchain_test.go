package model

import (
	"testing"
)

func TestAddBlock(t *testing.T) {
	var blockchain Blockchain
	genesisBlock := createGenesisBlock()
	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)

	blockchain.AddBlock([]byte("Test block 1"), "validator1")
	blockchain.AddBlock([]byte("Test block 2"), "validator2")

	if len(blockchain.Blocks) != 3 {
		t.Errorf("Blockchain length mismatch! Expected %d but got %d", 3, len(blockchain.Blocks))
	}

	if blockchain.Blocks[1].Index != 1 || blockchain.Blocks[2].Index != 2 {
		t.Errorf("Block index mismatch! Expected %d and %d but got %d and %d", 1, 2, blockchain.Blocks[1].Index, blockchain.Blocks[2].Index)
	}

	if blockchain.Blocks[1].Hash != blockchain.Blocks[2].PrevHash {
		t.Errorf("Block hash mismatch! Expected %s and %s", blockchain.Blocks[1].Hash, blockchain.Blocks[2].PrevHash)
	}
}
