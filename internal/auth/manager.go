package auth

import (
	"AuthenticationService/internal/config"
	"crypto/rsa"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	defaultRefreshTokenTTL = time.Hour * 24 * 3
	defaultAccessTokenTtl  = time.Minute * 15
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type TokenManager struct {
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
	accessTokenTtl  time.Duration
	refreshTokenTtl time.Duration
	issuer          string
	log             *logrus.Logger
}

func NewTokenManager(
	logger *logrus.Logger,
	privateKey *rsa.PrivateKey,
	publicKey *rsa.PublicKey,
) *TokenManager {
	accessTokenTtl := config.GetAccessTokenTtl()
	refreshTokenTtl := config.GetRefreshTokenTtl()
	issuer := config.GetTokenIssuer()

	if privateKey == nil || publicKey == nil {
		logger.Fatal("TokenManager keys missing")
	}

	if accessTokenTtl == 0 {
		accessTokenTtl = defaultAccessTokenTtl
		logger.Warnf("TokenManager access token ttl is 0, set to default: %v", accessTokenTtl)
	}

	if refreshTokenTtl == 0 {
		refreshTokenTtl = defaultRefreshTokenTTL
		logger.Warnf("TokenManager refresh tokne ttl is 0, set to default: %v", refreshTokenTtl)
	}

	if issuer == "" {
		logger.Fatal("TokenManager issuer is empty")
	}

	manager := &TokenManager{
		privateKey:      privateKey,
		publicKey:       publicKey,
		accessTokenTtl:  accessTokenTtl,
		refreshTokenTtl: refreshTokenTtl,
		issuer:          issuer,
		log:             logger,
	}

	return manager
}
