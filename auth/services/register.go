package services

import (
	"context"
	"main/auth/repo"
)

type RegisterUserService struct {
	authRepo repo.AuthRepository
}

func NewRegisterUserService(authRepo repo.AuthRepository) RegisterUserService {
	return RegisterUserService{authRepo: authRepo}
}

func (rs RegisterUserService) RegisterNewUser(ctx context.Context, username string, password string) error {
	err := rs.authRepo.SaveNewUser(ctx, username, password)

	if err != nil {
		return err
	}

	return nil

}
