package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"main/finances/entrypoints/middleware"
	"main/notifications/api"
	"main/notifications/resolver"
	"net/http"
	"os"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	redisdb, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Print(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisdb,
	})

	startPort := os.Getenv("NOTIFICATION_STARTUP_PORT")

	eventResolver := resolver.EventsResolver{Redis: rdb}

	go func() {
		eventResolver.StartResolve()
	}()

	entrypoint := api.NotificationAPI{Redis: rdb}

	mux := http.NewServeMux()

	mux.Handle("/set", http.HandlerFunc(entrypoint.SetNotificationStatus))
	mux.Handle("/get", http.HandlerFunc(entrypoint.GetNotificationUserStatus))

	http.ListenAndServe(fmt.Sprintf(":%s", startPort), middleware.AuthMiddleware(mux))

}
