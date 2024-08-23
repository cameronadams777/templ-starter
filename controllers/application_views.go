package controllers

import (
	"app/structs"
	"app/views/pages"

	"github.com/labstack/echo/v4"
)

type ApplicationViewHandler struct {}

func (av *ApplicationViewHandler) HandleHomeIndex(c echo.Context) error {
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
	return render_with_context(c, pages.HomeIndex(), app_context)
}

func (av *ApplicationViewHandler) HandleNotFoundIndex(c echo.Context) error {
  return render(c, pages.NotFoundIndex())
}

func (av *ApplicationViewHandler) HandleErrorIndex(c echo.Context) error {
  return render(c, pages.ErrorIndex())
}
