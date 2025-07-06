package repository

import "github.com/B-Bridger/server/model"

// User 관련 데이터 엑세스를 추상화한 인터페이스입니다.
// SOLID 원칙에 따라, Interface 구현은 mariaDB에서 진행합니다.
type UserRepository interface {
	// UserID를 통해 user 객체를 반환합니다.
	//
	// 매개 변수
	//   - id: 사용자의 고유 ID
	//
	// 반환 값
	//   - *User: 불러온 user 객체
	//   - error: 실패 시 error 메세지
	FindByID(id string) (*model.User, error)

	// email을 통해 user 객체를 반환합니다.
	//
	// 매개 변수
	//   - id: 사용자의 고유 ID
	//
	// 반환 값
	//   - *User: 불러온 user 객체
	//   - error: 실패 시 error 메세지
	FindByEmail(email string) (*model.User, error)

	// 사용자 레코드를 생성합니다.
	//
	// 매개 변수
	//   - user: user 객체 포인터
	//
	// 반환 값
	//   - error: 실패 시 error 메세지
	Create(user *model.User) error

	// 기존에 존재하는 사용자 정보를 수정합니다.
	//
	// 매개 변수
	//   - user: user 객체 포인터
	//
	// 반환 값
	//   - *User: 수정된 user 객체
	//   - error: 실패 시 error 메세지
	Update(user *model.User) (*model.User, error)

	// 사용자 레코드를 삭제합니다.
	//
	// 매개 변수
	//   - id: 사용자의 고유 ID
	//
	// 반환 값
	//   - error: 실패 시 error 메세지
	Delete(id string) error

	// 사용자의 프로필 이미지 경로를 저장합니다.
	//
	// 매개 변수
	// 	 - id: 사용자의 고유 ID
	// 	 - imageUrl: 이미지 경로
	//
	// 반환 값
	//   - error: 실패 시 error 메세지
	UpdateProfileImage(id string, imageURL string) error
}
