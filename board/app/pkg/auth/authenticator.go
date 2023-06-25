package auth

import "github.com/google/uuid"

func IsAdminToken(token string) bool {
	return true
}

func IsCustomer(id string) (uuid.UUID, error) {
	return uuid.New(), nil
}
