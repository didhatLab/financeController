package spend

import (
	"context"
	"main/finances/models/finance"
	"main/finances/models/user"
)

type CreatingSpendService interface {
	CreateNewSpend(ctx context.Context, user user.User, spending finance.Spending) (error, int)
}

type GettingUserSpendsService interface {
	GetUserSpends(ctx context.Context, user user.User) (error, []finance.Spending)
}
