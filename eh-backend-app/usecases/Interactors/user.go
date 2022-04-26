package interactors

import (
	"app/usecases/ports"
	"context"
	"domain/models"
)

type UserInteractor struct {
	Repository ports.UserRepository
}

func NewUserInputPort(repository ports.UserRepository) ports.UserInputPort {
	return &UserInteractor{
		Repository: repository,
	}
}

func (u *UserInteractor) AddUser(ctx context.Context, user *models.User) error {

	err := u.Repository.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserInteractor) GetUser(ctx context.Context, userId string) (*models.User, error) {
	user, err := u.Repository.FetchByUserId(ctx, userId)
	if err != nil {
		return user, err
	}

	return user, nil
}
