package signatory

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type SignService struct {
	secretKey []byte
}

func NewSignService(secretKey []byte) SignService {
	return SignService{secretKey: secretKey}
}

func (ss SignService) SignToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": map[string]int{"id": userId},
	})

	tokenString, err := token.SignedString([]byte(ss.secretKey))

	if err != nil {
		log.Printf("Can not sign token!")
		return "", err
	}

	return tokenString, nil
}

func (ss SignService) ValidateToken(token string) (jwt.MapClaims, bool) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return ss.secretKey, nil
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
