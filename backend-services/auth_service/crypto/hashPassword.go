package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func Crypt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "none"
	}
	return string(hash)
	// database.DB.Exec("INSERT INTO users(email, password_hash) VALUES (?, ?)", "user@mail.com", hash)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
