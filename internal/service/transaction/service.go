package transaction

import (
	"context"
	"database/sql"
	"nam_0515/internal/model"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		blockChain BlockchainRepository
	}

	BlockchainRepository interface {
		CreateTransaction(ctx context.Context, req *model.Transaction) (*model.Transaction, error)
	}
)

func NewTransactionService(DatabaseConn *sql.DB, blockChain BlockchainRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
	}
}
