package utils

import "golang.org/x/crypto/bcrypt"

func Panic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
