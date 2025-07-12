package handler

import (
	"fmt"
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

// GetUser godoc
// @Summary 사용자 정보 조회
// @Description JWT 토큰에 포함된 사용자 정보를 가져옵니다.
// @Tags 사용자
// @Security BearerAuth
// @Success 200 {object} model.UserResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.MustGet("userID").(string)
	user, err := h.Service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "사용자를 찾을 수 없습니다", Detail: err.Error(), Status: 404})
		return
	}
	c.JSON(http.StatusOK, model.UserResponse{Message: "사용자를 성공적으로 조회하였습니다", Status: 200, User: *user})
}

// CreateUser godoc
// @Summary 사용자 생성
// @Description 새로운 사용자를 생성합니다.
// @Tags 사용자
// @Accept json
// @Produce json
// @Param user body model.CreateUserModel true "사용자 정보"
// @Success 201 {object} model.UserResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "요청 형식이 잘못되었습니다", Detail: err.Error(), Status: 400})
		return
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Language: req.Language,
		Password: req.Password,
	}

	if err := h.Service.CheckUserField(user.UserID, user.Email); err != nil {
		c.JSON(http.StatusConflict, model.ErrorResponse{Message: "해당 사용자가 이미 존재합니다", Detail: err.Error(), Status: 409})
		return
	}

	if err := h.Service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "사용자 생성에 실패하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	c.JSON(http.StatusCreated, model.UserResponse{Message: "사용자를 성공적으로 생성하였습니다", Status: 201, User: *user})
}

// UpdateUser godoc
// @Summary 사용자 정보 수정
// @Description JWT 토큰 기반으로 현재 사용자 정보를 수정합니다.
// @Tags 사용자
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body model.User true "수정할 사용자 정보"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.MustGet("userID").(string)
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "요청 형식이 잘못되었습니다", Detail: err.Error(), Status: 400})
		return
	}
	user.UserID = id

	updated, err := h.Service.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "유저 정보 수정에 실패하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{Message: "유저 정보를 성공적으로 수정하였습니다", Status: 200, User: *updated})
}

// DeleteUser godoc
// @Summary 사용자 삭제
// @Description JWT 토큰 기반으로 현재 사용자를 삭제합니다.
// @Tags 사용자
// @Security BearerAuth
// @Success 200 {object} model.OKResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.MustGet("userID").(string)
	if err := h.Service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "삭제 실패", Detail: err.Error(), Status: 500})
		return
	}
	c.JSON(http.StatusOK, model.OKResponse{Message: "삭제 완료", Status: 200})
}

// Login godoc
// @Summary 로그인
// @Description 이메일과 비밀번호를 사용하여 로그인합니다.
// @Tags 인증
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body LoginRequest true "로그인 정보"
// @Success 200 {object} model.TokenResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "요청 형식이 잘못되었습니다", Detail: err.Error(), Status: 400})
		return
	}

	user, token, err := h.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "로그인에 실패하였습니다", Detail: err.Error(), Status: 401})
		return
	}

	c.JSON(http.StatusOK, model.TokenResponse{Message: "로그인에 성공하였습니다", Status: 200, User: *user, Token: token})
}

// UploadProfileImage godoc
// @Summary 프로필 이미지 업로드
// @Description JWT 토큰에 기반한 사용자의 프로필 이미지를 업로드하고 경로를 DB에 저장합니다.
// @Tags 사용자
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "업로드할 이미지 파일"
// @Security BearerAuth
// @Success 200 {object} model.OKResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/profile-image [post]
func (h *UserHandler) UploadProfileImage(c *gin.Context) {
	id := c.MustGet("userID").(string)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "이미지 파일이 없습니다", Detail: err.Error(), Status: 400})
		return
	}

	filename := fmt.Sprintf("%s_%s", id, file.Filename)
	savePath := "static/uploads/" + filename

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "이미지 저장에 실패하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	imageURL := "/static/uploads/" + filename
	if err := h.Service.UpdateProfileImage(id, imageURL); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "이미지 저장에 실패하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	c.JSON(http.StatusOK, model.OKResponse{Message: "이미지를 성공적으로 저장하였습니다", Status: 200})
}
