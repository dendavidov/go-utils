package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// Claims carries application fields embedded with standard JWT registered claims.
type Claims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	IsAccess bool   `json:"isAccessToken"`
	jwtv5.RegisteredClaims
}

// ParsePrivateKeyFromPEM parses an RSA private key from a PEM-encoded block.
func ParsePrivateKeyFromPEM(pem []byte) (*rsa.PrivateKey, error) {
	if len(pem) == 0 {
		return nil, errors.New("jwt: empty PEM")
	}
	key, err := jwtv5.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, fmt.Errorf("jwt: parse PEM: %w", err)
	}
	return key, nil
}

// ParsePrivateKeyFromBase64 decodes a base64-encoded PEM block and parses an RSA private key.
func ParsePrivateKeyFromBase64(b64 string) (*rsa.PrivateKey, error) {
	if b64 == "" {
		return nil, errors.New("jwt: empty key material")
	}
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("jwt: decode private key: %w", err)
	}
	return ParsePrivateKeyFromPEM(pemBytes)
}

// GenerateToken signs claims with RS256 using the given private key and TTL from now.
func GenerateToken(priv *rsa.PrivateKey, c Claims, ttl time.Duration) (string, error) {
	if priv == nil {
		return "", errors.New("jwt: private key is nil")
	}
	if ttl <= 0 {
		return "", errors.New("jwt: ttl must be positive")
	}
	now := time.Now()
	exp := now.Add(ttl)
	c.RegisteredClaims = jwtv5.RegisteredClaims{
		ExpiresAt: jwtv5.NewNumericDate(exp),
		IssuedAt:  jwtv5.NewNumericDate(now),
		NotBefore: jwtv5.NewNumericDate(now),
	}
	t := jwtv5.NewWithClaims(jwtv5.SigningMethodRS256, &c)
	return t.SignedString(priv)
}

// ParseToken verifies an RS256 token and returns validated claims.
func ParseToken(tokenString string, pub *rsa.PublicKey) (*Claims, error) {
	if pub == nil {
		return nil, errors.New("jwt: public key is nil")
	}
	claims := &Claims{}
	_, err := jwtv5.ParseWithClaims(tokenString, claims, func(t *jwtv5.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwtv5.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("jwt: unexpected signing method %v", t.Header["alg"])
		}
		return pub, nil
	}, jwtv5.WithValidMethods([]string{jwtv5.SigningMethodRS256.Alg()}))
	if err != nil {
		return nil, err
	}
	return claims, nil
}
