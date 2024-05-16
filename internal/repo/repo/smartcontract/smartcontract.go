package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	contracts "nam_0515/internal/repo/smartcontract"
	"nam_0515/pkg/smartcontract"
	"strings"
)

type Repository struct {
	ethClient *ethclient.Client
	config    smartcontract.SmartContractConfig
}

func NewSmartContractRepository(ethClient *ethclient.Client, config smartcontract.SmartContractConfig) *Repository {
	return &Repository{
		ethClient: ethClient,
		config:    config,
	}
}

func (r *Repository) ListenEvent(contractAddress common.Address) error {
	// Create a new instance of the smart contract
	contractInstance, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to the event
	logs := make(chan *contracts.ContractsStringAdded)
	sub, err := contractInstance.WatchStringAdded(nil, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			return err
		case event := <-logs:
			fmt.Println("Received event:", event)
			return nil
		}
	}
}

func (r *Repository) DeployContract(gasLimit uint64) (*common.Address, error) {
	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(int64(r.config.ChainID))

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := r.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := r.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to retrieve nonce: %v", err)
	}

	abiValue, err := abi.JSON(strings.NewReader(contracts.ContractsMetaData.ABI))
	if err != nil {
		return nil, err
	}

	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)

	contractAddress, tx, _, err := bind.DeployContract(auth, abiValue, common.FromHex(contracts.ContractsMetaData.Bin), r.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy the contract: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), r.ethClient, tx)
	if err != nil {
		return nil, err
	}

	contractAddress = receipt.ContractAddress

	return &contractAddress, nil
}
