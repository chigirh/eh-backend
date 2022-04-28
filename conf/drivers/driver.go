package drivers

import (
	"context"
	"eh-backend-api/adapter/controllers/user"
	"eh-backend-api/conf/config"
	"fmt"

	"github.com/labstack/echo"
)

type Server interface {
	Start(ctx context.Context)
}

type UserDriver struct {
	echo       *echo.Echo
	controller user.UserApi
}

func NewUserDriver(echo *echo.Echo, controller user.UserApi) Server {
	return &UserDriver{
		echo:       echo,
		controller: controller,
	}
}

func (driver *UserDriver) Start(ctx context.Context) {
	driver.echo.GET("/users/:userId", driver.controller.Get(ctx))
	driver.echo.POST("/users", driver.controller.Post(ctx))
	driver.echo.Logger.Fatal(driver.echo.Start(fmt.Sprintf(":%d", config.Config.Server.ServerPort)))
}
