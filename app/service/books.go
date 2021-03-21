package service

import "server/app/model"

func CreateBooks() {
	model.CreateData(&model.Books{
		Name: "Yexk",
	})
}
