package utils

import (
	"RyftFramework/configuration"
	"crypto/aes"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ---
//
// This function is used to hash the password using bcrypt
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	return string(bytes)
}

// CheckPasswordHash ---
//
// This function is used to check if the password is correct
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EncryptString(plainText string) (string, error) {

	c, err := aes.NewCipher([]byte(configuration.ApplicationConfig.Security.Key))

	if err != nil {
		ErrorLogger.Print(err)
		return "", err
	}

	msgByte := make([]byte, len(plainText))
	c.Encrypt(msgByte, []byte(plainText))
	return hex.EncodeToString(msgByte), nil

}
