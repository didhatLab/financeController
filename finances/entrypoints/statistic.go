package entrypoints

import (
	"main/finances/entrypoints/middleware"
	"main/finances/entrypoints/webmodels"
	"main/finances/services/statistic"
	"net/http"
)

type StatisticEntryPoint struct {
	StatLoader statistic.StatLoaderService
}

func (se StatisticEntryPoint) LoadStatsForUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	realUser, ok := middleware.UserFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, stat := se.StatLoader.LoadStatsForUser(ctx, realUser.UserId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusOK, stat)
	return

}
