package models

import (
	"fmt"

	"github.com/rs/xid"
)

type ThirdPartyProvider struct {
	// Model

	UserID string
	User   User

	ID                string `gorm:"primary_key" json:"id"`
	Provider          string `json:"provider"`
	Email             string `json:"email"`
	ProviderID        string `json:"provider_id"`
	Avatar            string `json:"avatar"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func CreateProvider(data map[string]interface{}) error {
	fmt.Sprintf(`%v`, data["email"])
	provider := ThirdPartyProvider{
		ID:                xid.New().String(),
		Provider:          data["provider"].(string),
		Email:             data["email"].(string),
		ProviderID:        data["provider_id"].(string),
		Avatar:            data["avatar"].(string),
		AccessToken:       data["access_token"].(string),
		AccessTokenSecret: data["access_token_secret"].(string),

		UserID: data["user_id"].(string),
	}
	if err := db.Create(&provider).Error; err != nil {
		return err
	}

	return nil
}

func CheckProvider(provider, providerID string) (bool, error) {
	var providerModel ThirdPartyProvider
	result := db.Select("id").Where(ThirdPartyProvider{Provider: provider, ProviderID: providerID}).First(&providerModel)

	// fmt.Printf("FATAL: %+v\n", errors.Wrap(result.Error, "error"))
	// if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
