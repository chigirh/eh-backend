package redis

import (
	"context"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"

	"github.com/go-redis/redis/v8"
)

type TokenGateway struct {
	cli *redis.Client
}

func (it *TokenGateway) Insert(
	ctx context.Context,
	sessionToken models.SessionToken,
	userName models.UserName,
) error {
	err := it.cli.Set(ctx, string(sessionToken), string(userName), 0).Err()
	return err
}
func (it *TokenGateway) Fetch(
	ctx context.Context,
	sessionToken models.SessionToken,
) (*models.UserName, error) {
	ret, err := it.cli.Get(ctx, string(sessionToken)).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	un := models.UserName(ret)
	return &un, nil

}

// di
func NewTokenRepository() ports.TokenRepository {
	return &TokenGateway{
		cli: NewClient(),
	}
}
