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

	return err
}

type GetSpendsService struct {
	repository repo.FinanceRepository
}

func CreateGetSpendsService(repo repo.FinanceRepository) GetSpendsService {
	return GetSpendsService{repository: repo}
}

func (gs GetSpendsService) GetUserSpends(ctx context.Context, user user.User) (error, []finance.Spending) {
	err, spends := gs.repository.GetUserFinanceSpends(ctx, user.UserId)

	return err, spends
}

type DeleteSpendsService struct {
	repository repo.FinanceRepository
}

func CreateDeleteSpendsService(repo repo.FinanceRepository) DeleteSpendsService {
	return DeleteSpendsService{repository: repo}
}

func (ds DeleteSpendsService) DeleteUserSpend(ctx context.Context, userId int, spendId int) error {
	err := ds.repository.DeleteFinanceSpending(ctx, userId, spendId)

	return err

}
