package mysql

import (
	"eh-backend-api/conf/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

func NewDbConnection() (database *gorm.DB, err error) {
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.Mysql.DbUserName,
		config.Config.Mysql.DbUserPassword,
		config.Config.Mysql.DbHost,
		config.Config.Mysql.DbPort,
		config.Config.Mysql.DbName,
	)

	db, err := gorm.Open(config.Config.Mysql.DbDriverName, dbConnectInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
