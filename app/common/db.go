package common

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	config, err := Init()
	if err != nil {
		panic(err)
	}

	db, err := connect(&config)
	if err != nil {
		panic(err)
	}
	return db
}

func connect(config *Config) (db *gorm.DB, err error) {
	if config.Mysql.Driver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			config.Mysql.Username,
			config.Mysql.Password,
			config.Mysql.Host,
			config.Mysql.Port,
			config.Mysql.Database,
			config.Mysql.Charset)

		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}))
		return
	}
	err = errors.New("db is null")
	return
}
