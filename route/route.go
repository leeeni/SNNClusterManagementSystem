package route

import (
	"snns_srv/controller/v1/monitor"
	"snns_srv/controller/v1/neuron"
	"snns_srv/controller/v1/node"
	"snns_srv/controller/v1/pcluser"
	"snns_srv/controller/v1/pynnpod"
	"snns_srv/controller/v1/task"
	"snns_srv/controller/v1/timelist"
	"snns_srv/controller/v1/user"
	"snns_srv/controller/v1/usercenter"
	"snns_srv/controller/v1/userdir"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

// InitRouter -
func InitRouter(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
	})
	app.Use(crs)
	api := app.Party("/").AllowMethods(iris.MethodOptions)
	{
		v1 := api.Party("/api/v1")
		{

			// 用户管理
			v1.PartyFunc("/user", func(p iris.Party) {
				p.Post("/signup", user.Signup)
				p.Post("/login", user.Login)
				p.Post("/setrole", user.Login)
				p.Post("/setbanned", user.Login)
				p.Get("/selectuser", user.SelectUser)
				p.Delete("/logout", user.Logout)
				p.Get("/deluser/{id:string}", user.DelUser)
				p.Post("/updateuser", user.UpdateUser)
			})

			// 节点管理
			v1.PartyFunc("/node", func(p iris.Party) {
				p.Post("/insertnode", node.InsertNode)
				p.Post("/updatenode", node.UpdateNode)
				p.Get("/selectnode", node.SelectNode)
				p.Get("/delnode/{id:string}", node.DelNode)
			})

			// 时间管理
			v1.PartyFunc("/timelist", func(p iris.Party) {
				p.Post("/inserttimelist", timelist.InsertTimelist)
				p.Post("/updatetimelist", timelist.UpdateTimelist)
				p.Get("/selecttimelist", timelist.SelectTimelist)
				p.Get("/selecttimelistun/{username:string}", timelist.SelectTimelistun)
				p.Get("/deltimelist/{id:string}", timelist.DelTimelist)
			})

			// 任务管理
			v1.PartyFunc("/task", func(p iris.Party) {
				p.Post("/inserttask", task.InsertTask)
				//p.Post("/updatetask", task.UpdateTask)
				p.Get("/selecttask/{username:string}", task.SelectTask)
				p.Get("/deltask/{id:string}", task.DelTask)
			})

			// 节点管理
			v1.PartyFunc("/monitor", func(p iris.Party) {
				p.Get("/getcpumems/{hostname:string}/{len:int}", monitor.GetMonitorDatas)
				p.Get("/gethostmem/{hostname:string}", monitor.GetHostMem)
				p.Get("/gethostcpu/{hostname:string}", monitor.GetHostCPU)
			})

			// 神经元管理
			v1.PartyFunc("/neuron", func(p iris.Party) {
				p.Get("/getspking/{gid:string}", neuron.GetNeuronSpiking)
				p.Get("/getvl/{gid:string}/{tt:int64}", neuron.GetNeuronV)
				p.Get("/getOnevl/{gid:string}/{id:int64}", neuron.GetNeuronOneV)
				p.Get("/getOnesvl/{gid:string}/{id:int64}", neuron.GetNeuronOnesV)
				p.Get("/getspking3/{gid:string}/{id:int64}", neuron.GetNeuronSpiking3)

				p.Get("/getheatmapNvlfrombin/{username:string}/{gid:string}/{tt:string}", neuron.GetHeatMapNvlFromBin)
				p.Get("/getmulneuvlgrombin/{username:string}/{gid:string}/{minId:string}/{maxId:string}", neuron.GetMulNeuVlFromBin)
				p.Get("/getspkingfrombin/{username:string}/{gid:string}/{minId:string}/{maxId:string}", neuron.GetSpkingFromBin)
			})

			// pod管理
			v1.PartyFunc("/pynnpod", func(p iris.Party) {
				p.Post("/update", pynnpod.Update)
				p.Post("/build", pynnpod.Build)
				p.Post("/get", pynnpod.Get)
				p.Delete("/delete", pynnpod.Delete)
			})

			// userdir管理
			v1.PartyFunc("/userdir", func(p iris.Party) {
				p.Post("/build", userdir.Build)
				p.Get("/get", userdir.Get)
				p.Delete("/delete", userdir.Delete)
			})

			// usercenter入口
			v1.PartyFunc("/usercenter", func(p iris.Party) {
				p.Post("/token", usercenter.Token)
				p.Get("/userinfo", usercenter.UserInfo)
			})

			// pcluser入口
			v1.PartyFunc("/pcluser", func(p iris.Party) {
				p.Post("/login", pcluser.Login)
			})
		}
	}
}
