package group

import (
	"context"
	"errors"
	"fmt"
	"main/finances/repo"
	"main/finances/services/notify"
	"main/finances/services/privacy"
)

type DeleteSpendGroupService struct {
	groupRepo     repo.GroupRepository
	accessChecker privacy.GroupAccessChecker
	notifier      notify.Notifier
}

func NewDeleteSpendGroupService(groupRepo repo.GroupRepository, accessChecker privacy.GroupAccessChecker, notifier notify.Notifier) DeleteSpendGroupService {
	return DeleteSpendGroupService{groupRepo: groupRepo, accessChecker: accessChecker, notifier: notifier}
}

func (dgr DeleteSpendGroupService) DeleteSpendGroup(ctx context.Context, groupId int, deleterId int) error {
	group, err := dgr.groupRepo.GetSpendingGroup(ctx, groupId)
	if err != nil {
		return err
	}

	ok := privacy.IsUserHasAccessToGroup(group, deleterId)

	if !ok {
		return errors.New("user have not access to delete this group")
	}

	err = dgr.groupRepo.DeleteSpendGroup(ctx, groupId)

	if err != nil {
		return err
	}
	userIds := make([]int, 0, len(group.Members))

	for _, mem := range group.Members {
		userIds = append(userIds, mem.UserId)
	}

	err = dgr.notifier.Notify(ctx, fmt.Sprintf("group %s was deleted", group.Name), userIds)

	return err
}
