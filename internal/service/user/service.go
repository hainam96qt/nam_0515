package user

import (
	"context"
	"database/sql"
	"nam_0515/internal/model"

	db "nam_0515/internal/repo/dbmodel"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		userRepo       UserRepository
		blockChainRepo BlockchainRepository
	}

	BlockchainRepository interface {
		CreateAccount(address string) *model.Account
	}

	UserRepository interface {
		CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error)
	}
)

func NewUserService(DatabaseConn *sql.DB, userRepo UserRepository, blockChainRepo BlockchainRepository) *Service {
	return &Service{
		DatabaseConn:   DatabaseConn,
		userRepo:       userRepo,
		blockChainRepo: blockChainRepo,
	}
}
