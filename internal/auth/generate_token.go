package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (m *TokenManager) Generate(userID string, email string) (string, error) {
	now := time.Now()

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(m.privateKey)
}
