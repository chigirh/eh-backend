package interactors

import (
	"context"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/errors"
	"eh-backend-api/domain/models"
	"fmt"
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
	password models.Password,
) error {

	user, err := it.userRepository.FetchByUserId(ctx, userName)

	if err != nil {
		log.Println(err)
		return &errors.SystemError{Message: err.Error()}
	}

	if user == nil {
		return &errors.NotFoundError{Sources: fmt.Sprintf("user_name=%s", string(userName))}
	}

	has, err := it.authRepository.HasUserName(ctx, userName)

	if err != nil {
		log.Println(err)
		return &errors.SystemError{Message: err.Error()}
	}

	if has {
		if err := it.authRepository.Update(ctx, userName, password); err != nil {
			return &errors.SystemError{Message: err.Error()}
		}
	} else {
		if err := it.authRepository.Insert(ctx, userName, password); err != nil {
			return &errors.SystemError{Message: err.Error()}
		}
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
