package mysql

import (
	"context"
	"crypto/sha256"
	"eh-backend-api/adapter/gateways/entities"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AuthGateway struct{}

// return false on error.
func (it *AuthGateway) HasUserName(
	ctx context.Context,
	userName models.UserName,
) (bool, error) {

	db, err := NewDbConnection()
	if err != nil {
		return false, err
	}

	result := []*entities.Password{}
	err = db.Where("user_id = ?", string(userName)).Find(&result).Error

	if err != nil {
		return false, err
	}

	if len(result) == 0 {
		return false, nil
	}

	db.Close()
	return true, nil
}

// return false on error.
func (it *AuthGateway) HasPassword(
	ctx context.Context,
	userName models.UserName,
	password models.Password,
) (bool, error) {

	db, err := NewDbConnection()
	if err != nil {
		return false, err
	}

	pw := []byte(string(password))
	sha256 := sha256.Sum256(pw)

	result := []*entities.Password{}
	err = db.Where("user_id = ? AND password = ?", userName, fmt.Sprintf("%x", sha256)).Find(&result).Error

	if err != nil {
		return false, err
	}

	if len(result) == 0 {
		return false, nil
	}

	db.Close()
	return true, nil
}

func (it *AuthGateway) Insert(
	ctx context.Context,
	userName models.UserName,
	password models.Password,
) error {

	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	pw := []byte(string(password))
	sha256 := sha256.Sum256(pw)

	err = tx.Create(&entities.Password{
		UserId:   string(userName),
		Password: fmt.Sprintf("%x", sha256),
	}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	db.Close()
	return nil
}

func (it *AuthGateway) Update(
	ctx context.Context,
	userName models.UserName,
	password models.Password,
) error {

	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	pw := []byte(string(password))
	sha256 := sha256.Sum256(pw)
	err = tx.Debug().Model(&entities.Password{}).
		Where("user_id = ?", string(userName)).
		Updates(&entities.Password{
			Password: fmt.Sprintf("%x", sha256),
		}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	db.Close()
	return err
}

func (it *AuthGateway) FetchRoles(
	ctx context.Context,
	userName models.UserName,
) ([]models.Role, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	result := []*entities.Role{}
	error := db.Where("user_id = ?", string(userName)).Find(&result).Error

	if error != nil {
		return nil, error
	}

	if len(result) == 0 {
		return []models.Role{}, nil
	}

	mdls := []models.Role{}
	for i := 0; i < len(result); i++ {
		switch models.Role(result[i].Role) {
		case models.RoleAadmin:
			mdls = append(mdls, models.RoleAadmin)
			break
		case models.RoleCorp:
			mdls = append(mdls, models.RoleCorp)
			break
		case models.RoleGene:
			mdls = append(mdls, models.RoleGene)
			break
		}
	}

	db.Close()
	return mdls, nil
}

// di
func NewAnthRepository() ports.AuthRepository {
	return &AuthGateway{}
}
