package interactors

import (
	"context"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/errors"
	"eh-backend-api/domain/models"
)

type UserInteractor struct {
	Repository ports.UserRepository
}

func (it *UserInteractor) AddUser(ctx context.Context, user models.User) error {

	userName := user.UserId
	u, err := it.Repository.FetchByUserId(ctx, userName)

	if err != nil {
		return err
	}

	if u != nil {
		return &errors.AlreadyExistsError{Sources: string(user.UserId)}
	}

	err = it.Repository.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (it *UserInteractor) GetUser(ctx context.Context, userId models.UserName) (*models.User, error) {
	user, err := it.Repository.FetchByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.NotFoundError{Sources: string(userId)}
	}

	return user, nil
}

// di
func NewUserInputPort(repository ports.UserRepository) ports.UserInputPort {
	return &UserInteractor{
		Repository: repository,
	}
}
