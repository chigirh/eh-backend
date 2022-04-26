package controllers

import (
	"context"
	"domain/models"
	"net/http"

	"github.com/labstack/echo"

	"app/usecases/ports"
)

type UserApi interface {
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type InputFactory func(ports.UserRepository) ports.UserInputPort
type RepositoryFactory func() ports.UserRepository

type UserController struct {
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
}

func NewUserController(
	inputFactory InputFactory,
	repositoryFactory RepositoryFactory,
) UserApi {
	return &UserController{
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
	}
}

func (u *UserController) newInputPort(c echo.Context) ports.UserInputPort {
	repository := u.repositoryFactory()
	return u.inputFactory(repository)
}

func (u *UserController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		userId := c.Param("userId")
		user, _ := u.newInputPort(c).GetUser(ctx, userId)
		res := new(GetResponse)
		res.User = UserDto{Id: user.UserId, FirstName: user.Firstname, FamilyName: user.FamilyName}
		return c.JSON(http.StatusOK, res)
	}
}

func (u *UserController) Post(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		req := new(PostRequest)
		if error := c.Bind(req); error != nil {
			return error
		}

		u.newInputPort(c).AddUser(ctx, &models.User{UserId: req.User.Id, Firstname: req.User.FirstName, FamilyName: req.User.FamilyName})

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
