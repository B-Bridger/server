package handler

import (
	"fmt"
	"net/http"

	"github.com/B-Bridger/server/model"
	"github.com/B-Bridger/server/service"
	"github.com/gin-gonic/gin"
)

type ChatRoomHandler struct {
	Service *service.ChatRoomService
}

// GetChatRoom godoc
// @Summary 채팅방 정보 조회
// @Description 채팅방 고유 ID를 통해 채팅방 정보를 조회합니다.
// @Tags 채팅방
// @Param id path string true "조회할 채팅방 고유 ID"
// @Success 200 {object} model.ChatRoomResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /chat-room/{id} [get]
func (h *ChatRoomHandler) GetChatRoom(c *gin.Context) {
	id := c.Param("id")
	chatRoom, err := h.Service.GetChatRoomByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "채팅방을 찾을 수 없습니다.", Detail: err.Error(), Status: 404})
		return
	}
	c.JSON(http.StatusOK, model.ChatRoomResponse{Message: "채팅방을 성공적으로 조회하였습니다", Status: 200, ChatRoom: *chatRoom})
}

// GetChatRoomByOwner godoc
// @Summary 채팅방 정보 조회
// @Description 소유자 고유 ID를 통해 채팅방 정보를 조회합니다.
// @Tags 채팅방
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.ChatRoomsResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 403 {object} model.ErrorResponse
// @Router /chat-rooms [get]
func (h *ChatRoomHandler) GetChatRoomByOwner(c *gin.Context) {
	id := c.MustGet("userID").(string)
	chatRooms, err := h.Service.GetChatRoomByUserID(id)
	if err != nil {
		c.JSON(http.StatusForbidden, model.ErrorResponse{Message: "접근 권한이 없습니다.", Detail: err.Error(), Status: 403})
		return
	}

	c.JSON(http.StatusOK, model.ChatRoomsResponse{Message: "채팅방을 성공적으로 조회하였습니다", Status: 200, ChatRooms: *chatRooms})
}

// CreateChatRoom godoc
// @Summary 채팅방 생성
// @Description 새로운 채팅방을 생성합니다.
// @Tags 채팅방
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body model.CreateChatRoomModel true "채팅방 정보"
// @Success 201 {object} model.ChatRoomResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /chat-room [post]
func (h *ChatRoomHandler) CreateChatRoom(c *gin.Context) {
	id := c.MustGet("userID").(string)
	var req model.CreateChatRoomModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "요청 형식이 잘못되었습니다", Detail: err.Error(), Status: 400})
		return
	}

	chatRoom := model.ChatRoom{
		UserID: id,
	}

	if err := h.Service.CreateChatRoom(&chatRoom); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "채팅방 생성에 실패하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	readChatRoom, err := h.Service.GetChatRoomByID(chatRoom.ChatRoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "요청을 처리하는 중 오류가 발생하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	c.JSON(http.StatusCreated, model.ChatRoomResponse{Message: "채팅방을 성공적으로 생성하였습니다", Status: 201, ChatRoom: *readChatRoom})
}

// UpdateChatRoom Docs
// @Summary 채팅방 정보 업데이트
// @Description 채팅방 정보를 수정합니다. JWT Token을 기반으로 Owner가 아닐 경우, 403 코드를 반환합니다.
// @Tags 채팅방
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "수정할 채팅방 고유 ID"
// @Param chatRoom body model.ChatRoom true "수정할 채팅방 정보"
// @Success 200 {object} model.ChatRoomResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 403 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /chat-room/{id} [PUT]
func (h *ChatRoomHandler) UpdateChatRoom(c *gin.Context) {
	// FIXME: 업데이트 로직 변경
	userID := c.MustGet("userID").(string)
	chatRoomID := c.Param("id")
	fmt.Println(chatRoomID)
	var req model.ChatRoom
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "요청 형식이 잘못되었습니다", Detail: err.Error(), Status: 400})
		return
	}

	if userID != req.UserID {
		c.JSON(http.StatusForbidden, model.ErrorResponse{Message: "접근 권한이 없습니다", Detail: "JWT token value doesn't match Owner", Status: 403})
		return
	}

	updatedChatRoom, err := h.Service.UpdateChatRoom(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "정보 수정중 오류가 발생하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	c.JSON(http.StatusOK, model.ChatRoomResponse{Message: "정보를 성공적으로 수정하였습니다", Status: 200, ChatRoom: *updatedChatRoom})
}

// DeleteChatRoom godoc
// @Summary 채팅방 제거
// @Description 채팅방 정보를 제거합니다. JWT Token을 기반으로 Owner가 아닐 경우, 403 코드를 반환합니다.
// @Tags 채팅방
// @Security BearerAuth
// @Param id path string true "제거할 채팅방 고유 ID"
// @Success 200 {object} model.OKResponse
// @Failure 403 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /chat-room/{id} [delete]
func (h *ChatRoomHandler) DeleteChatRoom(c *gin.Context) {
	id := c.MustGet("userID").(string)
	chatRoomID := c.Param("id")
	chatRoom, err := h.Service.GetChatRoomByID(chatRoomID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "채팅방을 찾을 수 없습니다", Detail: err.Error(), Status: 404})
		return
	}

	if chatRoom.UserID != id {
		c.JSON(http.StatusForbidden, model.ErrorResponse{Message: "접근 권한이 없습니다", Detail: "You are not chat room's owner", Status: 403})
		return
	}

	err = h.Service.DeleteChatRoom(chatRoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "제거에 실패하였습니다", Detail: err.Error(), Status: 500})
		return
	}

	c.JSON(http.StatusOK, model.OKResponse{Message: "성공적으로 채팅방을 제거하였습니다", Status: 200})
}
