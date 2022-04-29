package controllers

import (
	"eh-backend-api/domain/errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
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

type CustomValidator struct {
	validator *validator.Validate
}

// SEE:https://github.com/go-playground/validator
// https://pkg.go.dev/gopkg.in/go-playground/validator.v9#section-readme
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

type RequestMapper struct{}

func (it *RequestMapper) Parse(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}
	if err := c.Validate(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}
	return nil
}

func NewRequestMapper() RequestMapper {
	return RequestMapper{}
}
