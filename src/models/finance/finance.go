package finance

import "main/src/entrypoints/webmodels"

type Spending struct {
	UserId   int
	Name     string
	Type     string
	Amount   int
	Currency string
}

type FactoryImp struct{}

func SpendingFromUserInput(spending webmodels.TestSpending, userId int) Spending {
	return Spending{UserId: userId, Name: spending.Name, Type: spending.Type,
		Amount:   spending.Amount,
		Currency: spending.Currency}
}

func (fi FactoryImp) CreateSpending(userId int, name string, Type string, amount int, currency string) Spending {
	return Spending{Name: name, Type: Type, UserId: userId, Amount: amount, Currency: currency}
}

type Factory interface {
	CreateSpending(userId int, name string, Type string, amount int, currency string) Spending
}
