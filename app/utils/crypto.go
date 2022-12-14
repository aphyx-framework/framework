package utils

import (
	"crypto/aes"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func (_ Util) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	hashedPassword := string(bytes)
	return hashedPassword, nil
}

func (_ Util) VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u Util) EncryptWithAppKey(plaintext string) (string, error) {
	return u.EncryptString(plaintext, u.Config.Security.Key)
}

func (_ Util) EncryptString(plaintext string, key string) (string, error) {
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	msgByte := make([]byte, len(plaintext))
	c.Encrypt(msgByte, []byte(plaintext))

	return hex.EncodeToString(msgByte), nil
}
