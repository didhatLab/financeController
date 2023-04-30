package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"main/finances/entrypoints"
	"main/finances/repo"
	"main/finances/services"
	"main/finances/services/privacy"
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
		CreateSpendService:    services.NewCreateSpendsService(finRepo, groupAccessChecker),
		GetSpendsService:      services.NewGetSpendsService(finRepo),
		DeleteSpendService:    services.NewDeleteSpendsService(finRepo, groupAccessChecker),
		UpdateSpendService:    services.NewUpdateSpendsService(finRepo, groupAccessChecker),
		GetGroupSpendsService: services.NewGroupSpendsService(finRepo, groupAccessChecker),
		Ctx:                   ctx,
	}

	return nil, App{AppMux: finEntry.FinanceEntrypoint()}
}
