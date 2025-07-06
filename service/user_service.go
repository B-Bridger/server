package service

import (
	"errors"
	"os"
	"time"

	"github.com/B-Bridger/server/model"
	"github.com/B-Bridger/server/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// UserService는 사용자 도메인과 관련된 비즈니스 로직을 담당합니다.
// 이 서비스는 UserRepository 인터페이스에 의존하여 DB 구현과 분리된 구조를 가집니다.
//
// Methods:
//   - GetUser (사용자 조회)
//   - CreateUser (사용자 생성)
//   - UpdateUser (사용자 정보 수정)
//   - DeleteUser (사용자 삭제)
//   - Authenticate (로그인 인증)
type UserService struct {
	Repo repository.UserRepository
}

// UserID를 통해 user 객체를 반환합니다.
//
// 매개 변수
//   - id: 사용자의 고유 ID
//
// 반환 값
//   - *User: 불러온 user 객체
//   - error: 실패 시 error 메세지
func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.Repo.FindByID(id)
}

// 사용자를 생성합니다.
// 비밀번호는 bcrypt를 통해 암호화 합니다.
//
// 매개 변수
//   - user: user 객체 포인터
//
// 반환 값
//   - error: 실패 시 error 메세지
func (s *UserService) CreateUser(user *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.Repo.Create(user)
}

// 기존에 존재하는 사용자 정보를 수정합니다.
//
// 매개 변수
//   - user: user 객체 포인터
//
// 반환 값
//   - *User: 수정된 user 객체
//   - error: 실패 시 error 메세지
func (s *UserService) UpdateUser(user *model.User) (*model.User, error) {
	return s.Repo.Update(user)
}

// 사용자 레코드를 삭제합니다.
//
// 매개 변수
//   - id: 사용자의 고유 ID
//
// 반환 값
//   - error: 실패 시 error 메세지
func (s *UserService) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}

// Authenticate는 주어진 이메일과 비밀번호를 검증하여 로그인 인증을 수행합니다.
// 비밀번호는 bcrypt로 비교되며, 인증에 성공하면 사용자 정보와 토큰을 반환합니다.
//
// 매개 변수:
//   - email: 사용자의 이메일 주소
//   - password: 사용자의 비밀번호 (평문)
//
// 반환 값:
//   - *User: 인증된 사용자 정보
//   - string: 인증 성공 시 발급되는 토큰 문자열
//   - error: 인증 실패 시 오류 메시지 반환
func (s *UserService) Authenticate(email, password string) (*model.User, string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("사용자를 찾을 수 없습니다")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("비밀번호가 일치하지 않습니다")
	}

	jwtSecret := os.Getenv("SECRET")
	if jwtSecret == "" {
		return nil, "", errors.New("JWT 비밀 키가 설정되지 않았습니다")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.UserID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, "", errors.New("토큰 생성 실패")
	}

	return user, tokenString, nil
}

// CheckUserField는
func (s *UserService) CheckUserField(userID, email string) error {
	if findUser, _ := s.Repo.FindByID(userID); findUser != nil {
		return errors.New("userID가 이미 존재합니다")
	}
	if findUser, _ := s.Repo.FindByEmail(email); findUser != nil {
		return errors.New("email이 이미 존재합니다")
	}

	return nil
}
