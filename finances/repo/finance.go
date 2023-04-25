package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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
	_, err := pfr.pool.Exec(ctx, "INSERT INTO spend(name, type, user_id, amount, currency) values ($1, $2, $3, $4, $5)",
		spending.Name, spending.Type, userId, spending.Amount, spending.Currency)
	if err != nil {
		return err
	}

	return nil
}

func (pfr PostgresFinanceRepository) GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending) {
	spends := make([]finance.Spending, 0, 30)

	rows, err := pfr.pool.Query(ctx, "SELECT name, type, COALESCE(amount, 0), coalesce(currency, ''), time FROM spend WHERE user_id=$1", userId)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var sp finance.Spending
		sp.UserId = userId
		err = rows.Scan(&sp.Name, &sp.Type, &sp.Amount, &sp.Currency, &sp.Time)
		spends = append(spends, sp)
	}

	return nil, spends
}

func (pfr PostgresFinanceRepository) DeleteFinanceSpending(ctx context.Context, userId int, id int) error {
	return nil
}

type FinanceRepository interface {
	CreateFinanceSpending(ctx context.Context, userId int, spending finance.Spending) error
	GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending)
	DeleteFinanceSpending(ctx context.Context, userId int, id int) error
}
