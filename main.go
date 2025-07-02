package main

import (
	"log"
	"os"

	"github.com/B-Bridger/server/database"
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

	userRepo := &mariaDB.MariaDBUserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	r := SetupRouter(userHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("서버 실행 중: http://localhost:%s", port)
	r.Run(":" + port)
}
