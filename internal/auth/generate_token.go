package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (m *TokenManager) GenerateTokenPair(userID string, email string) (string, string, error) {
	now := time.Now()

	accessTokenClaims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.accessTokenTtl)),
		},
	}

	refreshTokenClaims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.refreshTokenTtl)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString(m.privateKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString(m.privateKey)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}
