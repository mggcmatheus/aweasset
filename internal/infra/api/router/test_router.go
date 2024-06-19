package router

import (
	"aweasset/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

func TestRouter(r *echo.Group) {
	h := handler.NewTestHandler()

	r.GET("/", h.Test)

}
