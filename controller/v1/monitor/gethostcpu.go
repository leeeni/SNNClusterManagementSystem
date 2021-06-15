package monitor

import (
	"fmt"
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// GetHostCPUResponseData :
type GetHostCPUResponseData struct{}

// GetHostCPU :
func GetHostCPU(ctx iris.Context) {

	//var req SelectNodeRequest
	var cpu repository.CPU
	fmt.Printf(ctx.Params().Get("hostname"))
	cpu, err := service.GetCPUByHost(ctx.Params().Get("hostname"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, cpu)
}
