package models

import (
	"RyftFramework/database"
	"RyftFramework/utils"
	"gorm.io/gorm"
	"time"
)

type PersonalAccessToken struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;not null"`
	Name      string    `gorm:"not null" json:"name"`
	Token     string    `gorm:"not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}

func (_ PersonalAccessToken) CreateTokenForUser(user User, name string, permanent bool) (PersonalAccessToken, error) {
	plaintextToken := utils.RandStringRunes(40)

	var expiry time.Time

	if permanent {
		// If the token is permanent, set the expiry to 100 years from now
		expiry = time.Now().AddDate(100, 0, 0)
	} else {
		// If the token is not permanent, set the expiry to 1 month from now
		expiry = time.Now().AddDate(0, 1, 0)
	}

	token := PersonalAccessToken{
		UserID:    user.ID,
		Name:      name,
		Token:     utils.HashPassword(plaintextToken),
		ExpiresAt: expiry,
	}

	err := database.DB.Create(&token).Error

	return PersonalAccessToken{
		UserID:    user.ID,
		Name:      name,
		Token:     plaintextToken,
		ExpiresAt: expiry,
	}, err
}
