package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"main/finances/models/group"
)

type PostgresGroupRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresGroupRepository(pool *pgxpool.Pool) PostgresGroupRepository {
	return PostgresGroupRepository{pool: pool}
}

func (pgr PostgresGroupRepository) CreateSpendingGroup(ctx context.Context, newGroup group.SpendGroup, userCreatorId int) (int, error) {
	var createdGroupId int

	err := pgr.pool.QueryRow(ctx, "WITH created_group AS (INSERT INTO spend_group (name, description) values ($1, $2) RETURNING id"+
		" INSERT INTO group_member (user_id, group_id) values ($3, (SELECT created_group.id FROM created_group))", newGroup.Name, newGroup.Description).Scan(&createdGroupId)

	return createdGroupId, err
}

func (pgr PostgresGroupRepository) GetSpendingGroup(ctx context.Context, groupId int) (group.SpendGroup, error) {
	members := make([]group.Member, 0, 10)
	var spendGroup group.SpendGroup

	query, err := pgr.pool.Query(ctx, "SELECT user_id, id, name, description FROM spend_group LEFT JOIN group_member gm on spend_group.id = gm.group_id WHERE id=$1", groupId)

	if err != nil {
		return group.SpendGroup{}, err
	}

	if !query.Next() {
		return group.SpendGroup{}, errors.New("group does not exist")
	}
	var memberId *int

	query.Scan(&memberId, &spendGroup.Id, &spendGroup.Name, &spendGroup.Description)

	if memberId != nil {
		members = append(members, group.NewGroupMember(*memberId, false))
	}

	for query.Next() {
		query.Scan(&memberId)
		if memberId != nil {
			members = append(members, group.NewGroupMember(*memberId, false))
		}
	}

	spendGroup.Members = members

	return spendGroup, nil
}

type GroupRepository interface {
	CreateSpendingGroup(ctx context.Context, newGroup group.SpendGroup, userCreatorId int) (int, error)
	GetSpendingGroup(ctx context.Context, groupId int) (group.SpendGroup, error)
}
