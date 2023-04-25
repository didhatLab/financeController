package services

import (
	"context"
	"main/finances/entrypoints/webmodels"
	"main/finances/repo"
)

type UpdateSpendsService struct {
	repository repo.FinanceRepository
}

func NewUpdateSpendsService(repo repo.PostgresFinanceRepository) UpdateSpendsService {
	return UpdateSpendsService{repository: repo}
}

func (us UpdateSpendsService) UpdateUserSpend(ctx context.Context, req webmodels.UpdateRequest, userId int) error {
	err := us.repository.UpdateFinanceSpending(ctx, req, userId)

	return err
}
