package repository

import (
	"context"
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

type FinanceRepository interface {
	CreateFinanceSpending(ctx context.Context, userId int, name string) error
	DeleteFinanceSpending(ctx context.Context, userId int, id int) error
	GetUserFinanceSpends(ctx context.Context, userId int) (error, []finance.Spending)
}
