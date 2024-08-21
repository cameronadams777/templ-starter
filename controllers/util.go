package controllers

import (
	"context"
	"math/rand"
	"app/structs"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func render_with_context(c echo.Context, component templ.Component, app_context structs.AppContext) error {
	ctx := context.WithValue(context.Background(), app_context.Key, app_context.Value)
	return component.Render(ctx, c.Response())
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!@#$%^&()[]{};,.<>?+-="

func generate_token() string {
	// Generate a random 32 character string made up of letters, number and symbols
	b := make([]byte, 64)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
