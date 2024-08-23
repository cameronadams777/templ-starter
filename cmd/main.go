package main

import (
	"app/controllers"
	"app/database"
  middleware_handlers "app/middleware"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

  db := database.ConnectDB()

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})

	app.Static("/assets", "assets")

	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
	}))

  application_views_controller := controllers.ApplicationViewHandler{}
  app.GET("/error", application_views_controller.HandleErrorIndex)
  app.GET("/not_found", application_views_controller.HandleNotFoundIndex)
  app.GET("/", application_views_controller.HandleHomeIndex)

  authentication_controller := controllers.AuthenticationController {
    DB: db,
  }
  app.GET("/login", authentication_controller.HandleLoginIndex)
  app.POST("/auth/login", authentication_controller.HandleLoginCreate)
  app.GET("/register", authentication_controller.HandleRegisterIndex)
  app.POST("/auth/register", authentication_controller.HandleRegisterCreate)
  app.GET("/forgot-password", authentication_controller.HandleForgotPasswordIndex)
  app.GET("/auth/send-reset-password-email", authentication_controller.HandleSendResetPasswordEmail)
  app.GET("/logout", authentication_controller.HandleLogout)

	app.Use(middleware_handlers.SetSessionInContext)
	app.Use(middleware_handlers.NoSessionRedirect)

  users := app.Group("/users")
  users_controller := controllers.UsersController {
    DB: db,
  }
  users.GET("/edit", users_controller.HandleUsersEdit)

	app.Logger.Fatal(app.Start(":4000"))
}
