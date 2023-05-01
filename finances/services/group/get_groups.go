package group

import (
	"context"
	"main/finances/models/group"
	"main/finances/repo"
)

type GetGroupsService struct {
	groupRepo repo.GroupRepository
}

func NewGetGroupService(groupRepo repo.GroupRepository) GetGroupsService {
	return GetGroupsService{groupRepo: groupRepo}
}

func (ggs GetGroupsService) GetUserGroups(ctx context.Context, userId int) ([]group.SpendGroup, error) {
	userGroups, err := ggs.groupRepo.GetUserGroups(ctx, userId)

	return userGroups, err

}
