package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"testing"
)

func TestStringStorage(t *testing.T) {
	host := "http://127.0.0.1:8545"
	smartContractAddress := "0x62aA83147edC13F02BB2c9Ef297775eF6fB9494a"
	SmartContractCreatorKey := "16e7a6a2e1c67a3a3a0667f33574d5c94e399d0caba48728568cd8dd78363f0f"

	// Kết nối đến nút Ethereum
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Current block number:", blockNumber)

	contractAddress := common.HexToAddress(smartContractAddress)

	code, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(code) == 0 {
		t.Error("No contract code at given address")
	}

	contract, err := NewContracts(contractAddress, client)
	if err != nil {
		t.Error(err)
	}

	privateAdminKey, err := crypto.HexToECDSA(SmartContractCreatorKey)
	if err != nil {
		t.Fatal("private key not found", err)
	}

	authAdmin, err := bind.NewKeyedTransactorWithChainID(privateAdminKey, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("AddString", func(t *testing.T) {
		_, err := contract.AddString(authAdmin, "teststring", big.NewInt(1652985473))
		if err != nil {
			t.Fatalf("Failed to add string: %v", err)
		}
	})

	t.Run("GetString", func(t *testing.T) {
		time, err := contract.GetString(nil, "teststring")
		if err != nil {
			t.Fatalf("Failed to add string: %v", err)
		}

		assert.Equal(t, time.Int64(), int64(1652985473))
	})

	t.Run("GetString invalid", func(t *testing.T) {
		time, err := contract.GetString(nil, "teststringxxxxxx")
		if err != nil {
			t.Fatalf("Failed to add string: %v", err)
		}

		assert.Equal(t, time.Int64(), int64(0))

	})
}
