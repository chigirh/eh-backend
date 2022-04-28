//go:build wireinject
// +build wireinject

package drivers

import (
	"context"
	"eh-backend-api/adapter/controllers/user"
	"eh-backend-api/adapter/gateways/mysql"
	"eh-backend-api/app/usecases/interactors"

	"github.com/google/wire"
	"github.com/labstack/echo"
)

func InitializeUserDriver(ctx context.Context) (Server, error) {
	wire.Build(echo.New, NewInputFactory, NewRepositoryFactory, user.NewUserController, NewUserDriver)
	return &UserDriver{}, nil
}

func NewInputFactory() user.InputFactory {
	return interactors.NewUserInputPort
}

func NewRepositoryFactory() user.RepositoryFactory {
	return mysql.NewUserRepository
}
