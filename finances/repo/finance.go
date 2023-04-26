package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/finances/entrypoints/webmodels"
	"main/finances/models/finance"
)

type PostgresFinanceRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresFinanceRepository(pool *pgxpool.Pool) PostgresFinanceRepository {
	return PostgresFinanceRepository{
		pool: pool,
	}
}

func (pfr PostgresFinanceRepository) CreateFinanceSpending(ctx context.Context, userId int, spending finance.Spending) error {
	_, err := pfr.pool.Exec(ctx, "INSERT INTO spend(name, type, user_id, amount, currency, description) values ($1, $2, $3, $4, $5, $6)",
		spending.Name, spending.Type, userId, spending.Amount, spending.Currency, spending.Description)
	if err != nil {
		return err
	}

	return nil
}

func (pfr PostgresFinanceRepository) GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending) {
	spends := make([]finance.Spending, 0, 30)

	rows, err := pfr.pool.Query(ctx, "SELECT name, type, COALESCE(amount, 0), coalesce(currency, ''), time, id, description FROM spend WHERE user_id=$1 ORDER BY time DESC ", userId)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var sp finance.Spending
		sp.UserId = userId
		err = rows.Scan(&sp.Name, &sp.Type, &sp.Amount, &sp.Currency, &sp.Time, &sp.Id, &sp.Description)
		spends = append(spends, sp)
	}

	return nil, spends
}

func (pfr PostgresFinanceRepository) DeleteFinanceSpending(ctx context.Context, userId int, id int) error {
	_, err := pfr.pool.Exec(ctx, "DELETE FROM spend WHERE id=$1 AND user_id=$2", id, userId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pfr PostgresFinanceRepository) UpdateFinanceSpending(ctx context.Context, request webmodels.UpdateRequest, userId int) error {

	_, err := pfr.pool.Exec(ctx, "UPDATE spend "+
		"SET name=COALESCE($1, name), type=COALESCE($2, type), amount=COALESCE($3, amount), description=COALESCE($4, description)"+
		" WHERE id=$5 AND user_id=$6", request.Name, request.Type, request.Amount, request.Description, request.SpendId, userId)

	if err != nil {
		log.Println(err)
	}

	return err
}

type FinanceRepository interface {
	CreateFinanceSpending(ctx context.Context, userId int, spending finance.Spending) error
	GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending)
	DeleteFinanceSpending(ctx context.Context, userId int, id int) error
	UpdateFinanceSpending(ctx context.Context, request webmodels.UpdateRequest, userId int) error
}
