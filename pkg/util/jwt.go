package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/negaihoshi/daigou/pkg/setting"
)

var jwtSecret = []byte(setting.AppConfig.App.JwtSecret)

type Claims struct {
	// Username string `json:"username"`
	// Password string `json:"password"`

	// Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	// func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		// EncodeMD5(username),
		// EncodeMD5(password),
		// Username: user.Username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
