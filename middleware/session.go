package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func NoSessionRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		no_session_paths := []string{"/", "/login", "/auth/login", "/register", "/auth/register", "/not_found", "/error"}
		if strings.Contains(c.Path(), "/assets") || slices.Contains(no_session_paths, c.Path()) {
			return next(c)
		}

		sess, err := session.Get("app_session", c)
		if err != nil {
			return err
		}

		if sess.Values["id"] == nil {
			return c.Redirect(http.StatusPermanentRedirect, "/login")
		}

		return next(c)
	}
}

func SetSessionInContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("app_session", c)

		if err != nil {
			return next(c)
		}

    c.Get("user_id")

    c.Set("user_id", sess.Values["id"])

    return next(c)
	}
}
