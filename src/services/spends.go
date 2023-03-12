package services

import (
	"context"
	"main/src/adapter/repository"
	"main/src/models/finance"
	"main/src/models/user"
)

type CreateSpendService struct {
	repository repository.FinanceRepository
}

func CreateCreateSpendsService(repo repository.FinanceRepository) CreateSpendService {
	return CreateSpendService{repository: repo}
}

func (cs CreateSpendService) CreateNewSpend(ctx context.Context, user user.User, spendName string) error {
	err := cs.repository.CreateFinanceSpending(ctx, user.UserId, spendName)

	if err != nil {
		return err
	}
	return nil
}

type GetSpendsService struct {
	repository repository.FinanceRepository
}

func CreateGetSpendsService(repo repository.FinanceRepository) GetSpendsService {
	return GetSpendsService{repository: repo}
}

func (gs GetSpendsService) GetUserSpends(ctx context.Context, user user.User) (error, []finance.Spending) {
	err, spends := gs.repository.GetUserFinanceSpends(ctx, user.UserId)

	if err != nil {
		return err, nil
	}

	return nil, spends
}

