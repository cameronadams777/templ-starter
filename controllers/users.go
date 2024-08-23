package controllers

import (
	"app/services"
	"app/structs"
	"app/views/pages/user_pages"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersController struct {
	DB *gorm.DB
}

func (uc UsersController) HandleUsersEdit(c echo.Context) error {
	user_id := c.Get("user_id")

	if user_id == nil {
		user_id = ""
	}

	app_context := structs.AppContext{
		Key: "session",
		Value: structs.SessionContext{
			UserID: user_id.(string),
		},
	}

	user_service := services.UserService{
		DB: uc.DB,
	}

  user, err := user_service.FindByID(user_id.(string))

  if err != nil {
    return c.Redirect(302, "/error")
  }

	return render_with_context(c, user_pages.UserEdit(user_pages.UserEditPageProps{
		User: *user,
	}), app_context)
}
