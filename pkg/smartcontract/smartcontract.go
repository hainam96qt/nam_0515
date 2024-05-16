package smartcontract

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type SmartContractConfig struct {
	Host            string `yaml:"host"`
	ContractAddress string `yaml:"contractAddress"`
	PrivateKey      string `yaml:"privateKey"`
	ChainID         string `yaml:"chainID"`
}

func ConnectSmartContract(args SmartContractConfig) (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}
