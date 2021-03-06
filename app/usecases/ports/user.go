package ports

import (
	"context"
	"eh-backend-api/domain/models"
)

// receiving user data
type UserInputPort interface {
	AddUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, userId models.UserName) (*models.User, error)
}

// CRUD user data to something
type UserRepository interface {
	AddUser(ctx context.Context, user models.User) error
	FetchByUserId(ctx context.Context, userId models.UserName) (*models.User, error)
}
