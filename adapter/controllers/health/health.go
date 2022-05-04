package health

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"eh-backend-api/adapter/controllers"
)

type HealthApi interface {
	Get(ctx context.Context) func(c echo.Context) error
}

type HealthController struct{}

func (it *HealthController) Get(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

// di
func NewHealthController() HealthApi {
	return &HealthController{}
}
