package main

// @title B-Bridger API 문서
// @version 1.0
// @description B2B 통신 플랫폼을 위한 백엔드 API
// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Example: "Bearer {token}"

import (
	"log"
	"os"

	"github.com/B-Bridger/server/database"
	_ "github.com/B-Bridger/server/docs"
	"github.com/B-Bridger/server/handler"
	"github.com/B-Bridger/server/model"
	"github.com/B-Bridger/server/repository/mariaDB"
	"github.com/B-Bridger/server/service"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.ChatRoom{})

	userRepo := &mariaDB.MariaDBUserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}
	chatRoomRepo := &mariaDB.MariaDBChatRoomRepository{DB: db}
	chatRoomService := &service.ChatRoomService{Repo: chatRoomRepo}
	chatRoomHandler := &handler.ChatRoomHandler{Service: chatRoomService}

	r := SetupRouter(userHandler, chatRoomHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("서버 실행 중: http://localhost:%s", port)
	r.Run(":" + port)
}
