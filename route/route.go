package route

import (
	"SNNClusterManagementSystem/controller/v2/pynnpod"
	"SNNClusterManagementSystem/controller/v2/user"
	"SNNClusterManagementSystem/controller/v2/usercenter"
	"SNNClusterManagementSystem/controller/v2/userdir"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {

	api := app.Party("/").AllowMethods()
	{
		// v2版本：PyNN镜像接口
		v2 := api.Party("/api/v2")
		{
			// pod管理
			v2.PartyFunc("/pynnpod", func(p iris.Party) {
				p.Post("/update", pynnpod.Update)
				p.Post("/build", pynnpod.Build)
				p.Get("/get", pynnpod.Get)
				p.Delete("/delete", pynnpod.Delete)
			})

			// userdir管理
			v2.PartyFunc("/userdir", func(p iris.Party) {
				p.Post("/build", userdir.Build)
				p.Get("/get", userdir.Get)
				p.Delete("/delete", userdir.Delete)
			})

			// user管理
			v2.PartyFunc("/user", func(p iris.Party) {
				p.Post("/register", user.Register)
				p.Post("/login", user.Login)
				p.Delete("/logout", user.Logout)
			})

			// usercenter管理
			v2.PartyFunc("/usercenter", func(p iris.Party) {
				p.Post("/token", usercenter.Token)
				p.Get("/userinfo", usercenter.UserInfo)
			})
		}
	}
}
