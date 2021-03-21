package model

import (
	"fmt"
	"server/app/common"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Name     string
	Category uint8
	Tags     string
}

func CreateData(book *Books) {
	fmt.Print("1231")
	common.GetDB().Create(&book)
}
