package mariaDB

import (
	"github.com/B-Bridger/server/model"
	"gorm.io/gorm"
)

type MariaDBChatRoomRepository struct {
	DB *gorm.DB
}

// TODO: Owner 불러오기 실행
func (r *MariaDBChatRoomRepository) FindByID(id string) (*model.ChatRoom, error) {
	var chatRoom model.ChatRoom

	if err := r.DB.Preload("Owner").First(&chatRoom, "chatRoomID = ?", id).Error; err != nil {
		return nil, err
	}

	return &chatRoom, nil
}

func (r *MariaDBChatRoomRepository) FindByOwner(id string) (*[]model.ChatRoom, error) {
	var chatRooms []model.ChatRoom

	if err := r.DB.Preload("Owner").Find(&chatRooms, "ownerUserID = ?", id).Error; err != nil {
		return nil, err
	}

	return &chatRooms, nil
}

func (r *MariaDBChatRoomRepository) Create(chatRoom *model.ChatRoom) error {
	return r.DB.Create(chatRoom).Error
}

func (r *MariaDBChatRoomRepository) Update(chatRoom *model.ChatRoom) (*model.ChatRoom, error) {
	if err := r.DB.Save(chatRoom).Error; err != nil {
		return nil, err
	}
	return chatRoom, nil
}

func (r *MariaDBChatRoomRepository) Delete(id string) error {
	return r.DB.Delete(&model.ChatRoom{}, "chatRoomID = ?", id).Error
}
