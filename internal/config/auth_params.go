package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

func ReadPrivateKey() (*rsa.PrivateKey, error) {
	filePath := getEnv("JWT_PRIVATE_KEY", "./keys/private_key.pem")
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	if _, ok := block.Headers["Proc-Type"]; ok {
		return nil, fmt.Errorf("encrypted private keys are not supported without passphrase handling")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		if rsaKey, ok := key.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, fmt.Errorf("not an RSA private key")
	default:
		return nil, fmt.Errorf("unsupported PEM type: %s", block.Type)
	}
}
func ReadPublicKey() (*rsa.PublicKey, error) {
	filePath := getEnv("JWT_PUBLIC_KEY", "./keys/private_key.pem.pub")
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("found a non-RSA public key")
	}

	return rsaPub, nil
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
	return getEnv("TOKEN_ISSUER", "test_issuer")
}
