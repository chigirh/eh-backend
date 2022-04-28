package controllers

import (
	"eh-backend-api/domain/errors"
	"net/http"

	"github.com/labstack/echo"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ErrorHandle(c echo.Context, err error) error {
	switch err {
	case *errors.NotFoundError:
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	case *errors.AuthenticationError:
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
	case *errors.SystemError:
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: errors.SystemError{"illegal error"}.Error()})
}
