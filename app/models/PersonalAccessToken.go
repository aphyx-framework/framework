package models

import (
	"github.com/rama-adi/RyFT-Framework/app"
	utils2 "github.com/rama-adi/RyFT-Framework/app/utils"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"go.uber.org/fx"
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

func (pat PersonalAccessToken) Logout() error {
	return app.DB.Delete(&pat).Error
}

func (_ PersonalAccessToken) RevokeToken(token string) error {

	var config *configuration.Configuration
	fx.Populate(&config)

	tokenEnc, err := utils2.EncryptString(token, config.Security.Key)

	if err != nil {
		return err
	}

	return app.DB.Delete(&PersonalAccessToken{}, "token = ?", tokenEnc).Error
}

func (_ PersonalAccessToken) CreateTokenForUser(user User, name string, permanent bool) (PersonalAccessTokenResponse, error) {
	plaintextToken := utils2.RandStringRunes(40)

	var expiry time.Time

	if permanent {
		// If the token is permanent, set the expiry to 100 years from now
		expiry = time.Now().AddDate(100, 0, 0)
	} else {
		// If the token is not permanent, set the expiry to 1 month from now
		expiry = time.Now().AddDate(0, 1, 0)
	}

	var config *configuration.Configuration
	fx.Populate(&config)

	tokenEnc, err := utils2.EncryptString(plaintextToken, config.Security.Key)

	if err != nil {
		return PersonalAccessTokenResponse{}, err
	}

	token := PersonalAccessToken{
		UserID:    user.ID,
		Name:      name,
		Token:     tokenEnc,
		ExpiresAt: expiry,
	}

	err = app.DB.Create(&token).Error

	return PersonalAccessTokenResponse{
		ID:        token.ID,
		UserID:    user.ID,
		Name:      name,
		Token:     plaintextToken,
		ExpiresAt: expiry,
	}, err
}
