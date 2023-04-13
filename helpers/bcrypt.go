package helpers

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string) (string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	return string(hash)
}

func ComparePassword(hash, password []byte) bool {
	hash,pass := []byte(hash),[]byte(password)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	
	return err == nil
}