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

type PersonalAccessTokenResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (_ PersonalAccessToken) CreateTokenForUser(user User, name string, permanent bool) (PersonalAccessTokenResponse, error) {
	plaintextToken := utils.RandStringRunes(40)

	var expiry time.Time

	if permanent {
		// If the token is permanent, set the expiry to 100 years from now
		expiry = time.Now().AddDate(100, 0, 0)
	} else {
		// If the token is not permanent, set the expiry to 1 month from now
		expiry = time.Now().AddDate(0, 1, 0)
	}

	tokenEnc, err := utils.EncryptString(plaintextToken)

	if err != nil {
		return PersonalAccessTokenResponse{}, err
	}

	token := PersonalAccessToken{
		UserID:    user.ID,
		Name:      name,
		Token:     tokenEnc,
		ExpiresAt: expiry,
	}

	err = database.DB.Create(&token).Error

	return PersonalAccessTokenResponse{
		ID:        token.ID,
		UserID:    user.ID,
		Name:      name,
		Token:     plaintextToken,
		ExpiresAt: expiry,
	}, err
}
