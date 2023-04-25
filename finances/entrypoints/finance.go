package entrypoints

import (
	"context"
	"main/finances/entrypoints/middleware"
	"main/finances/entrypoints/webmodels"
	"main/finances/models/finance"
	"main/finances/services"
	"net/http"
)

type FinanceEntryPoint struct {
	CreateSpendService services.CreatingSpendService
	GetSpendsService   services.GettingUserSpendsService
	DeleteSpendService services.DeleteSpendsService

	Ctx context.Context
}

func (fe FinanceEntryPoint) FinanceEntrypoint() *http.ServeMux {
	financeMux := http.NewServeMux()

	spending := fe.spendingEntrypoint()

	financeMux.Handle("/", http.StripPrefix("/spending", spending))

	return financeMux
}

func (fe FinanceEntryPoint) spendingEntrypoint() *http.ServeMux {
	spending := http.NewServeMux()

	spending.Handle("/save", middleware.AuthMiddleware(http.HandlerFunc(fe.saveNewSpending)))
	spending.Handle("/get", middleware.AuthMiddleware(http.HandlerFunc(fe.getUserSpends)))
	spending.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(fe.deleteUserSpend)))

	return spending
}

func (fe FinanceEntryPoint) saveNewSpending(w http.ResponseWriter, req *http.Request) {
	var newSpending webmodels.TestSpending
	err := webmodels.DecodeJSONBody(w, req, &newSpending)

	ctx := req.Context()
	realUser, ok := middleware.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}
	err = fe.CreateSpendService.CreateNewSpend(fe.Ctx, realUser,
		finance.SpendingFromUserInput(newSpending, realUser.UserId))

	if err != nil {
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusCreated, struct{}{})

}

func (fe FinanceEntryPoint) getUserSpends(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	realUser, ok := middleware.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, spends := fe.GetSpendsService.GetUserSpends(fe.Ctx, realUser)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusInternalServerError, struct{}{})
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusOK, spends)

}

func (fe FinanceEntryPoint) deleteUserSpend(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	realUser, ok := middleware.UserFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var deleteReq webmodels.DeleteRequest

	err := webmodels.DecodeJSONBody(w, req, &deleteReq)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	err = fe.DeleteSpendService.DeleteUserSpend(ctx, realUser.UserId, deleteReq.SpendId)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
