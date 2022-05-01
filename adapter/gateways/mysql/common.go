package mysql

import (
	"eh-backend-api/conf/config"
	"fmt"
	"log"

	"time"

	"github.com/jinzhu/gorm"
)

// gorm
// SEE:https://gorm.io/ja_JP/docs/index.html
func NewDbConnection() (*gorm.DB, error) {
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

func toSqlString(t time.Time) string {
	return t.Format("2006-01-02")
}

func ToDate(src string) time.Time {
	dt, _ := time.Parse(time.RFC3339, src)
	return dt
}
