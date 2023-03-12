package app

import (
	"context"
	"main/src/adapter/repository"
	"main/src/entrypoints"
	"main/src/services"
	"net/http"
)

type App struct {
	AppMux *http.ServeMux
}

func NewApplication(ctx context.Context) (error, App) {

	finRepo := repository.NewMemoryFinanceRepository()

	finEntry := entrypoints.FinanceEntryPoint{
		CreateSpendService: services.CreateCreateSpendsService(finRepo),
		GetSpendsService:   services.CreateGetSpendsService(finRepo),
		Ctx:                ctx,
	}

	return nil, App{AppMux: finEntry.FinanceEntrypoint()}
}
