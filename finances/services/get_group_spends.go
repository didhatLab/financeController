package services

import (
	"context"
	"errors"
	"main/finances/models/finance"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type GetGroupSpendsService struct {
	financeRepository repo.FinanceRepository
	accessChecker     privacy.GroupAccessChecker
}

func NewGroupSpendsService(financeRepo repo.FinanceRepository, accessChecker privacy.GroupAccessChecker) GetGroupSpendsService {
	return GetGroupSpendsService{financeRepository: financeRepo, accessChecker: accessChecker}
}

func (gss GetGroupSpendsService) GetGroupSpends(ctx context.Context, groupId int, userId int) ([]finance.Spending, error) {
	ok, err := gss.accessChecker.CheckAccessToGroupByUser(ctx, groupId, userId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("user have not access to group")
	}
	spends, err := gss.financeRepository.GetGroupFinanceSpends(ctx, groupId)

	if err != nil {
		return nil, err
	}

	return spends, nil

}
