package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

func LoadPrivateKeyFromEnv() (*rsa.PrivateKey, error) {
	b64 := os.Getenv("JWT_PRIVATE_SECRET")
	if b64 == "" {
		return nil, fmt.Errorf("JWT_PRIVATE_SECRET not set")
	}
	der, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(der)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func LoadPublicKeyFromEnv() (*rsa.PublicKey, error) {
	b64 := os.Getenv("JWT_PUBLIC_KEY")
	if b64 == "" {
		return nil, fmt.Errorf("JWT_PUBLIC_KEY not set")
	}
	der, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(der)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func GetTokenAccessTTL() time.Duration {
	ttlString := getEnv("TOKEN_ACCESS_TTL", "")
	if ttlString == "" {
		return 0
	}

	ttl, err := time.ParseDuration(ttlString)
	if err != nil {
		return 0
	}

	return ttl
}

func GetTokenIssuer() string {
	return getEnv("TOKEN_ISSUER", "")
}
