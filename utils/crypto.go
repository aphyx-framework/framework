package utils

import "golang.org/x/crypto/bcrypt"

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
