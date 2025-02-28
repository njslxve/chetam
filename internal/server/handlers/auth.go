package handlers

import (
	"chetam/internal/model"
	"chetam/internal/validation"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthInterface interface {
	CreateUser(model.RegisterRequest) (string, error)
}

func Register(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.RegisterRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err := validation.ValidateAuthRequest(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		token, err := auth.CreateUser(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var resp model.RegisterResponse
		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}

func Login(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusBadGateway, "fsddsfs")
	}
}
