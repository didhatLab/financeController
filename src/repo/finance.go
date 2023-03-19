package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/src/models/finance"
	"sync"
)

type MemoryFinanceRepository struct {
	spends map[int]finance.Spending
	lock   *sync.RWMutex

	financeFactory finance.Factory
}

var idForSpend int = 1

func NewMemoryFinanceRepository() MemoryFinanceRepository {
	return MemoryFinanceRepository{
		spends:         map[int]finance.Spending{},
		lock:           &sync.RWMutex{},
		financeFactory: finance.FactoryImp{},
	}
}

func (mfr MemoryFinanceRepository) CreateFinanceSpending(ctx context.Context, userId int, name string) error {
	mfr.lock.RLock()
	defer mfr.lock.RUnlock()

	newSpend := mfr.financeFactory.CreateSpending(userId, name, "common")
	mfr.spends[idForSpend] = newSpend
	idForSpend++
	return nil
}

func (mfr MemoryFinanceRepository) DeleteFinanceSpending(ctx context.Context, userId int, id int) error {
	mfr.lock.RLock()
	defer mfr.lock.RUnlock()

	delete(mfr.spends, id)

	return nil
}

func (mfr MemoryFinanceRepository) GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending) {
	userSpends := make([]finance.Spending, 0, 10)

	for _, spend := range mfr.spends {
		if spend.UserId == userId {
			userSpends = append(userSpends, spend)
		}
	}

	return nil, userSpends
}

type PostgresFinanceRepository struct {
	pool *pgxpool.Pool

	financeFactory finance.Factory
}

func NewPostgresFinanceRepository(pool *pgxpool.Pool) PostgresFinanceRepository {
	return PostgresFinanceRepository{
		pool:           pool,
		financeFactory: finance.FactoryImp{},
	}
}

func (pfr PostgresFinanceRepository) CreateFinanceSpending(ctx context.Context, userId int, name string) error {
	_, err := pfr.pool.Exec(ctx, "INSERT INTO spends(name, type, user_id) values ($1, $2, $3)", name, "test", userId)
	if err != nil {
		return err
	}

	return nil
}

func (pfr PostgresFinanceRepository) GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending) {
	spends := make([]finance.Spending, 0, 30)

	rows, err := pfr.pool.Query(ctx, "SELECT name, type FROM spends WHERE user_id=$1", userId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var sp finance.Spending
		sp.UserId = userId
		err = rows.Scan(&sp.Name, &sp.Type)
		spends = append(spends, sp)
	}

	return nil, spends
}

func (pfr PostgresFinanceRepository) DeleteFinanceSpending(ctx context.Context, userId int, id int) error {
	return nil
}

type FinanceRepository interface {
	CreateFinanceSpending(ctx context.Context, userId int, name string) error
	DeleteFinanceSpending(ctx context.Context, userId int, id int) error
	GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending)
}
