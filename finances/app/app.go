package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"main/finances/entrypoints"
	"main/finances/repo"
	"main/finances/services/privacy"
	"main/finances/services/spend"
	"net/http"
)

type App struct {
	AppMux *http.ServeMux
}

func NewApplication(ctx context.Context, pool *pgxpool.Pool) (error, App) {

	groupRepo := repo.NewPostgresGroupRepository(pool)
	finRepo := repo.NewPostgresFinanceRepository(pool)

	groupAccessChecker := privacy.NewGroupAccessChecker(groupRepo)

	finEntry := entrypoints.FinanceEntryPoint{
		CreateSpendService:    spend.NewCreateSpendsService(finRepo, groupAccessChecker),
		GetSpendsService:      spend.NewGetSpendsService(finRepo),
		DeleteSpendService:    spend.NewDeleteSpendsService(finRepo, groupAccessChecker),
		UpdateSpendService:    spend.NewUpdateSpendsService(finRepo, groupAccessChecker),
		GetGroupSpendsService: spend.NewGroupSpendsService(finRepo, groupAccessChecker),
		Ctx:                   ctx,
	}

	return nil, App{AppMux: finEntry.FinanceEntrypoint()}
}
