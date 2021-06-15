package node

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// DelNodeRequest :
type DelNodeRequest struct {
	pageindex int `bson:"pageindex"`
	pagerows  int `bson:"pagerows"`
}

// DelaskResponseData :
type DelaskResponseData struct{}

// DelNode :
func DelNode(ctx iris.Context) {
	print(ctx.Params().GetString("id"))
	err := service.DelNode(ctx.Params().GetString("id"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	//fmt.Print(users)
	common.SuccessResponse(ctx, "")
}
