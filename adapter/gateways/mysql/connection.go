package mysql

import (
	"eh-backend-api/conf/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// gorm
// SEE:https://gorm.io/ja_JP/docs/index.html
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
		log.Fatalln(err.Error())
	}
	return db, nil
}
