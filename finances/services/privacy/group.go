package privacy

import (
	"context"
	"main/finances/models/group"
	"main/finances/repo"
)

type GroupAccessCheckerImp struct {
	spendGroupRepository repo.GroupRepository
}

func NewGroupAccessChecker(spendGroupRepo repo.GroupRepository) GroupAccessCheckerImp {
	return GroupAccessCheckerImp{spendGroupRepository: spendGroupRepo}
}

func (gac GroupAccessCheckerImp) CheckAccessToGroupByUser(ctx context.Context, groupId int, userId int) (bool, error) {
	spendGroup, err := gac.spendGroupRepository.GetSpendingGroup(ctx, groupId)
	if err != nil {
		return false, err
	}
	ok := IsUserHasAccessToGroup(spendGroup, userId)

	if !ok {
		return false, nil
	}

	return true, nil
}

func IsUserHasAccessToGroup(group group.SpendGroup, userId int) bool {
	for _, member := range group.Members {
		if member.UserId == userId {
			return true
		}
	}
	return false
}

type GroupAccessChecker interface {
	CheckAccessToGroupByUser(ctx context.Context, groupId int, userId int) (bool, error)
}
