package mariaDB

import (
	"github.com/B-Bridger/server/model"
	"gorm.io/gorm"
)

type MariaDBUserRepository struct {
	DB *gorm.DB
}

func (r *MariaDBUserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	// TODO: SQL Injection 여부 확인 필요
	if err := r.DB.First(&user, "userID = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MariaDBUserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	// TODO: SQL Injection 여부 확인 필요
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MariaDBUserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *MariaDBUserRepository) Update(user *model.User) (*model.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MariaDBUserRepository) Delete(id string) error {
	// ToDO: SQL Injection 여부 확인 필요
	return r.DB.Delete(&model.User{}, "userID = ?", id).Error
}

func (r *MariaDBUserRepository) UpdateProfileImage(userID string, imageURL string) error {
	return r.DB.Model(&model.User{}).
		Where("userID = ?", userID).
		Update("profile", imageURL).
		Error
}
