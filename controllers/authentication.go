package controllers

import (
	"app/views/pages"

	"github.com/labstack/echo/v4"
)

type AuthenticationController struct {}

func (ac *AuthenticationController) HandleLoginIndex(c echo.Context) error {
  return render(c, pages.LoginIndex(
    pages.LoginIndexPageProps{},
  ))
}

func (ac *AuthenticationController) HandleLoginCreate(c echo.Context) error {
  return nil
}

func (ac *AuthenticationController) HandleRegisterIndex(c echo.Context) error {
  return render(c, pages.RegisterIndex(
    pages.RegisterIndexPageProps{},
  ))
}

func (ac *AuthenticationController) HandleRegisterCreate(c echo.Context) error {
  return nil
}

func (ac *AuthenticationController) HandleForgotPasswordIndex(c echo.Context) error {
  return render(c, pages.ForgotPasswordIndex(
    pages.ForgotPasswordIndexPageProps{},
  ))
}

func (ac *AuthenticationController) HandleSendResetPasswordEmail(c echo.Context) error {
  return nil
}


