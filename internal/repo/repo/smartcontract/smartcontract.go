package smartcontract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	contracts "nam_0515/internal/repo/smartcontract"
	"nam_0515/pkg/smartcontract"
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

func (r *Repository) DeployContract(gasLimit uint64) (*common.Address, error) {
	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := r.ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// Deploy the contract using the transactor
	_, tx, _, err := contracts.DeployContracts(auth, r.ethClient)
	if err != nil {
		return nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), r.ethClient, tx)
	if err != nil {
		return nil, err
	}

	// Get the contract address from the receipt
	contractAddress := receipt.ContractAddress

	return &contractAddress, nil

	//privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//chainID, err := r.ethClient.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Get the nonce for the deployer address
	//nonce, err := r.ethClient.PendingNonceAt(context.Background(), auth.From)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Get the gas price
	//gasPrice, err := r.ethClient.SuggestGasPrice(context.Background())
	//if err != nil {
	//	return nil, err
	//}
	//
	//bytecode, err := os.ReadFile("/internal/repo/smartcontract/attendance.bin")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Create the transaction to deploy the contract
	//tx := types.NewTx(&types.LegacyTx{
	//	Nonce:    nonce,
	//	To:       nil,
	//	Value:    big.NewInt(0),
	//	Gas:      gasLimit,
	//	GasPrice: gasPrice,
	//	Data:     bytecode,
	//})
	//
	//// Sign the transaction
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), privateKey)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Send the transaction
	//err = r.ethClient.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Wait for the transaction to be mined
	//receipt, err := bind.WaitMined(context.Background(), r.ethClient, signedTx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Get the contract address from the receipt
	//contractAddress := receipt.ContractAddress
	//
	//return &contractAddress, nil
}
