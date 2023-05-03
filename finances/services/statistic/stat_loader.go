package statistic

import (
	"context"
	"encoding/json"
	"main/finances/models/statistic"
	"main/finances/repo"
	"net/http"
)

type StatLoaderService struct {
	financeRepo  repo.FinanceRepository
	spendCounter SpendCounter
}

func NewStatLoader(finRepo repo.FinanceRepository) StatLoaderService {
	counter := NewSpendsCounter()

	return StatLoaderService{financeRepo: finRepo, spendCounter: counter}
}

func (sls StatLoaderService) LoadStatsForUser(ctx context.Context, UserId int) (error, statistic.PeriodSpendStat) {
	err, userSpends := sls.financeRepo.GetUserFinanceSpends(ctx, UserId)

	res, err := http.Get("http://127.0.0.1:4002/currency_for_stats")
	var rate statistic.CurrencyRate
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&rate)

	if err != nil {
		return err, statistic.PeriodSpendStat{}
	}

	if err != nil {
		return err, statistic.PeriodSpendStat{}
	}

	return nil, sls.spendCounter.spendsForLastMonth(userSpends, rate)

}
