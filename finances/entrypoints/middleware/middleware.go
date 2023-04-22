package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"main/finances/models/user"
	"net/http"
)

const secretKey = "secretKey"

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		userId, err := extractUserId(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := req.Context()
		ctx = ContextWithUser(ctx, user.NewUser(userId))
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

func extractUserId(req *http.Request) (int, error) {
	token := req.Header.Get("Auth-Token")

	if token == "" {
		return -1, errors.New("not found token in Header")
	}

	claims, ok := validateToken(token)

	if !ok {
		return -1, errors.New("invalid token")
	}

	//data := claims["id"](int)

	data := claims["data"].(map[string]interface{})

	userId := data["id"].(float64)

	return int(userId), nil

}

func validateToken(token string) (jwt.MapClaims, bool) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
