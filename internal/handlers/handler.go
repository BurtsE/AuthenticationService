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

	middlewares := []gin.HandlerFunc{
		middleware.LoggerMiddleware(log),
		middleware.PanicHandlerMiddleware(log),
		internalMiddlware.JWTAuth(tokenManager),
	}

	handler.engine.Use(middlewares...)
	handler.registerRoutes()

	return handler
}

func (h *UserHandler) registerRoutes() {
	gr := h.engine.Group("/users")
	gr.POST("/register", h.CreateUser)
	gr.DELETE("/", h.DeleteUser)
}

func (h *UserHandler) Start() error {
	return h.engine.Run(fmt.Sprintf(":%s", config.GetApplicationPort()))
}
