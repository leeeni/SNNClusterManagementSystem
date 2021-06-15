package usercenter

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

type TokenRequest struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

func Token(ctx iris.Context) {

	var req TokenRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	// 向用户中心发送用户验证
	responData, err := service.GetToken(req.Username, req.Password)

	if err != nil || responData.AccessToken == "" {
		common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
		return
	}

	// 成功了返回信息
	common.SuccessResponse(ctx, responData)

}
