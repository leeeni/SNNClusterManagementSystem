package userdir

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

type DeleteUserDirRequest struct {
	Username string `json:"username"   validate:"required"`
}

type DeleteUserDirResponseData struct{}

func Delete(ctx iris.Context) {

	var req DeleteUserDirRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 数据库中有无记录
	if exist, err := service.CheckUserDirExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if !exist {
		common.ParamErrorResponse(ctx, "USERDIR_EXIST_ERROR")
		return
	}

	// 存在就删除
	// 根据名称获取对应的user
	userdir := service.GetUserDirByUsername(req.Username)

	// 服务器上删除
	if !service.DeleteUserDirFromServer(userdir.Uid.Hex()) {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// 数据库上删除
	if err := service.DeleteUserDirFromDB(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	common.SuccessResponse(ctx, DeleteUserDirResponseData{})
}
