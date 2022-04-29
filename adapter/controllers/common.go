package controllers

import (
	"eh-backend-api/domain/errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type EmptyResponse struct{}

var DefaultResponse = EmptyResponse{}

func ErrorHandle(c echo.Context, err error) error {
	switch err.(type) {
	case *errors.NotFoundError:
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	case *errors.AuthenticationError:
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
	case *errors.SystemError:
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("Internal server error. message:%s", err.Error())})
}
