package main

import (
	"github.com/B-Bridger/server/handler"
	"github.com/B-Bridger/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	authRequired := r.Group("/users", middleware.AuthMiddleware())
	{
		authRequired.GET("/", userHandler.GetUser)
		authRequired.PUT("/", userHandler.UpdateUser)
		authRequired.DELETE("/", userHandler.DeleteUser)
		authRequired.POST("/profile-image", userHandler.UploadProfileImage)
	}

	user := r.Group("/users")
	{
		user.POST("/", userHandler.CreateUser)
	}

	r.POST("/login", userHandler.Login)
	r.Static("/static", "./static")

	return r
}
