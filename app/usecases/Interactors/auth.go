package interactors

import (
	"context"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/errors"
	"eh-backend-api/domain/models"
	"log"

	"github.com/google/uuid"
)

type AuthInteractor struct {
	Repository ports.AuthRepository
}

func (it *AuthInteractor) AhtuAndCreateToken(
	ctx context.Context,
	userName models.UserName,
	password models.Password,
) (*models.SessionToken, error) {
	success, err := it.Repository.Has(ctx, userName, password)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, &errors.SystemError{Message: err.Error()}
	}

	if success {
		uuid, _ := uuid.NewRandom()
		var tkn models.SessionToken = models.SessionToken(uuid.String())
		return &tkn, nil
	}

	return nil, &errors.AuthenticationError{UserName: userName}

}

func NewAuthIputPort(repository ports.AuthRepository) ports.AuthInputPort {
	return &AuthInteractor{
		Repository: repository,
	}
}
