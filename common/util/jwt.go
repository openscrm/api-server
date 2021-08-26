package util

import (
	"encoding/base64"
	"openscrm/conf"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UID  string `json:"uid"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(uid string, role string, expireAt int64) (string, error) {
	claims := Claims{
		uid,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    conf.Settings.App.Name,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := base64.StdEncoding.DecodeString(conf.Settings.App.Key)
	if err != nil {
		return "", err
	}
	token, err := tokenClaims.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return base64.StdEncoding.DecodeString(conf.Settings.App.Key)
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
