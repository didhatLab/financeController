package spend

import (
	"context"
	"errors"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type DeleteSpendsService struct {
	repository    repo.FinanceRepository
	accessChecker privacy.GroupAccessChecker
}

func NewDeleteSpendsService(repo repo.FinanceRepository, accessChecker privacy.GroupAccessChecker) DeleteSpendsService {
	return DeleteSpendsService{repository: repo, accessChecker: accessChecker}
}

func (ds DeleteSpendsService) DeleteUserSpend(ctx context.Context, userId int, spendId int, groupId *int) error {

	if groupId != nil {
		ok, err := ds.accessChecker.CheckAccessToGroupByUser(ctx, *groupId, userId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("user have not access to group")
		}
	}

	err := ds.repository.DeleteFinanceSpending(ctx, userId, spendId, groupId)

	return err

}
