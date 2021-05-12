package user

import (
	"SNNClusterManagementSystem/controller/v2/common"
	"SNNClusterManagementSystem/service"

	"github.com/kataras/iris/v12"
)

type LoginRequest struct {
	Account  string `json:"account"  validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseData struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

func Login(ctx iris.Context) {

	var req LoginRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 根据账号获取数据库中用户对象
	user := service.GetUserByAccount(req.Account)
	// 查看请求的密码和数据库中的密码是否一样
	if !user.CheckPassword(req.Password) {
		common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
		return
	}

	// 一样的话创建登陆会话
	service.SetLoginSession(ctx, user.Uid)
	common.SuccessResponse(ctx, LoginResponseData{
		Uid:      user.Uid.Hex(),
		Username: user.Username,
		Role:     user.Role,
	})
}
