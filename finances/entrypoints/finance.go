package entrypoints

import (
	"context"
	"main/finances/entrypoints/webmodels"
	"main/finances/models/finance"
	"main/finances/models/user"
	"main/finances/services"
	"net/http"
)

type FinanceEntryPoint struct {
	CreateSpendService services.CreatingSpendService
	GetSpendsService   services.GettingUserSpendsService

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
	spending.HandleFunc("/save", fe.saveNewSpending)
	spending.HandleFunc("/get", fe.getUserSpends)

	return spending
}

func (fe FinanceEntryPoint) saveNewSpending(w http.ResponseWriter, req *http.Request) {
	var newSpending webmodels.TestSpending
	err := webmodels.DecodeJSONBody(w, req, &newSpending)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}
	mockUser := user.User{UserId: 1, Username: "dan"}
	err = fe.CreateSpendService.CreateNewSpend(fe.Ctx, mockUser,
		finance.SpendingFromUserInput(newSpending, mockUser.UserId))

	if err != nil {
		return
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusCreated, struct{}{})

}

func (fe FinanceEntryPoint) getUserSpends(w http.ResponseWriter, req *http.Request) {
	mockUser := user.User{UserId: 1, Username: "dan"}

	err, spends := fe.GetSpendsService.GetUserSpends(fe.Ctx, mockUser)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusInternalServerError, struct{}{})
	}

	webmodels.EncodeJSONResponseBody(w, http.StatusOK, spends)

}
