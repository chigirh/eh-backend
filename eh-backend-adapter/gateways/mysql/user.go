package mysql

import (
	"adapter/config"
	"adapter/gateways/entities"
	"app/usecases/ports"
	"context"
	"domain/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserGateway struct {
	db *gorm.DB
}

func NewUserRepository() ports.UserRepository {
	db, err := newDbConnection()
	if err != nil {
		panic(err.Error())
	}
	return &UserGateway{db}
}

func (gateway *UserGateway) AddUser(ctx context.Context, user *models.User) error {

	error := gateway.db.Create(&entities.User{
		UserId:     user.UserId,
		FirstName:  user.Firstname,
		FamilyName: user.FamilyName,
	}).Error

	if error != nil {
		return error
	}

	return nil
}

func (gateway *UserGateway) FetchByUserId(ctx context.Context, userId string) (*models.User, error) {

	result := []*entities.User{}
	error := gateway.db.Where("user_id = ?", userId).Find(&result).Error

	defer gateway.db.Close()

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

func newDbConnection() (database *gorm.DB, err error) {
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)

	db, err := gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		fmt.Print(err)
	}
	return db, nil
}
