package utils

import (
	"crypto/aes"
	"encoding/hex"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"golang.org/x/crypto/bcrypt"
)

type Crypto struct {
	config configuration.Configuration
}

func (_ Crypto) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	hashedPassword := string(bytes)
	return hashedPassword, nil
}

func (_ Crypto) VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u Crypto) EncryptWithAppKey(plaintext string) (string, error) {
	return u.EncryptString(plaintext, u.config.Security.Key)
}

func (_ Crypto) EncryptString(plaintext string, key string) (string, error) {
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	msgByte := make([]byte, len(plaintext))
	c.Encrypt(msgByte, []byte(plaintext))

	return hex.EncodeToString(msgByte), nil
}
