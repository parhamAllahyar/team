package authenticator

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TODO: get  secret key from config map
type JWT struct {
	Id uuid.UUID
}

func DecodeToken(userToken string) (JWT, error) {
	log.Println(userToken)
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		//TODO : read key from k8s secret
		return []byte("your key"), nil
	})

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
