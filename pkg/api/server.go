package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SethukumarJ/go-gin-clean-arch/cmd/api/docs"
	handler "github.com/SethukumarJ/go-gin-clean-arch/pkg/api/handler"
	middleware "github.com/SethukumarJ/go-gin-clean-arch/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	engine.POST("/login", middleware.LoginHandler)

	// Auth middleware
	// api := engine.Group("/api", middleware.AuthorizationMiddleware)

	engine.GET("/api/users", userHandler.FindAll)
	engine.GET("/users", userHandler.FindByID)
	engine.POST("/api/users", userHandler.Save)
	engine.DELETE("users", userHandler.Delete)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
