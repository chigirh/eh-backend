package drivers

import (
	"adapter/controllers"
	"context"

	"github.com/labstack/echo"
)

type User interface {
	ServeUsers(ctx context.Context, address string)
}

type UserDriver struct {
	echo       *echo.Echo
	controller controllers.UserApi
}

func NewUserDriver(echo *echo.Echo, controller controllers.UserApi) User {
	return &UserDriver{
		echo:       echo,
		controller: controller,
	}
}

func (driver *UserDriver) ServeUsers(ctx context.Context, address string) {
	driver.echo.GET("/users/:userId", driver.controller.Get(ctx))
	driver.echo.POST("/users", driver.controller.Post(ctx))
	driver.echo.Logger.Fatal(driver.echo.Start(address))
}
