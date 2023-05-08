package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	application "main/finances/app"
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

	postgresUri := os.Getenv("POSTGRES_URI")
	startupPort := os.Getenv("FINANCES_STARTUP_PORT")

	pool, err := pgxpool.New(ctx, postgresUri)
	err, app := application.NewApplication(ctx, pool)

	if err != nil {
		return
	}

	http.ListenAndServe(fmt.Sprintf(":%s", startupPort), cors.AllowAll().Handler(app.AppMux))

}
