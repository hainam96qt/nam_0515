package attendance

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	contracts "nam_0515/internal/repo/smartcontract"
	"nam_0515/pkg/smartcontract"
	"time"
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

func (r *Repository) CheckIn(ctx context.Context, req *contracts.AttendanceContractAttendanceRecord) error {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	_, err := r.ethClient.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return err
	}

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return err
	}

	_, err = contract.CheckIn(auth, req.EmployeeId, req.CheckInTime, req.Details)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CheckOut(ctx context.Context, req *contracts.AttendanceContractAttendanceRecord) error {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	_, err := r.ethClient.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return err
	}

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return err
	}

	_, err = contract.CheckOut(auth, req.EmployeeId, req.Index, req.CheckOutTime, req.Details)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetListAttendanceByUserID(ctx context.Context, userID int64, from *time.Time, to *time.Time) (result []contracts.AttendanceContractAttendanceRecord, err error) {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	_, err = r.ethClient.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		return nil, err
	}

	if from != nil && to != nil {
		result, err = contract.GetAttendanceByDateRange(nil, big.NewInt(userID), big.NewInt(from.Unix()), big.NewInt(to.Unix()))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		result, err = contract.GetAttendanceByEmployeeId(nil, big.NewInt(userID))
		if err != nil {
			log.Fatal(err)
		}
	}

	return result, nil
}

func (r *Repository) UpdateAttendanceTime(ctx context.Context, c contracts.AttendanceContractAttendanceRecord) error {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	_, err := r.ethClient.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return err
	}

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return err
	}

	tx, err := contract.UpdateCheckIn(auth, c.EmployeeId, c.Index, c.CheckInTime, c.Details)
	if err != nil {
		return err
	}

	// quan ly giao dich
	_ = tx

	return nil
}
