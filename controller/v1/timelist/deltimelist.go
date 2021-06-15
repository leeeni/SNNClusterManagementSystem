package timelist

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// DelTimelistRequest :
type DelTimelistRequest struct {
	pageindex int `bson:"pageindex"`
	pagerows  int `bson:"pagerows"`
}

// DelaskResponseData :
type DelaskResponseData struct{}

// DelTimelist :
func DelTimelist(ctx iris.Context) {
	print(ctx.Params().GetString("id"))
	err := service.DelTimelist(ctx.Params().GetString("id"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	//fmt.Print(users)
	common.SuccessResponse(ctx, "")
}
