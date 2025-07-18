package service

import (
	"github.com/B-Bridger/server/model"
	"github.com/B-Bridger/server/repository"
)

// ChatRoomService는 채팅방 도메인과 관련된 비즈니스 로직을 담당합니다.
//
// Methods:
type ChatRoomService struct {
	Repo repository.ChatRoomRepository
}

// ChatRoomID를 통해 ChatRoom 객체를 반환합니다.
//
// 매개 변수
//   - id: 채팅방의 고유 ID
//
// 반환 값
//   - ChatRoom: 불러온 ChatRoom 객체
//   - error: 실패 시 error 메세지
func (s *ChatRoomService) GetChatRoomByID(id string) (*model.ChatRoom, error) {
	return s.Repo.FindByID(id)
}

// OwnerUserID를 통해 ChatRoom 객체를 반환합니다.
//
// 매개 변수
//   - id: 사용자의 고유 ID
//
// 반환 값
//   - ChatRoom: 불러온 ChatRoom 객체
//   - error: 실패 시 error 메세지
func (s *ChatRoomService) GetChatRoomByUserID(id string) (*[]model.ChatRoom, error) {
	return s.Repo.FindByOwner(id)
}

// chatRoom 객체를 데이터베이스에 저장합니다.
//
// 매개 변수
//   - chatRoom: ChatRoom 객체 포인터
//
// 반환 값
//   - error: 실패 시 error 메세지
func (s *ChatRoomService) CreateChatRoom(chatRoom *model.ChatRoom) error {
	return s.Repo.Create(chatRoom)
}

// chatRoom 객체를 데이터베이스 수정합니다.
//
// 매개 변수
//   - chatRoom: ChatRoom 객체 포인터
//
// 반환 값
//   - ChatRoom: 수정된 ChatRoom 객체
//   - error: 실패 시 error 메세지
func (s *ChatRoomService) UpdateChatRoom(chatRoom *model.ChatRoom) (*model.ChatRoom, error) {
	return s.Repo.Update(chatRoom)
}

// chatRoom 객체를 데이터베이스에서 제거합니다.
//
// 매개 변수
//   - id: 채팅방의 고유 ID
//
// 반환 값
//   - error: 실패 시 error 메세지
func (s *ChatRoomService) DeleteChatRoom(id string) error {
	return s.Repo.Delete(id)
}
