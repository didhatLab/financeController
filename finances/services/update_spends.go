package services

import (
	"context"
	"errors"
	"main/finances/entrypoints/webmodels"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type UpdateSpendsService struct {
	repository    repo.FinanceRepository
	accessChecker privacy.GroupAccessChecker
}

func NewUpdateSpendsService(repo repo.PostgresFinanceRepository, accessChecker privacy.GroupAccessChecker) UpdateSpendsService {
	return UpdateSpendsService{repository: repo, accessChecker: accessChecker}
}

func (us UpdateSpendsService) UpdateUserSpend(ctx context.Context, req webmodels.UpdateRequest, userId int) error {
	if req.GroupId != nil {
		ok, err := us.accessChecker.CheckAccessToGroupByUser(ctx, *req.GroupId, userId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("user have not access to group")
		}
	}

	err := us.repository.UpdateFinanceSpending(ctx, req, userId)

	return err
}
