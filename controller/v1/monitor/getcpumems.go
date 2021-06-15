package monitor

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"


	"github.com/kataras/iris/v12"
)

// GetMDsResponseData :
type GetMDsResponseData struct{}

// GetMonitorDatas :
func GetMonitorDatas(ctx iris.Context) {

	//var req SelectNodeRequest
	
	var mds repository.MonitorDatas
	len,_ :=ctx.Params().GetInt("len");
	print(len)
	mds, err := service.GetCPUMEMs(ctx.Params().Get("hostname"), len)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, mds)
}
