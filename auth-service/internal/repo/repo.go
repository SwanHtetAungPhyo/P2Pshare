package repo

import (
	"github.com/SwanHtetAungPhyo/auth-service/internal/database"
	"github.com/SwanHtetAungPhyo/auth-service/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) Create(user *model.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) EmailExists(email string) (bool, error) {
	var count int64
	if err := database.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
