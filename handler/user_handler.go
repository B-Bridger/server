package handler

import (
	"net/http"

	"github.com/B-Bridger/server/model"
	"github.com/B-Bridger/server/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "사용자를 찾을 수 없습니다"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "요청 형식이 잘못되었습니다"})
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "사용자 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청"})
		return
	}
	user.UserID = id

	updated, err := h.Service.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "업데이트 실패"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "삭제 실패"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "삭제 완료"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식"})
		return
	}

	user, token, err := h.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
