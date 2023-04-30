package group

import (
	"context"
	"main/finances/models/group"
	"main/finances/repo"
)

type CreateSpendGroupService struct {
	spendGroupRepository repo.GroupRepository
}

func NewCreateSpendGroupService(groupRepo repo.GroupRepository) CreateSpendGroupService {
	return CreateSpendGroupService{spendGroupRepository: groupRepo}
}

func (cgs CreateSpendGroupService) CreateSpendGroup(ctx context.Context, newGroup group.SpendGroup, creatorId int) (int, error) {
	groupId, err := cgs.spendGroupRepository.CreateSpendingGroup(ctx, newGroup, creatorId)

	if err != nil {
		return -1, err
	}

	return groupId, nil

}
