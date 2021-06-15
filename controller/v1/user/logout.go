package user

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// LogoutRequest -
type LogoutRequest struct{}

// LogoutResponseData -
type LogoutResponseData struct{}

// Logout -
func Logout(ctx iris.Context) {
	service.SetLogoutSession(ctx)
	common.SuccessResponse(ctx, LogoutResponseData{})
}
