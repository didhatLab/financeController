package group

import (
	"context"
	"errors"
	"main/finances/repo"
	"main/finances/services/privacy"
)

type AddGroupMemberService struct {
	groupRepository repo.GroupRepository
	accessChecker   privacy.GroupAccessChecker
}

func NewAddGroupMemberService(groupRepo repo.GroupRepository, accessChecker privacy.GroupAccessChecker) AddGroupMemberService {
	return AddGroupMemberService{groupRepository: groupRepo, accessChecker: accessChecker}
}

func (ags AddGroupMemberService) AddMemberToGroup(ctx context.Context, adderUserId int, targetUserId int, targetGroupId int) error {
	ok, err := ags.accessChecker.CheckAccessToGroupByUser(ctx, targetGroupId, adderUserId)

	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user have not access to this group")
	}

	err = ags.groupRepository.AppendMemberToGroup(ctx, targetGroupId, targetUserId)

	return err
}
