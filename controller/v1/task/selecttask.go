package task

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// SelectTaskRequest :
type SelectTaskRequest struct {
	pageindex int `bson:"pageindex"`
	pagerows  int `bson:"pagerows"`
}

// SelectTaskResponseData :
type SelectTaskResponseData struct{}

// SelectTask :
func SelectTask(ctx iris.Context) {
	var tasks []repository.Task
	tasks, err := service.SelectTask(ctx.Params().Get("username"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	//fmt.Print(users)
	common.SuccessResponse(ctx, tasks)
}
