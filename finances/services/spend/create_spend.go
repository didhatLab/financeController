package spend

import (
	"context"
	"errors"
	"main/finances/models/finance"
	"main/finances/models/user"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type CreateSpendService struct {
	repository    repo.FinanceRepository
	accessChecker privacy.GroupAccessChecker
}

func NewCreateSpendsService(repo repo.FinanceRepository, accessChecker privacy.GroupAccessChecker) CreateSpendService {
	return CreateSpendService{repository: repo, accessChecker: accessChecker}
}

func (cs CreateSpendService) CreateNewSpend(ctx context.Context, user user.User, spending finance.Spending) (error, int) {
	if spending.GroupId != nil {
		ok, err := cs.accessChecker.CheckAccessToGroupByUser(ctx, *spending.GroupId, user.UserId)
		if err != nil {
			return err, -1
		}
		if !ok {
			return errors.New("user have not access to group"), -1
		}

	}
	err, newSpendId := cs.repository.CreateFinanceSpending(ctx, user.UserId, spending)

	return err, newSpendId
}
