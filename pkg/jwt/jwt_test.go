package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"
)

func TestGenerateAndParseToken(t *testing.T) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	const id = "user-1"
	const email = "a@example.com"
	c := Claims{
		ID:       id,
		Email:    email,
		IsAccess: false,
	}
	ttl := 30 * 24 * time.Hour
	token, err := GenerateToken(priv, c, ttl)
	if err != nil {
		t.Fatalf("GenerateToken: %v", err)
	}
	if token == "" {
		t.Fatal("empty token")
	}
	got, err := ParseToken(token, &priv.PublicKey)
	if err != nil {
		t.Fatalf("ParseToken: %v", err)
	}
	if got.ID != id {
		t.Errorf("ID: want %q got %q", id, got.ID)
	}
	if got.Email != email {
		t.Errorf("Email: want %q got %q", email, got.Email)
	}
	if got.IsAccess {
		t.Error("IsAccess: want false")
	}
	if got.ExpiresAt == nil {
		t.Fatal("ExpiresAt nil")
	}
	exp := got.ExpiresAt.Time
	wantExp := time.Now().Add(ttl)
	if exp.Sub(wantExp).Abs() > time.Minute {
		t.Errorf("expiration drift: got %v want ~%v", exp, wantExp)
	}
}

func TestParseTokenWrongKey(t *testing.T) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	other, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	token, err := GenerateToken(priv, Claims{ID: "x", Email: "y"}, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ParseToken(token, &other.PublicKey)
	if err == nil {
		t.Fatal("expected verification error with wrong public key")
	}
}

func TestGenerateTokenNilKey(t *testing.T) {
	_, err := GenerateToken(nil, Claims{}, time.Hour)
	if err == nil {
		t.Fatal("expected error for nil key")
	}
}

func TestGenerateTokenNonPositiveTTL(t *testing.T) {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	_, err := GenerateToken(priv, Claims{}, 0)
	if err == nil {
		t.Fatal("expected error for zero ttl")
	}
}

func TestParsePrivateKeyFromBase64Invalid(t *testing.T) {
	_, err := ParsePrivateKeyFromBase64("not-valid-base64!!!")
	if err == nil {
		t.Fatal("expected decode error")
	}
	_, err = ParsePrivateKeyFromBase64("")
	if err == nil {
		t.Fatal("expected error for empty string")
	}
}
