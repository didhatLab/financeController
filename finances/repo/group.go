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
		" INSERT INTO group_member (user_id, group_id) values ($3, (SELECT created_group.id FROM created_group))",
		newGroup.Name, newGroup.Description).Scan(&createdGroupId)

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

func (pgr PostgresGroupRepository) AppendMemberToGroup(ctx context.Context, groupId int, userId int) error {
	_, err := pgr.pool.Exec(ctx, "INSERT INTO group_member (user_id, group_id) values ($1, $2)", userId, groupId)

	return err
}

func (pgr PostgresGroupRepository) DeleteMemberFromGroup(ctx context.Context, groupId int, targetUserName string) error {
	_, err := pgr.pool.Exec(ctx, "DELETE FROM group_member g USING auth a WHERE a.username=$1 AND g.group_id=$2 AND a.user_id=g.user_id",
		targetUserName, groupId)

	return err

}

func (pgr PostgresGroupRepository) DeleteSpendGroup(ctx context.Context, groupId int) error {
	_, err := pgr.pool.Exec(ctx, "DELETE FROM spend_group WHERE id=$1", groupId)

	return err
}

func (pgr PostgresGroupRepository) GetUserGroups(ctx context.Context, userId int) ([]group.SpendGroup, error) {
	groups := make([]group.SpendGroup, 0, 10)

	sql := `SELECT target_user.group_id, sp.name, sp.description, member.user_id, a.username FROM (SELECT * FROM group_member WHERE user_id=$1) target_user
    		JOIN spend_group sp ON sp.id = target_user.group_id
			LEFT JOIN group_member member on sp.id = member.group_id JOIN auth a on member.user_id = a.user_id
			ORDER BY target_user.group_id`

	query, err := pgr.pool.Query(ctx, sql, userId)

	if err != nil {
		return nil, err
	}

	prevGroupId := -1

	for query.Next() {
		var groupParams group.SpendGroup
		var member group.Member

		query.Scan(&groupParams.Id, &groupParams.Name, &groupParams.Description, &member.UserId, &member.Username)
		member.IsAdmin = true
		if groupParams.Id != prevGroupId {
			prevGroupId = groupParams.Id
			groups = append(groups, groupParams)
		}
		groups[len(groups)-1].Members = append(groups[len(groups)-1].Members, member)
	}

	return groups, nil

}

type GroupRepository interface {
	CreateSpendingGroup(ctx context.Context, newGroup group.SpendGroup, userCreatorId int) (int, error)
	GetSpendingGroup(ctx context.Context, groupId int) (group.SpendGroup, error)
	AppendMemberToGroup(ctx context.Context, groupId int, userId int) error
	DeleteMemberFromGroup(ctx context.Context, groupId int, targetUserName string) error
	DeleteSpendGroup(ctx context.Context, groupId int) error
	GetUserGroups(ctx context.Context, userId int) ([]group.SpendGroup, error)
}
