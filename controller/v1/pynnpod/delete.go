package pynnpod

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

type DeletePodRequest struct {
	Username string `json:"username"   validate:"required"`
}

type DeletePodResponseData struct{}

func Delete(ctx iris.Context) {

	var req DeletePodRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 数据库中有无记录
	if exist, err := service.CheckPodExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist == false {
		common.ParamErrorResponse(ctx, "POD_EXIST_ERROR")
		return
	}

	// 存在就删除
	// 根据名称获取对应的user
	pod := service.GetPodByUsername(req.Username)

	// 服务器上删除
	if service.DeletPodFromServer(pod.Uid.Hex()) == false {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// 数据库上删除
	if err := service.DeletePodFromDB(pod.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	common.SuccessResponse(ctx, DeletePodResponseData{})
}
