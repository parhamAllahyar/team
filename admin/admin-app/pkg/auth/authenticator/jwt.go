package authenticator

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TODO: get  secret key from config map
type JWT struct {
	Id uuid.UUID
}

func DecodeToken(userToken string) (JWT, error) {

	// fmt.Println("send", userToken)
	userToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjVkMDQ3YWZkLTM5ZTEtNDQwZS05ZjZjLTliMjE2YWQxZTAyYSIsIm5iZiI6IjIwMjQtMDYtMjNUMDk6NTI6MjMuMDQ3NjY3OTM3WiJ9.hO38GRiO5LHnC3OYdQqAeE1D_SytYPm3MMGZwjr793Y"

	// fmt.Println("fixed", userToken)

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		//TODO : read key from k8s secret
		return []byte("your key"), nil
	})

	fmt.Println("err is", err)

	if err == nil && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return JWT{
			Id: uuid.MustParse(claims["id"].(string)),
		}, nil
	} else {
		return JWT{}, err
	}
}

// Generate Token
func ExtractClaims(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"nbf": time.Now().AddDate(1, 0, 0),

		// time.Now().Local().Add(time.Hour * time.Duration(Hours) +
		// time.Minute * time.Duration(Mins) +
		// time.Second * time.Duration(Sec))

	})
	//TODO : read key from k8s secret
	tokenString, err := token.SignedString([]byte("your key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Get token from Auth header
func GetTokenString(token string) string {
	const BEARER_SCHEMA = "Bearer "
	return token[len(BEARER_SCHEMA):]
}
