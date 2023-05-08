package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"main/auth/app"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_URI"))

	startPort := os.Getenv("AUTH_STARTUP_PORT")

	application, err := app.NewApplication(ctx, pool)

	err = http.ListenAndServe(fmt.Sprintf(":%s", startPort), cors.AllowAll().Handler(application.AppMux))
	if err != nil {
		log.Fatal(err)
		return
	}

}
