package services

import (
	"app/models"
	"app/repositories"

	"gorm.io/gorm"
)

type UserService struct {
  DB *gorm.DB
}

func (us UserService) FindByID(id string) (*models.User, error) {
  user_repo := repositories.UserRepository{
    DB: us.DB,
  }

  user, err := user_repo.FindUser(repositories.FindUserParams{
    ID: id,
  })

  if err != nil {
    return nil, err
  }

  return user, nil
}
