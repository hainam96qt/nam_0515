package hashstore

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"nam_0515/pkg/smartcontract"
)

type Repository struct {
	ethClient *ethclient.Client
	config    smartcontract.SmartContractConfig
}

func NewPostgresRepository(ethClient *ethclient.Client, config smartcontract.SmartContractConfig) *Repository {
	return &Repository{
		ethClient: ethClient,
		config:    config,
	}
}
