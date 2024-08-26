package controllers

import (
	"app/services"
	"app/views/components"
	"app/views/pages/user_pages"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type UsersController struct {
	DB *gorm.DB
}

func (uc *UsersController) HandleUsersEdit(c echo.Context) error {
	user_id := c.Get("user_id")

	app_context := get_app_context(c)

	user_service := services.UserService{
		DB: uc.DB,
	}

	user, err := user_service.FindByID(user_id.(string))

	if err != nil {
		return c.Redirect(302, "/error")
	}

	return render_with_context(c, user_pages.UserEdit(user_pages.UserEditPageProps{
		Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		User:  *user,
	}), app_context)
}

func (uc *UsersController) HandleUsersUpdate(c echo.Context) error {
	var form services.UpdateUserParams

	if err := c.Bind(&form); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid form data",
		}))
	}

	user_service := services.UserService{
		DB: uc.DB,
	}

	_, err := user_service.Update(form)

	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid form data",
		}))
	}

	c.Response().Header().Set("HX-Location", "/users/edit")
	return c.String(http.StatusOK, "")
}
