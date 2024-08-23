package controllers

import (
	"app/services"
	"app/views/components"
	"app/views/pages"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type AuthenticationController struct {
	DB *gorm.DB
}

func (ac *AuthenticationController) HandleLoginIndex(c echo.Context) error {
	return render(c, pages.LoginIndex(
		pages.LoginIndexPageProps{
			Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		},
	))
}

type LoginDTO struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (ac *AuthenticationController) HandleLoginCreate(c echo.Context) error {
	var form LoginDTO

	if err := c.Bind(&form); err != nil {
		fmt.Println(err)
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid email or password",
		}))
	}

	auth_service := services.AuthService{
		DB:  ac.DB,
		CTX: c,
	}
	err := auth_service.Login(form.Email, form.Password)

	if err != nil {
		fmt.Println(err)
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid email or password",
		}))
	}

	c.Response().Header().Set("HX-Location", "/")
	return c.String(http.StatusOK, "")
}

func (ac *AuthenticationController) HandleRegisterIndex(c echo.Context) error {
	return render(c, pages.RegisterIndex(
		pages.RegisterIndexPageProps{
			Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		},
	))
}

func (ac *AuthenticationController) HandleRegisterCreate(c echo.Context) error {
	var form services.RegisterDTO

	if err := c.Bind(&form); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid form data",
		}))
	}

	auth_service := services.AuthService{
		DB:  ac.DB,
		CTX: c,
	}
	err := auth_service.Register(form)

	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: errors.New("An error occurred while registering").Error(),
		}))
	}

	c.Response().Header().Set("HX-Location", "/login")
	return c.String(http.StatusOK, "")
}

func (ac *AuthenticationController) HandleForgotPasswordIndex(c echo.Context) error {
	return render(c, pages.ForgotPasswordIndex(
		pages.ForgotPasswordIndexPageProps{
			Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
    },
	))
}

func (ac *AuthenticationController) HandleSendResetPasswordEmail(c echo.Context) error {
	return nil
}

func (ac *AuthenticationController) HandleLogout(c echo.Context) error {
	auth_service := services.AuthService{
		DB:  ac.DB,
		CTX: c,
	}

	auth_service.SignOut()

	return c.Redirect(http.StatusPermanentRedirect, "/login")
}
