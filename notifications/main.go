package main

import (
	"github.com/redis/go-redis/v9"
	"main/finances/entrypoints/middleware"
	"main/notifications/api"
	"main/notifications/resolver"
	"net/http"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	eventResolver := resolver.EventsResolver{Redis: rdb}

	go func() {
		eventResolver.StartResolve()
	}()

	entrypoint := api.NotificationAPI{Redis: rdb}

	mux := http.NewServeMux()

	mux.Handle("/set", http.HandlerFunc(entrypoint.SetNotificationStatus))
	mux.Handle("/get", http.HandlerFunc(entrypoint.GetNotificationUserStatus))

	http.ListenAndServe(":4005", middleware.AuthMiddleware(mux))

}
