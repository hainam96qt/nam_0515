package smartcontract

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"nam_0515/pkg/smartcontract"
	"testing"
)

func TestContractFunctionality(t *testing.T) {
	host := "http://127.0.0.1:8545"
	SmartContractCreatorKey := "e077e5186405e47b50db3544b4762690385a3f60492bca028df83c3918fdfe1d"

	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}

	repo := NewSmartContractRepository(client, smartcontract.SmartContractConfig{
		Host:       host,
		PrivateKey: SmartContractCreatorKey,
	})

	contractAddress, err := repo.DeployContract(1000000)
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

	//privateKey, err := crypto.HexToECDSA(SmartContractCreatorKey)
	//if err != nil {
	//	log.Fatal("private key not found", err)
	//}
	//
	//chainID, err := repo.ethClient.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//

}
