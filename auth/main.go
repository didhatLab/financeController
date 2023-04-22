package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/auth/app"
	"net/http"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgresql://postgres:postgres@localhost:5432/postgres")

	//if err != nil {
	//	return
	//}
	//
	//err, app := application.NewApplication(ctx, pool)
	//
	//http.ListenAndServe(":4001", app.AppMux)
	//payload := NewPa
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{"user_id": "sss"})

	application, err := app.NewApplication(ctx, pool)

	err = http.ListenAndServe(":4001", application.AppMux)
	if err != nil {
		log.Fatal(err)
		return
	}

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"id": 126,
	//})
	//
	//tokenString, err := token.SignedString([]byte("key"))
	//fmt.Println(tokenString, err)
	//
	//parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("there's an error with the signing method")
	//	}
	//	return []byte("key"), nil
	//})
	//
	//if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
	//	fmt.Println(claims)
	//} else {
	//	log.Printf("Invalid JWT Token")
	//	fmt.Println("nope")
	//}

}
