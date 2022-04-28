package mysql

import (
	"context"
	"eh-backend-api/adapter/gateways/entities"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"

	"crypto/sha256"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AuthGateway struct {
	db *gorm.DB
}

// return false on error.
func (it *AuthGateway) Has(
	ctx context.Context,
	userName models.UserName,
	password models.Password,
) (bool, error) {

	pw := []byte(password)
	sha256 := sha256.Sum256(pw)

	result := []*entities.Password{}
	error := it.db.Where("user_id = ? AND password = ?", userName, sha256).Find(&result).Error

	if error != nil {
		return false, error
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func NewAnthRepository() ports.AuthRepository {
	db, err := NewDbConnection()
	if err != nil {
		panic(err.Error())
	}
	return &AuthGateway{db}
}
