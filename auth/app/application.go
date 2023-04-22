package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"main/auth/entrypoints"
	"main/auth/repo"
	"main/auth/services"
	"main/auth/signatory"
	"net/http"
)

type App struct {
	AppMux *http.ServeMux
}

func NewApplication(ctx context.Context, pool *pgxpool.Pool) (App, error) {
	secretKey := "secretKey"

	authRepo := repo.NewAuthRepository(pool)

	signService := signatory.NewSignService([]byte(secretKey))

	authEntry := entrypoints.AuthEntryPoint{
		Ctx:             ctx,
		RegisterService: services.NewRegisterUserService(authRepo),
		AuthService:     services.NewAuthService(authRepo, signService),
	}

	return App{AppMux: authEntry.AuthEntryPoint()}, nil
}
