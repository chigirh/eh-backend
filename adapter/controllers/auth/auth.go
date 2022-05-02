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
	requestMapper controllers.RequestMapper
	inputPort     ports.AuthInputPort
}

func (it *AuthController) Post(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(PostRequest)
		if error := it.requestMapper.Parse(c, req); error != nil {
			return error
		}

		// If session token is set, have any.
		token, err := it.requestMapper.GetSessionToken(c)
		if err != nil {
			return err
		}
		usrl, err := it.inputPort.GetUserRole(ctx, token)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		err = it.inputPort.UpdatePassword(
			ctx,
			usrl.UserName,
			models.Password(req.Before),
			models.Password(req.After),
		)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

func (it *AuthController) Login(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(LoginRequest)
		if error := it.requestMapper.Parse(c, req); error != nil {
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
type (
	LoginRequest struct {
		UserName string `json:"user_name" validate:"required,max=64"`
		Password string `json:"password" validate:"required"`
	}

	PostRequest struct {
		Before string `json:"before" validate:"required"`
		After  string `json:"after" validate:"required"`
	}

	LoginResponse struct {
		SessionToken models.SessionToken `json:"session_token"`
	}
)

// di
func NewAuthController(requestMapper controllers.RequestMapper, inputPort ports.AuthInputPort) AuthApi {
	return &AuthController{
		requestMapper: requestMapper,
		inputPort:     inputPort,
	}
}
