package task

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// InsertTaskRequest -
type InsertTaskRequest struct {
	TaskID   string `bson:"taskid"`
	Name     string `bson:"name"`
	Describe string `bson:"describe"`
	User     string `bson:"user"`
}

// InsertTaskResponseData -
type InsertTaskResponseData struct{}

// InsertTask -
func InsertTask(ctx iris.Context) {

	var req InsertTaskRequest
	var task repository.Task
	// Get Request
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	// Prepare task
	task = repository.Task{
		TaskID:   req.TaskID,
		Name:     req.Name,
		Describe: req.Describe,
		User:     req.User,
	}
	if err := service.InsertTask(&task); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, InsertTaskResponseData{})
}
