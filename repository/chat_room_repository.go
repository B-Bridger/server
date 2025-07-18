package repository

import "github.com/B-Bridger/server/model"

// ChatRoom 관련 데이터 엑세스를 추상화한 인터페이스입니다.
type ChatRoomRepository interface {
	// ChatRoomID를 통해 ChatRoom 객체를 반환합니다.
	//
	// 매개 변수
	//   - id: 채팅방의 고유 ID
	//
	// 반환 값
	//   - ChatRoom: 불러온 ChatRoom 객체
	//   - error: 실패 시 error 메세지
	FindByID(id string) (*model.ChatRoom, error)

	// OwnerUserID를 통해 ChatRoom 객체를 반환합니다.
	//
	// 매개 변수
	//   - id: 사용자의 고유 ID
	//
	// 반환 값
	//   - ChatRoom: 불러온 ChatRoom 객체
	//   - error: 실패 시 error 메세지
	FindByOwner(id string) (*[]model.ChatRoom, error)

	// 채팅방 레코드를 생성합니다.
	//
	// 매개 변수
	//   - chatRoom: ChatRoom 객체 포인터
	//
	// 반환 값
	//   - error: 실패 시 error 메세지
	Create(chatRoom *model.ChatRoom) error

	// 기존에 존재하는 채팅방 정보를 변경합니다.
	//
	// 매개 변수
	//   - chatRoom: ChatRoom 객체 포인터
	//
	// 반환 값
	//   - ChatRoom: 수정된 ChatRoom 객체
	//   - error: 실패 시 error 메세지
	Update(chatRoom *model.ChatRoom) (*model.ChatRoom, error)

	// 채팅방 레코드를 삭제합니다.
	//
	// 매개 변수
	//   - id: 채팅방의 고유 ID
	//
	// 반환 값
	//   - error: 실패 시 error 메세지
	Delete(id string) error
}
