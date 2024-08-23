package services

import (
	"errors"

	"app/models"
	"app/repositories"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SESSION_MAX_AGE = 86400 * 7

type AuthService struct {
	DB  *gorm.DB
	CTX echo.Context
}

func (s *AuthService) Login(email string, password string) error {
	user_repo := repositories.UserRepository{DB: s.DB}
	user, err := user_repo.GetUser(repositories.GetUserParams{Email: email})

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	sess, err := session.Get("app_session", s.CTX)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   SESSION_MAX_AGE,
		HttpOnly: true,
	}
	sess.Values["id"] = user.ID.String()

	session_save_err := sessions.Save(s.CTX.Request(), s.CTX.Response())

	if session_save_err != nil {
		return session_save_err
	}

	return nil
}

type RegisterDTO struct {
	FirstName       string `form:"first_name"`
	LastName        string `form:"last_name"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

func (s *AuthService) Register(params RegisterDTO) error {
	user_repo := repositories.UserRepository{DB: s.DB}

	if params.Password != params.ConfirmPassword {
		return errors.New("Passwords do not match")
	}

	hashed_password, hashing_err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if hashing_err != nil {
		return hashing_err
	}

	user := models.User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(hashed_password),
	}

	err := user_repo.CreateUser(user)

	if err != nil {
		return err
	}

	sess, err := session.Get("app_session", s.CTX)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["id"] = user.ID.String()

	session_save_err := sessions.Save(s.CTX.Request(), s.CTX.Response())

	if session_save_err != nil {
		return session_save_err
	}

	return nil
}

func (s *AuthService) SignOut() error {
	sess, err := session.Get("app_session", s.CTX)

  if err != nil {
    return err
  }

	sess.Options.MaxAge = -1
	session_save_err := sess.Save(s.CTX.Request(), s.CTX.Response())

	if session_save_err != nil {
		return session_save_err
	}

	return nil
}
