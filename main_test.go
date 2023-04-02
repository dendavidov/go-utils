package main

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	privateKeyInterface, _ := getPrivateKey(jwtSecret)
	privateKey, ok := privateKeyInterface.(*rsa.PrivateKey)

	if !ok {
		t.Fatal("privateKey is not of type *rsa.PrivateKey")
	}

	token, err := getToken(privateKey, false)

	if err != nil {
		t.Errorf("getToken returned an error: %v", err)
	}

	if token == "" {
		t.Error("getToken returned an empty token")
	}

	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey.Public(), nil
	})

	if err != nil {
		t.Errorf("jwt.ParseWithClaims returned an error: %v", err)
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		t.Error("parsedToken.Claims is not of type *Claims")
	}

	if claims.ID != userID {
		t.Errorf("Expected ID to be '%s', but got '%s'", userID, claims.ID)
	}

	if claims.Email != userEmail {
		t.Errorf("Expected Email to be '%s', but got '%s'", userEmail, claims.Email)
	}

	if claims.IsAccess != false {
		t.Error("Expected IsAccess to be false, but got true")
	}

	expirationTime := time.Unix(claims.ExpiresAt, 0)
	expectedExpiration := time.Now().Add(30 * 24 * time.Hour)
	if expirationTime.Sub(expectedExpiration) > time.Minute {
		t.Errorf("Token expiration time is not within expected range")
	}
}
