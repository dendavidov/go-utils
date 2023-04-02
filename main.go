package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	IsAccess bool   `json:"isAccessToken"`
	jwt.StandardClaims
}

func getToken(privateKey interface{}, isAccessToken bool) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{

		ID:       userID,
		Email:    userEmail,
		IsAccess: isAccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(privateKey)
}

func getPrivateKey(jwtSecret string) (interface{}, error) {

	jwtPrivateKey, err := base64.StdEncoding.DecodeString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(jwtPrivateKey)
}

func main() {
	privateKey, _ := getPrivateKey(jwtSecret)

	token, _ := getToken(privateKey, false)

	fmt.Printf(token)
}
