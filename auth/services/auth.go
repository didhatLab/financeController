package services

import (
	"context"
	"log"
	"main/auth/repo"
	"main/auth/signatory"
)

type AuthService struct {
	authRepo  repo.AuthRepository
	signToken signatory.SignService
}

func NewAuthService(authRepo repo.AuthRepository, sign signatory.SignService) AuthService {
	return AuthService{authRepo: authRepo, signToken: sign}
}

func (as AuthService) AuthUser(ctx context.Context, username string, password string) (string, int, error) {
	authUser, err := as.authRepo.GetUserDataByUserNameAndHash(ctx, username, password)

	if err != nil {
		return "", -1, err
	}
	log.Printf("auth userId: %d", authUser.UserId)

	token, err := as.signToken.SignToken(authUser.UserId)

	if err != nil {
		return "", -1, err
	}

	return token, authUser.UserId, nil
}
