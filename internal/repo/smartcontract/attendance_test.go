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
	"time"
)

func TestAttendanceContract(t *testing.T) {
	host := "http://127.0.0.1:8545"
	smartContractAddress := "0xC71310988e50927AfCA4e8c59333eCa46959a7ED"
	SmartContractCreatorKey := "e077e5186405e47b50db3544b4762690385a3f60492bca028df83c3918fdfe1d"
	SmartContractUserKey := "18567cb7d31940156181848057e04b990ea2ca40f7616fd4f42e60fd9b3d3b7e"

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

	t.Run("AuthorizeEntity", func(t *testing.T) {
		privateUserKey, err := crypto.GenerateKey()
		if err != nil {
			t.Fatalf("Failed to generate private key: %v", err)
		}

		sender := crypto.PubkeyToAddress(privateUserKey.PublicKey)
		_, err = contract.AuthorizeEntity(authAdmin, sender, big.NewInt(1))
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("User Not AuthorizeEntity", func(t *testing.T) {
		privateUserKey, err := crypto.GenerateKey()
		if err != nil {
			t.Fatalf("Failed to generate private key: %v", err)
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateUserKey, big.NewInt(1337))
		if err != nil {
			t.Fatal(err)
		}

		checkInTime := time.Now().Unix()
		_, err = contract.CheckIn(auth, big.NewInt(1), big.NewInt(checkInTime), "nv1")
		if err != nil {
			assert.Equal(t, "VM Exception while processing transaction: revert Unauthorized access", err.Error())
		}
	})

	t.Run("User did AuthorizeEntity Check In", func(t *testing.T) {
		privateUserKey, err := crypto.HexToECDSA(SmartContractUserKey)
		if err != nil {
			t.Fatalf("Failed to generate private key: %v", err)
		}

		sender := crypto.PubkeyToAddress(privateUserKey.PublicKey)

		_, err = contract.AuthorizeEntity(authAdmin, sender, big.NewInt(2))
		if err != nil {
			t.Fatal(err)
		}

		authUser, err := bind.NewKeyedTransactorWithChainID(privateUserKey, big.NewInt(1337))
		if err != nil {
			t.Fatal(err)
		}

		checkInTime := time.Now().Unix()
		_, err = contract.CheckIn(authUser, big.NewInt(2), big.NewInt(checkInTime), "nv2")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("User did AuthorizeEntity Check In another user", func(t *testing.T) {
		privateUserKey, err := crypto.HexToECDSA(SmartContractUserKey)
		if err != nil {
			t.Fatalf("Failed to generate private key: %v", err)
		}

		sender := crypto.PubkeyToAddress(privateUserKey.PublicKey)

		_, err = contract.AuthorizeEntity(authAdmin, sender, big.NewInt(2))
		if err != nil {
			t.Fatal(err)
		}

		authUser, err := bind.NewKeyedTransactorWithChainID(privateUserKey, big.NewInt(1337))
		if err != nil {
			t.Fatal(err)
		}

		checkInTime := time.Now().Unix()
		_, err = contract.CheckIn(authUser, big.NewInt(1), big.NewInt(checkInTime), "nv1")
		if err != nil {
			assert.Equal(t, "VM Exception while processing transaction: revert Unauthorized access", err.Error())
		}
	})
	
	// same for other case
}
