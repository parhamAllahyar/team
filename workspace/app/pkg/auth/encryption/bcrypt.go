package encryption

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hasheValue, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasheValue), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
