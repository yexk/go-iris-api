package model

import (
	"server/app/common"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Name     string
	Category uint8
	Tags     string
}

func _init() (err error) {
	// create table
	err = common.GetDB().AutoMigrate(&Books{})
	return err
}

func CreateData(book *Books) (db *gorm.DB, err error) {
	_init()

	db = common.GetDB().Create(&book)
	return
}
