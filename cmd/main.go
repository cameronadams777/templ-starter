package main

import (
	"app/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))

	app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})

	app.Static("/assets", "assets")

	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
	}))

  authentication_controller := controllers.AuthenticationController {}
  app.GET("/login", authentication_controller.HandleLoginIndex)
  app.POST("/auth/login", authentication_controller.HandleLoginCreate)
  app.GET("/register", authentication_controller.HandleRegisterIndex)
  app.POST("/auth/register", authentication_controller.HandleRegisterCreate)
  app.GET("/forgot-password", authentication_controller.HandleForgotPasswordIndex)
  app.GET("/auth/send-reset-password-email", authentication_controller.HandleSendResetPasswordEmail)

  application_views_controller := controllers.ApplicationViewHandler{}
  app.GET("/error", application_views_controller.HandleErrorIndex)
  app.GET("/not_found", application_views_controller.HandleNotFoundIndex)
  app.GET("/", application_views_controller.HandleHomeIndex)

	app.Logger.Fatal(app.Start(":4000"))
}
