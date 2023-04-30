package spend

import (
	"context"
	"main/finances/models/finance"
	"main/finances/models/user"
	"main/finances/repo"
)

type GetSpendsService struct {
	repository repo.FinanceRepository
}

func NewGetSpendsService(repo repo.FinanceRepository) GetSpendsService {
	return GetSpendsService{repository: repo}
}

func (gs GetSpendsService) GetUserSpends(ctx context.Context, user user.User) (error, []finance.Spending) {
	err, spends := gs.repository.GetUserFinanceSpends(ctx, user.UserId)

	return err, spends
}
