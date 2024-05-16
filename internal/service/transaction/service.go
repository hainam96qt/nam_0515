package transaction

import (
	"context"
	"nam_0515/internal/model"
)

type (
	Service struct {
		blockChainRepo BlockchainRepository
	}

	BlockchainRepository interface {
		CreateTransaction(ctx context.Context, req *model.Transaction, publicKey model.PublicKeyJSON, signature string) (*model.Transaction, error)
		ListTransactions(ctx context.Context, address string) ([]*model.Transaction, error)
	}
)

func NewTransactionService(blockChain BlockchainRepository) *Service {
	return &Service{
		blockChainRepo: blockChain,
	}
}

func (s *Service) CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.Transaction, error) {
	addressCtx := ctx.Value("Address").(string)

	return s.blockChainRepo.CreateTransaction(ctx, &model.Transaction{
		From:   addressCtx,
		To:     req.To,
		Amount: req.Amount,
	}, req.PublicKeyJSON, req.Signature)
}

func (s *Service) ListTransactions(ctx context.Context) (*model.ListTransactionResponse, error) {
	addressCtx := ctx.Value("Address").(string)

	transactions, err := s.blockChainRepo.ListTransactions(ctx, addressCtx)
	if err != nil {
		return nil, err
	}
	return &model.ListTransactionResponse{Transactions: transactions}, nil

}
