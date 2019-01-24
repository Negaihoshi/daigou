package models

import (
	"github.com/negaihoshi/daigou/pkg/util"
	"github.com/rs/xid"
)

type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(data map[string]interface{}) error {
	user := User{
		ID:       xid.New().String(),
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	data["id"] = user.ID

	return nil
}

func CheckUser(username, password string) (bool, error) {
	var user User
	// db.AutoMigrate(&Auth{})
	err := db.Select("id, password").Where(User{Username: username}).First(&user).Error

	if err != nil {
		// if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	match, err := util.ComparePasswordAndHash(password, user.Password)

	if err != nil {
		return false, err
	}

	if match {
		return true, nil
	}

	return false, nil
}

func GetUser(userID string) (User, error) {
	var user User
	// db.AutoMigrate(&Auth{})
	err := db.Select("id, username, email").Where(User{ID: userID}).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
