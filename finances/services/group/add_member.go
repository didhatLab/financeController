package group

import (
	"context"
	"errors"
	"log"
	"main/finances/repo"
	"main/finances/services/notify"
	"main/finances/services/privacy"
)

type AddGroupMemberService struct {
	groupRepository repo.GroupRepository
	accessChecker   privacy.GroupAccessChecker
	notifier        notify.Notifier
}

func NewAddGroupMemberService(groupRepo repo.GroupRepository, accessChecker privacy.GroupAccessChecker, notifier notify.Notifier) AddGroupMemberService {
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

	if err != nil {
		return err
	}

	err = ags.notifier.Notify(ctx, "you was added to group", []int{targetUserId})

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
