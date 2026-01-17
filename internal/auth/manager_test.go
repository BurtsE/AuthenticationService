package auth

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestTokenFlow(t *testing.T) {
	logger := logrus.New()
	manager := NewTokenManager(logger)

}
