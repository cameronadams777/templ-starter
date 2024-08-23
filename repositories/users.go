package repositories

import (
	"app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
  DB *gorm.DB
}

type GetUserParams struct {
  Email string
}

func (s *UserRepository) GetUser(params GetUserParams) (*models.User, error) {
  var user models.User

  err := s.DB.Where("email = ?", params.Email).First(&user).Error

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

