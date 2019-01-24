package user_service

import (
	"log"

	"github.com/negaihoshi/daigou/models"
	"github.com/negaihoshi/daigou/pkg/util"
)

type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// type User struct {
// 	*models.User
// }

type ThirdPartyProvider struct {
	ID                string
	Provider          string
	ProviderID        string
	Email             string
	Avatar            string
	AccessToken       string
	AccessTokenSecret string

	UserID string
	User   models.User
}

func (data *User) Add() error {
	encodedHash, err := util.GenerateFromPassword(data.Password)
	if err != nil {
		log.Fatal(err)
	}

	user := map[string]interface{}{
		"username": data.Username,
		"password": encodedHash,
		"email":    data.Email,
	}

	if err := models.CreateUser(user); err != nil {
		return err
	}

	data.ID = user["id"].(string)

	return nil
}

func (data *ThirdPartyProvider) AddProvider() error {
	providerData := map[string]interface{}{
		"provider":            data.Provider,
		"provider_id":         data.ProviderID,
		"email":               data.Email,
		"avatar":              data.Avatar,
		"access_token":        data.AccessToken,
		"access_token_secret": data.AccessTokenSecret,

		"user_id": data.UserID,
	}

	check, _ := models.CheckProvider(data.Provider, data.ProviderID)

	if !check {
		if err := models.CreateProvider(providerData); err != nil {
			return err
		}
	}

	return nil
}

func (user *User) Check() (bool, error) {
	return models.CheckUser(user.Username, user.Password)
}

func (user *User) Get() (models.User, error) {
	return models.GetUser(user.ID)
}
