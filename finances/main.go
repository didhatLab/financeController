package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
	application "main/finances/app"
	"net/http"
)

func main() {

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgresql://postgres:postgres@localhost:5432/postgres")
	err, app := application.NewApplication(ctx, pool)

	if err != nil {
		return
	}

	//entry := entrypoints.FinanceEntrypoint()

	http.ListenAndServe(":4000", cors.AllowAll().Handler(app.AppMux))

}
