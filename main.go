package main

import (
	"snns_srv/db"
	"snns_srv/route"
	"snns_srv/session"
	"snns_srv/util/i18n"
	"snns_srv/util/log"
	"snns_srv/util/validator"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	// Load setting from ./.env
	_ = godotenv.Load()

	//初始化时序数据库连接
	fmt.Print(os.Getenv("INFLUX_ADD"))
	db.Conntsdb = db.ConnInflux(os.Getenv("INFLUX_ADD"), os.Getenv("INFLUX_USERNAME"), os.Getenv("INFLUX_PASSWORD"))

	// Create app.
	app := iris.New()

	// Set logger.
	// Log Levels => https://github.com/kataras/golog/blob/master/README.md#log-levels
	log.Logger = app.Logger()
	log.Logger.SetLevel(os.Getenv("IRIS_MODE"))

	// Init redis and mongodb.
	session.ConnectRedis()
	if err := db.ConnectMgo(); err != nil {
		log.Logger.Errorf("Connect to mongodb failed: %s", err)
		panic(err)
	}

	// Init validator.
	app.Validator = validator.NewValidator()

	// Init i18n config.
	app.I18n.DefaultMessageFunc = i18n.DefaultMessageFunc
	if err := app.I18n.Load("./assets/locale/*/*", "en-US", "zh-CN"); err != nil {
		log.Logger.Errorf("Load i18n failed: %s", err)
		panic(err)
	}
	app.I18n.SetDefault("en-US")

	// Init router.
	route.InitRouter(app)

	_ = app.Run(iris.Addr(":" + os.Getenv("PORT")))
}
