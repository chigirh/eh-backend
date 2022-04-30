package user

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"

	"eh-backend-api/adapter/controllers"
	"eh-backend-api/conf/config"
)

type UserApi interface {
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type UserController struct {
	requestMapper controllers.RequestMapper
	authPort      ports.AuthInputPort
	inputPort     ports.UserInputPort
}

func (it *UserController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		userId := c.Param("userId")

		// If session token is set, have admin.
		token, err := it.requestMapper.GetSessionToken(c)
		if err != nil {
			return err
		}
		usrl, err := it.authPort.GetUserRole(ctx, token)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}
		if !usrl.HaveAdmin() {
			return c.JSON(http.StatusForbidden, controllers.ErrorResponse{Message: "Requires admin"})
		}

		user, err := it.inputPort.GetUser(ctx, models.UserName(userId))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := new(GetResponse)
		res.User = UserDto{
			Id:         string(user.UserId),
			FirstName:  user.Firstname,
			FamilyName: user.FamilyName,
			Password:   "",
		}
		for i := 0; i < len(user.Roles); i++ {
			res.User.Roles = append(res.User.Roles, string(user.Roles[i]))
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (it *UserController) Post(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		req := new(PostRequest)
		if err := it.requestMapper.Parse(c, req); err != nil {
			return err
		}

		// If session token is set, have admin.
		stkn := c.Request().Header.Get("x-session-token")
		if stkn != "" {
			usrl, err := it.authPort.GetUserRole(ctx, models.SessionToken(stkn))
			if err != nil {
				return controllers.ErrorHandle(c, err)
			}
			if !usrl.HaveAdmin() {
				return c.JSON(http.StatusForbidden, controllers.ErrorResponse{Message: "Requires admin"})
			}
		}

		// If master key is set, the master key must match.
		mk := c.Request().Header.Get("x-master-key")
		if mk != "" && mk != config.Config.Server.MasterKey {
			return c.JSON(http.StatusUnauthorized, controllers.ErrorResponse{Message: "Master key mismatch."})
		}

		if mk == "" && stkn == "" {
			return c.JSON(http.StatusBadRequest, controllers.ErrorResponse{Message: "Rrequires master key or sessiont token."})
		}

		roles := []models.Role{}
		for i := 0; i < len(req.User.Roles); i++ {
			roles = append(roles, models.Role(req.User.Roles[i]))
		}
		user := models.User{
			UserId:     models.UserName(req.User.Id),
			Firstname:  req.User.FirstName,
			FamilyName: req.User.FamilyName,
			Password:   models.Password(req.User.Password),
			Roles:      roles,
		}
		if err := it.inputPort.AddUser(ctx, user); err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := controllers.DefaultResponse
		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	GetResponse struct {
		User UserDto `json:"user" validate:"required"`
	}

	PostRequest struct {
		User UserDto `json:"user" validate:"required"`
	}

	UserDto struct {
		Id         string   `json:"id" validate:"required,max=64"`
		FirstName  string   `json:"first_name" validate:"required,max=300"`
		FamilyName string   `json:"family_name" validate:"required,max=300"`
		Password   string   `json:"password" validate:"required"`
		Roles      []string `json:"roles" validate:"min=1,unique"`
	}
)

// di
func NewUserController(
	requestMapper controllers.RequestMapper,
	authPort ports.AuthInputPort,
	inputPost ports.UserInputPort,
) UserApi {
	return &UserController{
		requestMapper: requestMapper,
		authPort:      authPort,
		inputPort:     inputPost,
	}
}
