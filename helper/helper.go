package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwrod string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(passwrod), 14)
	return string(byte), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
