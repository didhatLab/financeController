package group

import (
	"context"
	"errors"
	"fmt"
	"main/finances/models/group"
	"main/finances/repo"
	"main/finances/services/notify"
	"main/finances/services/privacy"
)

type DeleteGroupMemberService struct {
	spendGroupRepository repo.GroupRepository
	accessChecker        privacy.GroupAccessChecker
	notifier             notify.Notifier
}

func NewDeleteGroupMemberSrvice(groupRepo repo.GroupRepository, accessChecker privacy.GroupAccessChecker, notifier notify.Notifier) DeleteGroupMemberService {
	return DeleteGroupMemberService{spendGroupRepository: groupRepo, accessChecker: accessChecker, notifier: notifier}
}

func (dgr DeleteGroupMemberService) DeleteGroupMember(ctx context.Context, groupId int, deleterUserId int, userNameForDelete string) error {
	groupForDelete, err := dgr.spendGroupRepository.GetSpendingGroup(ctx, groupId)

	if err != nil {
		return err
	}
	ok := privacy.IsUserHasAccessToGroup(groupForDelete, deleterUserId)

	if !ok {
		return errors.New("user have not access for delete members")
	}
	userIdForDelete := getUserForDeleteByName(groupForDelete, userNameForDelete)
	err = dgr.spendGroupRepository.DeleteMemberFromGroup(ctx, groupId, userNameForDelete)

	err = dgr.notifier.Notify(ctx, fmt.Sprintf("you was deleted from group %s", groupForDelete.Name), []int{userIdForDelete})

	return err
}

func getUserForDeleteByName(groupForDelete group.SpendGroup, username string) int {

	for _, mem := range groupForDelete.Members {
		if mem.Username == username {
			return mem.UserId
		}
	}

	return -1
}
