package utils

import (
	"crypto/aes"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ---
//
// This function is used to hash the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	hashedPassword := string(bytes)

	return hashedPassword, nil
}

// CheckPasswordHash ---
//
// This function is used to check if the password is correct
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EncryptString(plainText string, key string) (string, error) {
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	msgByte := make([]byte, len(plainText))
	c.Encrypt(msgByte, []byte(plainText))

	return hex.EncodeToString(msgByte), nil

}
