package entrypoints

import (
	"context"
	"main/auth/entrypoints/webmodels"
	"main/auth/services"
	decoder "main/src/entrypoints/webmodels"
	"net/http"
)

type AuthEntryPoint struct {
	RegisterService services.RegisterUserService
	AuthService     services.AuthService

	Ctx context.Context
}

func (ap AuthEntryPoint) AuthEntryPoint() *http.ServeMux {
	authMux := http.NewServeMux()

	authMux.HandleFunc("/register", ap.registerNewUser)
	authMux.HandleFunc("/token", ap.getTokenForUser)

	return authMux

}

func (ap AuthEntryPoint) registerNewUser(w http.ResponseWriter, req *http.Request) {
	var newUser webmodels.UserModel
	err := decoder.DecodeJSONBody(w, req, &newUser)

	if err != nil {
		decoder.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	err = ap.RegisterService.RegisterNewUser(ap.Ctx, newUser.Username, newUser.Password)

	if err != nil {
		decoder.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	decoder.EncodeJSONResponseBody(w, http.StatusCreated, struct{}{})

}

func (ap AuthEntryPoint) getTokenForUser(w http.ResponseWriter, req *http.Request) {
	var userForAuth webmodels.UserModel
	err := decoder.DecodeJSONBody(w, req, &userForAuth)

	if err != nil {
		decoder.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	token, err := ap.AuthService.AuthUser(ap.Ctx, userForAuth.Username, userForAuth.Password)

	if err != nil {
		decoder.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{}{})
		return
	}

	decoder.EncodeJSONResponseBody(w, http.StatusOK, struct{ Token string }{Token: token})

}
