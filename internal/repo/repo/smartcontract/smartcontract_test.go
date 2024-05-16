package smartcontract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	contracts "nam_0515/internal/repo/smartcontract"
	"nam_0515/pkg/smartcontract"
	"testing"
)

func TestContractFunctionality(t *testing.T) {
	host := "http://127.0.0.1:7545"
	chainID := 1337
	SmartContractCreatorKey := "96a79ec829cdc73f5a52cda6a34ef058dedafb3bfc79c01c0936c2f417e954f1"

	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}

	repo := NewSmartContractRepository(client, smartcontract.SmartContractConfig{
		Host:       host,
		PrivateKey: SmartContractCreatorKey,
		ChainID:    chainID,
	})

	contractAddress, err := repo.DeployContract(6721975)
	if err != nil {
		t.Fatal(err)
	}

	code, err := client.CodeAt(context.Background(), *contractAddress, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(code) == 0 {
		t.Fatal("No contract code at given address")
	}

	privateKey, err := crypto.HexToECDSA(SmartContractCreatorKey)
	if err != nil {
		t.Fatal("private key not found", err)
	}

	// Tạo một thực thể hợp đồng từ địa chỉ
	contract, err := contracts.NewContracts(*contractAddress, client)
	if err != nil {
		t.Fatal(err)
	}

	_ = contract

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}

	_ = auth

	// TODO got something wrong with golang deploy. temporary using remix IDE fix later
	// VM Exception while processing transaction: invalid opcode
	_, err = contract.GetString(nil, "abc xyz")
	if err != nil {
		t.Fatal(err)
	}

	//// test with remix ide deploy
	//contractAddress2 := common.HexToAddress("0x57aca647E5D0F71220434758807307C0B74969B2")
	//contract2, err := contracts.NewContracts(contractAddress2, client)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//_ = contract2

	//err = repo.ListenEvent(*contractAddress)
	//if err != nil {
	//	t.Fatal(err)
	//}

}
