package finance

import (
	"main/finances/entrypoints/webmodels"
	"time"
)

type Spending struct {
	Id          int
	UserId      int
	Name        string
	Type        string
	Amount      int
	Currency    string
	Description string
	Time        time.Time
	GroupId     *int
}

type FactoryImp struct{}

func SpendingFromUserInput(spending webmodels.TestSpending, userId int) Spending {
	return Spending{
		UserId:      userId,
		Name:        spending.Name,
		Type:        spending.Type,
		Amount:      spending.Amount,
		Currency:    spending.Currency,
		Description: spending.Description,
		GroupId:     spending.GroupId,
	}
}

func (fi FactoryImp) CreateSpending(userId int, name string, Type string, amount int, currency string) Spending {
	return Spending{Name: name, Type: Type, UserId: userId, Amount: amount, Currency: currency}
}

type Factory interface {
	CreateSpending(userId int, name string, Type string, amount int, currency string) Spending
}
