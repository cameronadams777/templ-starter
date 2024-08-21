package controllers

import (
	"app/views/pages"

	"github.com/labstack/echo/v4"
)

type ApplicationViewHandler struct {}

func (av *ApplicationViewHandler) HandleHomeIndex(c echo.Context) error {
  return render(c, pages.HomeIndex())
}

func (av *ApplicationViewHandler) HandleNotFoundIndex(c echo.Context) error {
  return render(c, pages.NotFoundIndex())
}

func (av *ApplicationViewHandler) HandleErrorIndex(c echo.Context) error {
  return render(c, pages.ErrorIndex())
}
