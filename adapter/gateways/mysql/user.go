package mysql

import (
	"context"
	"eh-backend-api/adapter/gateways/entities"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserGateway struct {
	db *gorm.DB
}

func NewUserRepository() ports.UserRepository {
	db, err := NewDbConnection()
	if err != nil {
		panic(err.Error())
	}
	return &UserGateway{db}
}

func (it *UserGateway) AddUser(ctx context.Context, user *models.User) error {

	error := it.db.Create(&entities.User{
		UserId:     user.UserId,
		FirstName:  user.Firstname,
		FamilyName: user.FamilyName,
	}).Error

	if error != nil {
		return error
	}

	return nil
}

func (it *UserGateway) FetchByUserId(ctx context.Context, userId string) (*models.User, error) {

	result := []*entities.User{}
	error := it.db.Where("user_id = ?", userId).Find(&result).Error

	if error != nil {
		return nil, error
	}

	if len(result) == 0 {
		return nil, nil
	}

	entity := result[0]

	model := new(models.User)
	model.Set(entity.UserId, entity.FirstName, entity.FamilyName)
	return model, nil

}
