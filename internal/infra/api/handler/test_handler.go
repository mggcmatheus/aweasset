package handler

import (
	"aweasset/internal/infra/api/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c echo.Context) error {
	// Envia uma resposta JSON simples
	if err := c.JSON(http.StatusOK, echo.Map{
		"result": "Test",
	}); err != nil {
		return err
	}

	// Envia uma resposta JSON utilizando o helper.GenerateBaseResponse
	return c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))
}
