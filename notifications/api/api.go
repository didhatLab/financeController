package api

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"main/finances/entrypoints/middleware"
	"main/finances/entrypoints/webmodels"
	"net/http"
)

type NotificationAPI struct {
	Redis *redis.Client
}

func (na NotificationAPI) NotificationApi() http.Handler {
	mux := http.NewServeMux()

	return mux
}

func (na NotificationAPI) GetNotificationUserStatus(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	realUser, ok := middleware.UserFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	val, err := na.Redis.Get(ctx, fmt.Sprintf("%d:userNotificationStatus", realUser.UserId)).Result()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusOK, struct{ Status string }{Status: val})

}

func (na NotificationAPI) SetNotificationStatus(w http.ResponseWriter, req *http.Request) {
	var body SetNotificationStatusReq
	ctx := req.Context()

	realUser, ok := middleware.UserFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := webmodels.DecodeJSONBody(w, req, &body)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	na.Redis.Set(ctx, fmt.Sprintf("%d:userNotificationStatus", realUser.UserId), body.Status, 0)

	w.WriteHeader(http.StatusOK)
	return

}
