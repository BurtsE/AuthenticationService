package handlers

import (
	"AuthenticationService/internal/config"
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

func NewUserHandler(log *logrus.Logger, userService service.IUserService) *UserHandler {
	return &UserHandler{
		log:     log,
		engine:  gin.New(),
		service: userService,
	}
}

func (u *UserHandler) registerMiddleware() {
	u.engine.Use(middleware.LoggerMiddleware(u.log))
	u.engine.Use(middleware.PanicHandlerMiddleware(u.log))
}
func (u *UserHandler) registerRoutes() {
	gr := u.engine.Group("/users")
	gr.POST("/register", u.CreateUser)
	gr.DELETE("/", u.DeleteUser)
}

func (u *UserHandler) Start() error {
	u.registerMiddleware()
	u.registerRoutes()
	return u.engine.Run(fmt.Sprintf(":%s", config.GetApplicationPort()))
}
