package statistic

import (
	"main/finances/models/finance"
	"main/finances/models/statistic"
	"time"
)

type SpendsCounterImp struct {
}

func NewSpendsCounter() SpendsCounterImp {
	return SpendsCounterImp{}
}

func (sp SpendsCounterImp) spendsForLastMonth(spends []finance.Spending, rate statistic.CurrencyRate) statistic.PeriodSpendStat {
	var daySpend float32
	var monthSpend float32

	yesterday := time.Now().Add(-24 * time.Hour)
	monthAgo := time.Now().Add(-24 * time.Hour * 30)
	for _, sp := range spends {
		if sp.Time.After(yesterday) {
			daySpend += convertToRub(sp, rate)
		}
		if sp.Time.After(monthAgo) {
			monthSpend += convertToRub(sp, rate)
		}
	}

	return statistic.PeriodSpendStat{Day: daySpend, Month: monthSpend, Base: "RUB"}

}

func convertToRub(spend finance.Spending, rate statistic.CurrencyRate) float32 {
	if spend.Currency == "USD" {
		return rate.Usd * float32(spend.Amount)
	}
	if spend.Currency == "EUR" {
		return rate.Eur * float32(spend.Amount)
	}
	return float32(spend.Amount)
}

type SpendCounter interface {
	spendsForLastMonth(spends []finance.Spending, rate statistic.CurrencyRate) statistic.PeriodSpendStat
}
