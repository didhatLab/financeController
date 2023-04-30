package group

import (
	"context"
	"errors"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type DeleteGroupMemberService struct {
	spendGroupRepository repo.GroupRepository
	accessChecker        privacy.GroupAccessChecker
}

func NewDeleteGroupMemberSrvice(groupRepo repo.GroupRepository, accessChecker privacy.GroupAccessChecker) DeleteGroupMemberService {
	return DeleteGroupMemberService{spendGroupRepository: groupRepo, accessChecker: accessChecker}
}

func (dmr DeleteGroupMemberService) DeleteGroupMember(ctx context.Context, groupId int, deleterUserId int, userNameForDelete string) error {
	ok, err := dmr.accessChecker.CheckAccessToGroupByUser(ctx, groupId, deleterUserId)

	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user have not access for delete members")
	}

	err = dmr.spendGroupRepository.DeleteMemberFromGroup(ctx, groupId, userNameForDelete)

	return err
}
