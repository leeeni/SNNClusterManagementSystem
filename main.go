package main

import (
	"SNNClusterManagementSystem/repository"
	"SNNClusterManagementSystem/route"
	"SNNClusterManagementSystem/session"
	"SNNClusterManagementSystem/util/i18n"
	"SNNClusterManagementSystem/util/log"
	"SNNClusterManagementSystem/util/validator"

	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)



func main() {
	// 创建一个新的app
	app := iris.New()

	// Load setting from ./.env
	// 加载env文件中的配置，用os.Getenv获取
	_ = godotenv.Load()

	// 设置日志
	// Log Levels => https://github.com/kataras/golog/blob/master/README.md#log-levels
	log.Logger = app.Logger()
	log.Logger.SetLevel(os.Getenv("IRIS_MODE"))

	// 初始化redis和mongodb
	session.ConnectRedis()
	if err := repository.ConnectMgo(); err != nil {
		log.Logger.Errorf("Connect to mongodb failed: %s", err)
		panic(err)
	}

	// 初始化校验器
	app.Validator = validator.NewValidator()

	// 初始化i18n配置文件
	app.I18n.DefaultMessageFunc = i18n.DefaultMessageFunc
	if err := app.I18n.Load("./assets/locale/*/*", "en-US", "zh-CN"); err != nil {
		log.Logger.Errorf("Load i18n failed: %s", err)
		panic(err)
	}
	app.I18n.SetDefault("en-US")

	// 初始化路由
	route.InitRouter(app)
	
	// 运行后台程序，监听端口为80
	_ = app.Run(iris.Addr(":" + os.Getenv("PORT")))
}
