package interfaces

import (
	"context"
	"main/src/models/finance"
	"main/src/models/user"
)

type CreatingSpendService interface {
	CreateNewSpend(ctx context.Context, user user.User, spendName string) error
}

type GettingUserSpendsService interface {
	GetUserSpends(ctx context.Context, user user.User) (error, []finance.Spending)
}
