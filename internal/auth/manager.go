package auth

import (
	"AuthenticationService/internal/config"
	"crypto/rsa"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

const defaultTokenTTL = time.Hour * 24

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type TokenManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	ttl        time.Duration
	issuer     string
	log        *logrus.Logger
}

func NewTokenManager(
	logger *logrus.Logger,
	privateKey *rsa.PrivateKey,
	publicKey *rsa.PublicKey,
) *TokenManager {
	ttl := config.GetTokenAccessTTL()
	issuer := config.GetTokenIssuer()

	if privateKey == nil || publicKey == nil {
		logger.Fatal("TokenManager keys missing")
	}

	if ttl == 0 {
		ttl = defaultTokenTTL
		logger.Warnf("TokenManager ttl is 0, set to default: %v", ttl)
	}

	if issuer == "" {
		logger.Fatal("TokenManager issuer is empty")
	}

	manager := &TokenManager{
		privateKey: privateKey,
		publicKey:  publicKey,
		ttl:        ttl,
		issuer:     issuer,
		log:        logger,
	}

	return manager
}
