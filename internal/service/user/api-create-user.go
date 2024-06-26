package user

import (
	"context"
	"nam_0515/internal/model"
	db "nam_0515/internal/repo/dbmodel"
	"nam_0515/pkg/util/password"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
		Address:  req.Address,
	}
	u, err := s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	s.blockChainRepo.CreateAccount(req.Address)

	return &model.CreateUserResponse{
		ID:    u.ID,
		Token: "123456", // TODO
	}, nil
}
