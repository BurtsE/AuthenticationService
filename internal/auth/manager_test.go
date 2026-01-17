package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestTokenFlow tests generation and validation of tokens
func TestTokenFlow(t *testing.T) {
	t.Setenv("TOKEN_ISSUER", "me")
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)
	publicKey := &privateKey.PublicKey

	logger := logrus.New()
	manager := NewTokenManager(logger, privateKey, publicKey)

	accessToken, refreshToken, err := manager.GenerateTokenPair(uuid.New().String(), "myemail@gmail.com")
	assert.NoError(t, err)

	claims, err := manager.Validate(accessToken)
	assert.NoError(t, err)

	claims, err = manager.Validate(refreshToken)
	assert.NoError(t, err)
	
	t.Log(claims)
}
