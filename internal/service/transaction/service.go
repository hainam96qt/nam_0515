package transaction

import (
	"context"
	"nam_0515/internal/model"
	error2 "nam_0515/pkg/error"
	"net/http"
)

type (
	Service struct {
		blockChainRepo    BlockchainRepository
		smartContractRepo SmartContractRepository
	}

	BlockchainRepository interface {
		CreateTransaction(ctx context.Context, req *model.Transaction, publicKey model.PublicKeyJSON, signature string) (*model.Transaction, error)
		ListTransactions(ctx context.Context, address string) ([]*model.Transaction, error)
	}

	SmartContractRepository interface {
		AddHashToPublicBlockchain(ctx context.Context, hashStr string) error
	}
)

func NewTransactionService(blockChainRepo BlockchainRepository, smartContractRepo SmartContractRepository) *Service {
	return &Service{
		blockChainRepo:    blockChainRepo,
		smartContractRepo: smartContractRepo,
	}
}

func (s *Service) CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.Transaction, error) {
	addressCtx := ctx.Value("Address").(string)
	tx, err := s.blockChainRepo.CreateTransaction(ctx, &model.Transaction{
		From:   addressCtx,
		To:     req.To,
		Amount: req.Amount,
	}, req.PublicKeyJSON, req.Signature)
	if err != nil {
		return nil, err
	}

	err = s.smartContractRepo.AddHashToPublicBlockchain(ctx, tx.BlockHash)
	if err != nil {
		return nil, error2.NewXError("fail update to add hash to public blockchain", http.StatusInternalServerError)
	}

	return tx, err
}

func (s *Service) ListTransactions(ctx context.Context) (*model.ListTransactionResponse, error) {
	addressCtx := ctx.Value("Address").(string)

	transactions, err := s.blockChainRepo.ListTransactions(ctx, addressCtx)
	if err != nil {
		return nil, err
	}
	return &model.ListTransactionResponse{Transactions: transactions}, nil

}
