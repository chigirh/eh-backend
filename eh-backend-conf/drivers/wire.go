//go:build wireinject
// +build wireinject

package drivers

import (
	"adapter/controllers"
	"adapter/gateways/mysql"
	"app/usecases/interactors"
	"context"

	"github.com/google/wire"
	"github.com/labstack/echo"
)

func InitializeUserDriver(ctx context.Context) (User, error) {
	wire.Build(echo.New, NewInputFactory, NewRepositoryFactory, controllers.NewUserController, NewUserDriver)
	return &UserDriver{}, nil
}

func NewInputFactory() controllers.InputFactory {
	return interactors.NewUserInputPort
}

func NewRepositoryFactory() controllers.RepositoryFactory {
	return mysql.NewUserRepository
}
