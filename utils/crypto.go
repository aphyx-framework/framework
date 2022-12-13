package utils

import (
	"RyftFramework/bootstrapper/logging"
	"RyftFramework/configuration"
	"RyftFramework/di"
	"crypto/aes"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ---
//
// This function is used to hash the password using bcrypt
func HashPassword(password string) string {
	logger := di.Dependency.Get(di.Logger).(logging.ApplicationLogger)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
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
	config := di.Dependency.Get(di.Config).(configuration.Configuration)
	logger := di.Dependency.Get(di.Logger).(logging.ApplicationLogger)

	c, err := aes.NewCipher([]byte(config.Security.Key))

	if err != nil {
		logger.ErrorLogger.Print(err)
		return "", err
	}

	msgByte := make([]byte, len(plainText))
	c.Encrypt(msgByte, []byte(plainText))
	return hex.EncodeToString(msgByte), nil

}
