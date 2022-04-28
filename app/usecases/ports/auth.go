package ports

import (
	"context"
	"eh-backend-api/domain/models"
)

type AuthInputPort interface {
	AhtuAndCreateToken(ctx context.Context, userName models.UserName, password models.Password) (*models.SessionToken, error)
}

type AuthRepository interface {
	Has(ctx context.Context, userName models.UserName, password models.Password) (bool, error)
}
