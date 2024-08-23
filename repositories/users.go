package repositories

import (
	"app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type FindUserParams struct {
	ID    string
	Email string
}

func (ur *UserRepository) FindUser(params FindUserParams) (*models.User, error) {
	var user models.User

  tx := ur.DB

  if params.ID != "" {
    tx = tx.Where("id = ?", params.ID)
  }

  if params.Email != "" {
    tx = tx.Where("email = ?", params.Email)
  }

  err := tx.First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(user models.User) error {
	err := ur.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(id uuid.UUID, updated_user models.User) (*models.User, error) {
  var user_to_update models.User
  err := ur.DB.First(&user_to_update, id).Error

  if err != nil {
    return nil, err
  }

  ur.DB.Model(&user_to_update).Updates(updated_user)

  return &user_to_update, nil
}
