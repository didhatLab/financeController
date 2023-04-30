package entrypoints

import (
	"context"
	"main/finances/entrypoints/middleware"
	"main/finances/entrypoints/webmodels"
	"main/finances/models/finance"
	"main/finances/services/spend"
	"net/http"
)

type FinanceEntryPoint struct {
	CreateSpendService    spend.CreatingSpendService
	GetSpendsService      spend.GettingUserSpendsService
	DeleteSpendService    spend.DeleteSpendsService
	UpdateSpendService    spend.UpdateSpendsService
	GetGroupSpendsService spend.GetGroupSpendsService

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
	spending.Handle("/update", middleware.AuthMiddleware(http.HandlerFunc(fe.editUserSpend)))
	spending.Handle("/group/get", middleware.AuthMiddleware(http.HandlerFunc(fe.getGroupSpends)))

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
	err, newSpendId := fe.CreateSpendService.CreateNewSpend(fe.Ctx, realUser,
		finance.SpendingFromUserInput(newSpending, realUser.UserId))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusCreated, struct {
		SpendId int `json:"spend_id"`
	}{SpendId: newSpendId})

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

	err = fe.DeleteSpendService.DeleteUserSpend(ctx, realUser.UserId, deleteReq.SpendId, deleteReq.GroupId)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (fe FinanceEntryPoint) editUserSpend(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	realUser, ok := middleware.UserFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var updateReq webmodels.UpdateRequest

	err := webmodels.DecodeJSONBody(w, req, &updateReq)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	err = fe.UpdateSpendService.UpdateUserSpend(ctx, updateReq, realUser.UserId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (fe FinanceEntryPoint) getGroupSpends(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	realUser, ok := middleware.UserFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var getReq webmodels.GroupSpendRequest

	err := webmodels.DecodeJSONBody(w, req, &getReq)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	spends, err := fe.GetGroupSpendsService.GetGroupSpends(ctx, getReq.GroupId, realUser.UserId)
	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusOK, spends)

}
