package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"eh-backend-api/adapter/controllers"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"
)

type AuthApi interface {
	Post(ctx context.Context) func(c echo.Context) error
	Login(ctx context.Context) func(c echo.Context) error
}

type AuthController struct {
	inputPort ports.AuthInputPort
}

func (it *AuthController) Post(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(Request)
		if error := c.Bind(req); error != nil {
			return error
		}

		err := it.inputPort.UpdatePassword(
			ctx,
			models.UserName(req.UserName),
			models.Password(req.Password),
		)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

func (it *AuthController) Login(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(Request)
		if error := c.Bind(req); error != nil {
			return error
		}

		token, err := it.inputPort.AhtuAndCreateToken(
			ctx,
			models.UserName(req.UserName),
			models.Password(req.Password),
		)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := LoginResponse{SessionToken: *token}

		return c.JSON(http.StatusOK, res)
	}
}

// dto
type Request struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionToken models.SessionToken `json:"session_token"`
}

// di
type InputFactory func(ports.AuthRepository) ports.AuthInputPort
type RepositoryFactory func() ports.AuthRepository

func NewAuthController(inputPort ports.AuthInputPort) AuthApi {
	return &AuthController{
		inputPort: inputPort,
	}
}
