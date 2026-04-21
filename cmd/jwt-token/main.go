// Command jwt-token prints an RS256 JWT to stdout using flags or environment.
package main

import (
	"crypto/rsa"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	jwtpkg "github.com/dendavidov/go-utils/pkg/jwt"
)

// version is set by GoReleaser at link time.
var version = "dev"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		id       = flag.String("id", "", "subject user id (env JWT_SUBJECT_ID)")
		email    = flag.String("email", "", "email claim (env JWT_EMAIL)")
		access   = flag.Bool("access", false, "set isAccessToken claim")
		ttlStr   = flag.String("ttl", "720h", "token TTL as Go duration (e.g. 24h, 30m)")
		keyB64   = flag.String("key", "", "base64-encoded RSA private key PEM (env JWT_PRIVATE_KEY_B64)")
		keyFile  = flag.String("key-file", "", "path to PEM file (optional; overrides -key if set)")
		printVer = flag.Bool("version", false, "print version and exit")
	)
	flag.Parse()

	if *printVer {
		fmt.Println(version)
		return nil
	}

	idVal := firstNonEmpty(*id, os.Getenv("JWT_SUBJECT_ID"))
	emailVal := firstNonEmpty(*email, os.Getenv("JWT_EMAIL"))
	keyVal := firstNonEmpty(*keyB64, os.Getenv("JWT_PRIVATE_KEY_B64"))

	var priv *rsa.PrivateKey
	if *keyFile != "" {
		b, err := os.ReadFile(*keyFile)
		if err != nil {
			return fmt.Errorf("read key file: %w", err)
		}
		priv, err = jwtpkg.ParsePrivateKeyFromPEM(b)
		if err != nil {
			return err
		}
	}

	if idVal == "" || emailVal == "" {
		return errors.New("id and email are required (-id / JWT_SUBJECT_ID, -email / JWT_EMAIL)")
	}
	if priv == nil && keyVal == "" {
		return errors.New("private key required (-key, JWT_PRIVATE_KEY_B64, or -key-file)")
	}

	ttl, err := time.ParseDuration(*ttlStr)
	if err != nil {
		return fmt.Errorf("parse ttl: %w", err)
	}

	if priv == nil {
		var err error
		priv, err = jwtpkg.ParsePrivateKeyFromBase64(keyVal)
		if err != nil {
			return err
		}
	}

	token, err := jwtpkg.GenerateToken(priv, jwtpkg.Claims{
		ID:       idVal,
		Email:    emailVal,
		IsAccess: *access,
	}, ttl)
	if err != nil {
		return err
	}
	fmt.Print(token)
	return nil
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
