package monitor

import (
	"fmt"
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// GetHostMemResponseData -
type GetHostMemResponseData struct{}

// GetHostMem -
func GetHostMem(ctx iris.Context) {

	//var req SelectNodeRequest
	var mem repository.Mem
	fmt.Printf(ctx.Params().Get("hostname"))
	mem, err := service.GetMemByHost(ctx.Params().Get("hostname"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, mem)
}
