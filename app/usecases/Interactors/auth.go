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
	authRepository  ports.AuthRepository
	userRepository  ports.UserRepository
	tokenRepositroy ports.TokenRepository
}

func (it *AuthInteractor) UpdatePassword(
	ctx context.Context,
	userName models.UserName,
	before models.Password,
	after models.Password,
) error {

	success, err := it.authRepository.HasPassword(ctx, userName, before)

	if err != nil {
		log.Println(err)
		return &errors.SystemError{Message: err.Error()}
	}

	if !success {
		return &errors.AuthenticationError{Sources: string(userName)}
	}
	if err := it.authRepository.Update(ctx, userName, after); err != nil {
		return &errors.SystemError{Message: err.Error()}
	}

	return nil

}

func (it *AuthInteractor) AhtuAndCreateToken(
	ctx context.Context,
	userName models.UserName,
	password models.Password,
) (*models.SessionToken, error) {
	success, err := it.authRepository.HasPassword(ctx, userName, password)

	if err != nil {
		log.Println(err)
		return nil, &errors.SystemError{Message: err.Error()}
	}

	if !success {
		return nil, &errors.AuthenticationError{Sources: string(userName)}
	}

	uuid, _ := uuid.NewRandom()
	tkn := models.SessionToken(uuid.String())

	tknerr := it.tokenRepositroy.Insert(ctx, tkn, userName)
	if tknerr != nil {
		log.Println(tknerr)
		return nil, &errors.SystemError{Message: err.Error()}
	}

	return &tkn, nil

}

func (it *AuthInteractor) GetUserRole(
	ctx context.Context,
	sessionToken models.SessionToken,
) (*models.UserRole, error) {
	un, err := it.tokenRepositroy.Fetch(ctx, sessionToken)

	if err != nil {
		return nil, &errors.SystemError{Message: err.Error()}
	}

	if un == nil {
		return nil, &errors.AuthenticationError{Sources: string(sessionToken)}
	}

	roles, err := it.authRepository.FetchRoles(ctx, *un)

	if err != nil {
		return nil, &errors.SystemError{Message: err.Error()}
	}

	usrl := models.UserRole{
		UserName: *un,
		Roles:    roles,
	}

	return &usrl, nil
}

// di
func NewAuthIputPort(
	authRepository ports.AuthRepository,
	userRepository ports.UserRepository,
	tokenRepository ports.TokenRepository,
) ports.AuthInputPort {
	return &AuthInteractor{
		authRepository:  authRepository,
		userRepository:  userRepository,
		tokenRepositroy: tokenRepository,
	}
}
