//go:build wireinject
// +build wireinject

package drivers

import (
	"context"
	"eh-backend-api/adapter/controllers/auth"
	"eh-backend-api/adapter/controllers/user"
	"eh-backend-api/adapter/gateways/mysql"
	"eh-backend-api/adapter/gateways/redis"
	"eh-backend-api/app/usecases/interactors"

	"github.com/google/wire"
	"github.com/labstack/echo"
)

func InitializeDriver(ctx context.Context) (Server, error) {
	wire.Build(
		// Driver
		NewDriver,
		// echo
		echo.New,
		// user
		user.NewUserController,
		interactors.NewUserInputPort,
		mysql.NewUserRepository,
		// auth
		auth.NewAuthController,
		interactors.NewAuthIputPort,
		mysql.NewAnthRepository,
		redis.NewTokenRepository,
	)
	return &Driver{}, nil
}
