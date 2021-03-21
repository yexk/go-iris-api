package api

import (
	"server/configs"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func Main() {
	config, err := configs.Init()
	if err != nil {
		panic(err)
	}
	app := iris.New()
	app.Logger().SetLevel(config.AppEnv)

	app.Use(recover.New())
	app.Use(logger.New())

	configs.Router(app)
	app.Listen(config.ApiPort)
}
