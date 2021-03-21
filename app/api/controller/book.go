package controller

import (
	"fmt"
)

type BookController struct {
}

func (b *BookController) Get() string {
	fmt.Print("helle world")
	return "helle world root /"
}

func (b *BookController) GetShow() string {
	fmt.Print("helle world")
	// service.CreateBooks()
	return "helle world"
}

func (b *BookController) GetEcho(a string) string {
	fmt.Print("helle world")
	return "helle world" + a
}
