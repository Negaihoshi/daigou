package models

import (
	"github.com/rs/xid"
)

type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(data map[string]interface{}) error {
	user := User{
		ID:       xid.New().String(),
		Username: data["username"].(string),
		Password: data["password"].(string),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	data["id"] = user.ID

	return nil
}
