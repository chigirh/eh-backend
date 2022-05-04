package ports

import (
	"context"
	"eh-backend-api/domain/models"
)

type AuthInputPort interface {
	UpdatePassword(ctx context.Context, userName models.UserName, before models.Password, after models.Password) error
	AhtuAndCreateToken(ctx context.Context, userName models.UserName, password models.Password) (*models.SessionToken, *models.User, error)
	GetUserRole(ctx context.Context, sessionToken models.SessionToken) (*models.UserRole, error)
}

type AuthRepository interface {
	HasUserName(ctx context.Context, userName models.UserName) (bool, error)
	HasPassword(ctx context.Context, userName models.UserName, password models.Password) (bool, error)
	Insert(ctx context.Context, userName models.UserName, password models.Password) error
	Update(ctx context.Context, userName models.UserName, password models.Password) error
	FetchRoles(ctx context.Context, userName models.UserName) ([]models.Role, error)
}

type TokenRepository interface {
	Insert(ctx context.Context, sessionToken models.SessionToken, userName models.UserName) error
	Fetch(ctx context.Context, sessionToken models.SessionToken) (*models.UserName, error)
}
