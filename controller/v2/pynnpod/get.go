package pynnpod

import (
	"SNNClusterManagementSystem/controller/v2/common"
	"SNNClusterManagementSystem/service"
	"github.com/kataras/iris/v12"
)

type GetPodRequest struct {
	Username  string  `json:"username"   validate:"required"`
}

type GetPodResponseData struct {
	Uid         string        `json:"uid"`
	UserName    string		  `json:"username"`
	PodName    string        `json:"podname"`
	PodIP       string         `json:"podip"`
	PodStatus   string		 `json:"podstatus"`
	PodCpu		string			`json:"podcpu"`
	PodMemory	string			`json:"podmemory"`
}

func Get(ctx iris.Context) {

	var req GetPodRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 管理名称无法使用
	if req.Username == "root" || req.Username == "admin" {
		common.ParamErrorResponse(ctx, "USERNAME_IN_USED")
		return
	}

	// 数据库中有无记录
	if exist, err := service.CheckPodExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist == false{
		common.ParamErrorResponse(ctx, "POD_EXIST_ERROR")
		return
	}

	// 存在就数据库获取
	pod := service.GetPodByUsername(req.Username)

	// 获取响应
	response := GetPodResponseData{
		Uid: 		pod.Uid.Hex(),
		UserName:	pod.Username,
		PodName:    pod.PodName,
		PodIP:      pod.PodIp,
		PodStatus:  pod.PodStatus,
		PodCpu: 	pod.PodCpu,
		PodMemory:	pod.PodMemory,
	}

	common.SuccessResponse(ctx, response)
}