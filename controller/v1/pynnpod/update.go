package pynnpod

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

type UpdatePodRequest struct {
	Username  string `json:"username"    validate:"required"`
	PodCpu    int64  `json:"podcpu"      validate:"required"`
	PodMemory int64  `json:"podmemory"   validate:"required"`
}

type UpdatePodResponseData struct{}

func Update(ctx iris.Context) {

	var req UpdatePodRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// **************** 先删除，后创建 *******************
	// 数据库中有无记录
	if exist, err := service.CheckPodExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if !exist {
		common.ParamErrorResponse(ctx, "POD_Update_ERROR")
		return
	}

	// 根据名称获取对应的user
	pod := service.GetPodByUsername(req.Username)

	// 服务器上删除
	if !service.DeletPodFromServer(pod.Uid.Hex()) {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// 数据库上删除
	if err := service.DeletePodFromDB(pod.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// 创建
	// 获取数据库中对应的user对象
	user := service.GetUserByAccount(req.Username)
	// 对应user的uid
	uid := user.UID.Hex()

	// 创建Pod，service包含了检查是否存在
	if !service.BuildPodByCpuAndMemoryFromServer(uid, req.PodCpu, req.PodMemory) {
		return
	}

	// 获取现在运行的pod
	RunningPod := service.GetPodByUsernameFromServer(uid)
	podcpu := RunningPod.Spec.Containers[0].Resources.Requests.Cpu().String()
	podmemory := RunningPod.Spec.Containers[0].Resources.Requests.Memory().String()

	// Prepare userdir
	pynnpod := repository.PynnPod{
		Uid:         user.UID,
		Username:    strings.ToLower(user.Username),
		PodName:     strings.ToLower("pynn-pod-" + uid),
		PodIp:       strings.ToLower(RunningPod.Status.PodIP),
		PodStatus:   strings.ToLower(string(RunningPod.Status.Phase)),
		PodCpu:      strings.ToLower(podcpu),
		PodMemory:   strings.ToLower(podmemory),
		CreatedTime: time.Now().Unix(),
		LastUse:     time.Now().Unix(),
		Role:        repository.NormalUserDir,
		IsBanned:    false,
	}

	// 查看数据库中是否有对应pod对象
	if exist, err := service.CheckPodExistByUsername(user.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist {
		common.ParamErrorResponse(ctx, "POD_Update_ERROR")
		return
	}

	// 插入数据库
	if err := service.InsertPod(&pynnpod); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	common.SuccessResponse(ctx, UpdatePodResponseData{})
}
