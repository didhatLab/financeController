package entrypoints

import (
	"context"
	"main/src/entrypoints/webmodels"
	"main/src/models/user"
	"main/src/services/interfaces"
	"net/http"
)

type FinanceEntryPoint struct {
	CreateSpendService interfaces.CreatingSpendService
	GetSpendsService   interfaces.GettingUserSpendsService

	Ctx context.Context
}

func (fe FinanceEntryPoint) FinanceEntrypoint() *http.ServeMux {
	finance := http.NewServeMux()

	spending := fe.spendingEntrypoint()

	finance.Handle("/", http.StripPrefix("/spending", spending))

	return finance
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
		if err != nil {
			return
		}
		return
	}
	mockUser := user.User{UserId: 1, Username: "dan"}
	err = fe.CreateSpendService.CreateNewSpend(fe.Ctx, mockUser, newSpending.Name)

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
