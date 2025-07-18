package main

import (
	"github.com/B-Bridger/server/handler"
	"github.com/B-Bridger/server/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandler *handler.UserHandler, chatRoomHandler *handler.ChatRoomHandler) *gin.Engine {
	r := gin.Default()

	// 사용자 관련 라우팅 설정
	authRequiredUser := r.Group("/users", middleware.AuthMiddleware())
	{
		authRequiredUser.GET("/", userHandler.GetUser)
		authRequiredUser.PUT("/", userHandler.UpdateUser)
		authRequiredUser.DELETE("/", userHandler.DeleteUser)
		authRequiredUser.POST("/profile-image", userHandler.UploadProfileImage)
	}
	user := r.Group("/users")
	{
		user.POST("/", userHandler.CreateUser)
	}
	r.POST("/login", userHandler.Login)

	// 채팅방 관련 라우팅 설정
	authRequiredChatRoom := r.Group("/chat-room", middleware.AuthMiddleware())
	{
		authRequiredChatRoom.POST("/", chatRoomHandler.CreateChatRoom)
		authRequiredChatRoom.PUT("/:id", chatRoomHandler.UpdateChatRoom)
		authRequiredChatRoom.DELETE("/:id", chatRoomHandler.DeleteChatRoom)
	}
	chatRoom := r.Group("/chat-room")
	{
		chatRoom.GET("/:id", chatRoomHandler.GetChatRoom)
	}
	authRequiredChatRooms := r.Group("/chat-rooms", middleware.AuthMiddleware())
	{
		authRequiredChatRooms.GET("/", chatRoomHandler.GetChatRoomByOwner)
	}

	// Swagger & 정적 파일
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/static", "./static")

	return r
}
