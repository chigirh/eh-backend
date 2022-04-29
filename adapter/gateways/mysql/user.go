package mysql

import (
	"context"
	"crypto/sha256"
	"eh-backend-api/adapter/gateways/entities"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserGateway struct {
	db *gorm.DB
}

func (it *UserGateway) AddUser(ctx context.Context, user models.User) error {

	tx := it.db.Begin()

	// users
	if err := tx.Create(&entities.User{
		UserId:     string(user.UserId),
		FirstName:  user.Firstname,
		FamilyName: user.FamilyName,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// passwords
	pw := []byte(string(user.Password))
	sha256 := sha256.Sum256(pw)

	if err := tx.Create(&entities.Password{
		UserId:   string(user.UserId),
		Password: fmt.Sprintf("%x", sha256),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// roles
	for i := 0; i < len(user.Roles); i++ {
		if err := tx.Create(&entities.Role{UserId: string(user.UserId), Role: string(user.Roles[i])}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (it *UserGateway) FetchByUserId(ctx context.Context, userId models.UserName) (*models.User, error) {

	result := []*entities.User{}
	error := it.db.Where("user_id = ?", userId).Find(&result).Error

	if error != nil {
		return nil, error
	}

	if len(result) == 0 {
		return nil, nil
	}

	entity := result[0]

	model := models.User{
		UserId:     models.UserName(entity.UserId),
		Firstname:  entity.FirstName,
		FamilyName: entity.FamilyName,
	}
	return &model, nil

}

// di
func NewUserRepository() ports.UserRepository {
	db, err := NewDbConnection()
	if err != nil {
		panic(err.Error())
	}
	return &UserGateway{db}
}
