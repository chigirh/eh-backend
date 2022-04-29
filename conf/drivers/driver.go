package drivers

import (
	"context"
	"eh-backend-api/adapter/controllers/auth"
	"eh-backend-api/adapter/controllers/user"
	"eh-backend-api/conf/config"
	"fmt"

	"github.com/labstack/echo"
)

type Server interface {
	Start(ctx context.Context)
}

type Driver struct {
	echo           *echo.Echo
	userController user.UserApi
	authController auth.AuthApi
}

func NewDriver(
	echo *echo.Echo,
	userController user.UserApi,
	authController auth.AuthApi,
) Server {
	return &Driver{
		echo:           echo,
		userController: userController,
		authController: authController,
	}
}

func (driver *Driver) Start(ctx context.Context) {
	// users
	driver.echo.GET("/users/:userId", driver.userController.Get(ctx))
	driver.echo.POST("/users", driver.userController.Post(ctx))

	// auth
	driver.echo.POST("/login", driver.authController.Login(ctx))
	driver.echo.POST("/auth/", driver.authController.Post(ctx))

	driver.echo.Logger.Fatal(driver.echo.Start(fmt.Sprintf(":%d", config.Config.Server.ServerPort)))
}
