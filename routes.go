package main

import (
	"github.com/B-Bridger/server/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	user := r.Group("/users")
	{
		user.GET("/:id", userHandler.GetUser)
		user.POST("/", userHandler.CreateUser)
		user.PUT("/:id", userHandler.UpdateUser)
		user.DELETE("/:id", userHandler.DeleteUser)
	}

	r.POST("/login", userHandler.Login)

	return r
}
