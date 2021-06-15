package monitor

import (
	"fmt"
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// GetMDResponseData :
type GetMDResponseData struct{}

// GetMonitorData :
func GetMonitorData(ctx iris.Context) {

	//var req SelectNodeRequest
	var cpu repository.CPU

	cpu, err := service.GetCPUByHost(ctx.Params().Get("hostname"))
	k, err := ctx.Params().GetInt("id")
	fmt.Print(k)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, cpu)
}
