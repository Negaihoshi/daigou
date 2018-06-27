package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/negaihoshi/daigou/db"
	"github.com/negaihoshi/daigou/forms"
	"github.com/rs/xid"
	"gopkg.in/hlandau/passlib.v1"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	db := db.GetDB()
	id := xid.New()
	user := User{
		ID:        id.String(),
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}

	hash, err := passlib.Hash("password")
	if err != nil {
		// couldn't hash password for some reason
		return
	}
	fmt.Println(hash)

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("TableUsers"),
	}
	if _, err := db.PutItem(params); err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	return &user, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(id),
			},
		},
		TableName:      aws.String("TableUsers"),
		ConsistentRead: aws.Bool(true),
	}
	resp, err := db.GetItem(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user *User
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
