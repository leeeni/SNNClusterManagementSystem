package user

import (
	"SNNClusterManagementSystem/controller/v1/common"
	"SNNClusterManagementSystem/service"

	"github.com/kataras/iris/v12"
)

type LogoutRequest struct{}

type LogoutResponseData struct{}

func Logout(ctx iris.Context) {
	service.SetLogoutSession(ctx)
	common.SuccessResponse(ctx, LogoutResponseData{})
}
