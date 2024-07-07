package mysql

import (
	setting "GoBBS/settings"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(config *setting.MySQLConfig) (err error) {
	user := config.User
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.DbName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return
	} else {
		var sqlDB *sql.DB
		if sqlDB, err = db.DB(); err != nil {
			return
		}
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	return
}

func Close() {
	if sqlDB, err := db.DB(); err != nil {
		return
	} else {
		_ = sqlDB.Close()
	}
}
