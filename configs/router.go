package configs

import (
	"server/app/api/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Router 路由定义
func Router(app *iris.Application) {
	mvc.New(app.Party("/")).Handle(new(controller.IndexController))
	mvc.New(app.Party("/books")).Handle(new(controller.BookController))
}
