package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"main/finances/entrypoints"
	"main/finances/repo"
	"main/finances/services"
	"net/http"
)

type App struct {
	AppMux *http.ServeMux
}

func NewApplication(ctx context.Context, pool *pgxpool.Pool) (error, App) {

	finRepo := repo.NewPostgresFinanceRepository(pool)

	finEntry := entrypoints.FinanceEntryPoint{
		CreateSpendService: services.CreateCreateSpendsService(finRepo),
		GetSpendsService:   services.CreateGetSpendsService(finRepo),
		Ctx:                ctx,
	}

	return nil, App{AppMux: finEntry.FinanceEntrypoint()}
}
