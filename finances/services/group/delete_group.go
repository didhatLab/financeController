package group

import (
	"context"
	"errors"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type DeleteSpendGroupService struct {
	groupRepo     repo.GroupRepository
	accessChecker privacy.GroupAccessChecker
}

func NewDeleteSpendGroupService(groupRepo repo.GroupRepository, accessChecker privacy.GroupAccessChecker) DeleteSpendGroupService {
	return DeleteSpendGroupService{groupRepo: groupRepo, accessChecker: accessChecker}
}

func (dgr DeleteSpendGroupService) DeleteSpendGroup(ctx context.Context, groupId int, deleterId int) error {
	ok, err := dgr.accessChecker.CheckAccessToGroupByUser(ctx, groupId, deleterId)

	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user have not access to delete this group")
	}

	err = dgr.groupRepo.DeleteSpendGroup(ctx, groupId)

	return err
}
