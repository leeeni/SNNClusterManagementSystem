package usercenter

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

type UserInfoRequest struct {
	Authorization string `json:"authorization"`
}

func UserInfo(ctx iris.Context) {

	var req UserInfoRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 发送Get请求
	responData, err := service.GetUserInfo(req.Authorization)

	if err != nil || responData.Uid == "" {
		common.ParamErrorResponse(ctx, "USER_GET_ERROR")
		return
	}

	// 成功了返回信息
	common.SuccessResponse(ctx, responData)
}
