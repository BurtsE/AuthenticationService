package handlers

import (
	"AuthenticationService/internal/auth"
	"AuthenticationService/internal/config"
	internalMiddlware "AuthenticationService/internal/middleware"
	"AuthenticationService/internal/service"
	"AuthenticationService/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	refreshTokenCookieName = "refreshToken"
)

var cookieMaxAge = int(config.GetRefreshTokenTtl().Seconds())

type UserHandler struct {
	service      service.IUserService
	log          *logrus.Logger
	engine       *gin.Engine
	tokenManager *auth.TokenManager
}

func NewUserHandler(
	log *logrus.Logger,
	userService service.IUserService,
	tokenManager *auth.TokenManager,
) *UserHandler {
	handler := &UserHandler{
		log:          log,
		engine:       gin.New(),
		service:      userService,
		tokenManager: tokenManager,
	}

	// handler.engine.SetTrustedProxies([]string{"127.0.0.1"})
	//for later

	// Global middlewares
	handler.engine.Use(
		middleware.LoggerMiddleware(log),
		middleware.PanicHandlerMiddleware(log),
	)

	authRoutes := handler.engine.Group("/api/v1/auth")

	authRoutes.POST("/register", handler.CreateUser)
	authRoutes.POST("/login", handler.AuthorizeUser)
	authRoutes.POST("/refresh-token", handler.RefreshToken)

	userRoutes := handler.engine.Group("/api/v1/users")

	userRoutes.Use(internalMiddlware.JWTAuth(tokenManager))
	userRoutes.DELETE("/:id", handler.DeleteUser)

	return handler
}

func (h *UserHandler) Start() error {
	return h.engine.Run(fmt.Sprintf(":%s", config.GetApplicationPort()))
}

// retrieve engine for tests
func (h *UserHandler) Engine() *gin.Engine {
	return h.engine
}
