package authenticator

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TODO: get  secret key from config map
type JWT struct {
	///	secret string
}

func DecodeToken(userToken string) (jwt.MapClaims, error) {
	log.Println(userToken)
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {

		//TODO : read key from k8s secret
		return []byte("your key"), nil
	})
	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")
		claims, _ := token.Claims.(jwt.MapClaims)
		return claims, nil
	} else {
		return nil, err
	}
}

// Generate Token
func ExtractClaims(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	//TODO : read key from k8s secret
	tokenString, err := token.SignedString([]byte("your key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
