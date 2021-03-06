//go:build wireinject
// +build wireinject

package drivers

import (
	"context"
	"eh-backend-api/adapter/controllers"
	"eh-backend-api/adapter/controllers/auth"
	"eh-backend-api/adapter/controllers/health"
	"eh-backend-api/adapter/controllers/schedule"
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
		// commons
		controllers.NewRequestMapper,
		// health
		health.NewHealthController,
		// user
		user.NewUserController,
		interactors.NewUserInputPort,
		mysql.NewUserRepository,
		// auth
		auth.NewAuthController,
		interactors.NewAuthIputPort,
		mysql.NewAnthRepository,
		redis.NewTokenRepository,
		// schedule
		schedule.NewScheduleController,
		interactors.NewScheduleInputPort,
		mysql.NewScheduleRepository,
	)
	return &Driver{}, nil
}
