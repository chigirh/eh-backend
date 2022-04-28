package user

import (
	"context"
	"eh-backend-api/domain/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"eh-backend-api/app/usecases/ports"
)

type UserApi interface {
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type UserController struct {
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
}

func (it *UserController) newInputPort(c echo.Context) ports.UserInputPort {
	repository := it.repositoryFactory()
	return it.inputFactory(repository)
}

func (it *UserController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("api-key")
		fmt.Println(apiKey)

		userId := c.Param("userId")
		user, _ := it.newInputPort(c).GetUser(ctx, userId)
		res := new(GetResponse)
		res.User = UserDto{Id: user.UserId, FirstName: user.Firstname, FamilyName: user.FamilyName}
		return c.JSON(http.StatusOK, res)
	}
}

func (it *UserController) Post(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		req := new(PostRequest)
		if error := c.Bind(req); error != nil {
			return error
		}

		it.newInputPort(c).AddUser(ctx, &models.User{UserId: req.User.Id, Firstname: req.User.FirstName, FamilyName: req.User.FamilyName})

		res := PostResponse{}
		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type GetResponse struct {
	User UserDto `json:"user"`
}

type PostRequest struct {
	User UserDto `json:"user"`
}

type PostResponse struct {
}

type UserDto struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	FamilyName string `json:"family_name"`
}

// di
func NewUserController(
	inputFactory InputFactory,
	repositoryFactory RepositoryFactory,
) UserApi {
	return &UserController{
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
	}
}

type InputFactory func(ports.UserRepository) ports.UserInputPort
type RepositoryFactory func() ports.UserRepository
