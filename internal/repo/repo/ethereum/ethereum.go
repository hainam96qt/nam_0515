package ethereum

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nam_0515/pkg/smartcontract"
)

type Repository struct {
	ethClient *ethclient.Client
	config    smartcontract.SmartContractConfig
}

func NewEthereumRepository(ethClient *ethclient.Client, config smartcontract.SmartContractConfig) *Repository {
	return &Repository{
		ethClient: ethClient,
		config:    config,
	}
}

func (r *Repository) GetLatestBlock(ctx context.Context) (*types.Block, error) {
	blockNumber, err := r.ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	block, err := r.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (r *Repository) GetBalanceAddress(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := r.ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (r *Repository) CreateTransactionWithPrivateKey(ctx context.Context, privateKeyHex string, toAddressHex string, amount int, gasLimit int) error {
	//Sender's private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Sender's address derived from the private key
	fromAddress := crypto.PubkeyToAddress(privateKey.Public().(ecdsa.PublicKey))

	// Recipient's address
	toAddress := common.HexToAddress(toAddressHex)

	// Fetch the nonce for the sender's address
	nonce, err := r.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Specify the amount to send
	value := big.NewInt(int64(amount)) // 1 ether in wei

	// Gas limit and gas price
	gasPrice, err := r.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      uint64(gasLimit),
		GasPrice: gasPrice,
		Data:     nil,
	})

	// ChainID retrieval
	chainID, err := r.ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Send the transaction
	err = r.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *Repository) CreateSignedTransaction(ctx context.Context, signedTx *types.Transaction) error {
	// Send the transaction
	err := r.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	return nil
}
