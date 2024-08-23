package repositories

import (
	"app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type FindUserParams struct {
	ID    string
	Email string
}

func (s *UserRepository) FindUser(params FindUserParams) (*models.User, error) {
	var user models.User

  tx := s.DB

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

func (s *UserRepository) CreateUser(user models.User) error {
	err := s.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}
