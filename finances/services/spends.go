package services

import (
	"context"
	"main/finances/models/finance"
	"main/finances/models/user"
	"main/finances/repo"
)

type CreateSpendService struct {
	repository repo.FinanceRepository
}

func CreateCreateSpendsService(repo repo.FinanceRepository) CreateSpendService {
	return CreateSpendService{repository: repo}
}

func (cs CreateSpendService) CreateNewSpend(ctx context.Context, user user.User, spending finance.Spending) error {
	err := cs.repository.CreateFinanceSpending(ctx, user.UserId, spending)

	if err != nil {
		return err
	}
	return nil
}

type GetSpendsService struct {
	repository repo.FinanceRepository
}

func CreateGetSpendsService(repo repo.FinanceRepository) GetSpendsService {
	return GetSpendsService{repository: repo}
}

func (gs GetSpendsService) GetUserSpends(ctx context.Context, user user.User) (error, []finance.Spending) {
	err, spends := gs.repository.GetUserFinanceSpends(ctx, user.UserId)

	if err != nil {
		return err, nil
	}

	return nil, spends
}
