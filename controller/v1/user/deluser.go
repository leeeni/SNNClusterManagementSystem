package user

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// DelaskResponseData :
type DelaskResponseData struct{}

// DelUser :
func DelUser(ctx iris.Context) {
	print(ctx.Params().GetString("id"))
	err := service.DelUser(ctx.Params().GetString("id"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	//fmt.Print(users)
	common.SuccessResponse(ctx, "")
}
