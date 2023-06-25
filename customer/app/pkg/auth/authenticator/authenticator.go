package authenticator

import "github.com/google/uuid"

func IsAdminToken(token string) bool {
	return true
}

func IsCustomer() (uuid.UUID, error) {
	return uuid.New(), nil
}
