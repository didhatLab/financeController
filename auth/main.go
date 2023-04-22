package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/auth/app"
	"net/http"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgresql://postgres:postgres@localhost:5432/postgres")

	application, err := app.NewApplication(ctx, pool)

	err = http.ListenAndServe(":4001", application.AppMux)
	if err != nil {
		log.Fatal(err)
		return
	}

}
