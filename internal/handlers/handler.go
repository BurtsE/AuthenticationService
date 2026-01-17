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

type UserHandler struct {
	service service.IUserService
	log     *logrus.Logger
	engine  *gin.Engine
}

func NewUserHandler(
	log *logrus.Logger,
	userService service.IUserService,
	tokenManager *auth.TokenManager,
) *UserHandler {
	handler := &UserHandler{
		log:     log,
		engine:  gin.New(),
		service: userService,
	}

	handler.engine.Use(
		middleware.LoggerMiddleware(log),
		middleware.PanicHandlerMiddleware(log),
	)
	api := handler.engine.Group("/api/v1")

	api.POST("/register", handler.CreateUser)

	authorized := api.Group("/")
	authorized.Use(internalMiddlware.JWTAuth(tokenManager))
	authorized.DELETE("/", handler.DeleteUser)

	return handler
}

func (h *UserHandler) Start() error {
	return h.engine.Run(fmt.Sprintf(":%s", config.GetApplicationPort()))
}
