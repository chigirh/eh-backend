package drivers

import (
	"context"
	"eh-backend-api/adapter/controllers"
	"eh-backend-api/adapter/controllers/auth"
	"eh-backend-api/adapter/controllers/schedule"
	"eh-backend-api/adapter/controllers/user"
	"eh-backend-api/conf/config"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

type Server interface {
	Start(ctx context.Context)
}

type Driver struct {
	echo        *echo.Echo
	userApi     user.UserApi
	authApi     auth.AuthApi
	scheduleApi schedule.ScheduleApi
}

func NewDriver(
	echo *echo.Echo,
	userApi user.UserApi,
	authApi auth.AuthApi,
	scheduleApi schedule.ScheduleApi,
) Server {
	return &Driver{
		echo:        echo,
		userApi:     userApi,
		authApi:     authApi,
		scheduleApi: scheduleApi,
	}
}

func (driver *Driver) Start(ctx context.Context) {
	log.Println("api start.")
	// custom validator
	driver.echo.Validator = controllers.NewValidator()

	// users
	driver.echo.GET("/users/:userId", driver.userApi.Get(ctx))
	driver.echo.POST("/users", driver.userApi.Post(ctx))

	// auth
	driver.echo.POST("/login", driver.authApi.Login(ctx))
	driver.echo.POST("/auth", driver.authApi.Post(ctx))

	// shcedule
	driver.echo.GET("/schedules/aggregate", driver.scheduleApi.AggregateGet(ctx))
	driver.echo.GET("/schedules/details", driver.scheduleApi.DetailsGet(ctx))
	driver.echo.GET("/schedules/periods", driver.scheduleApi.PeriodsGet(ctx))
	driver.echo.POST("/schedules", driver.scheduleApi.Post(ctx))

	driver.echo.Logger.Fatal(driver.echo.Start(fmt.Sprintf(":%d", config.Config.Server.ServerPort)))
}
