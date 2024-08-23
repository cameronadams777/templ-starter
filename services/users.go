package services

import (
	"app/models"
	"app/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
  DB *gorm.DB
}

func (us *UserService) FindByID(id string) (*models.User, error) {
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

type UpdateUserParams struct {
  ID string `form:"id"`
  FirstName string `form:"first_name"`
  LastName string `form:"last_name"`
  Email string `form:"email"`
}

func (us *UserService) Update(updates UpdateUserParams) (*models.User, error) {
  user_repo := repositories.UserRepository {
    DB: us.DB,
  }

  updated_user := models.User {
    FirstName: updates.FirstName,
    LastName: updates.LastName,
    Email: updates.Email,
  }

  user_id, err := uuid.Parse(updates.ID)

  if err != nil {
    return nil, err
  }

  user, err := user_repo.Update(user_id, updated_user)

  if err != nil {
    return nil, err
  }

  return user, nil
}
