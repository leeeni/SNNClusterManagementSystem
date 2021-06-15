package pynnpod

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

// 初次访问注册用户目录和Pod
type BuildPodRequest struct {
	Username string `json:"username"   validate:"required"`
}

type BuildPodResponseData struct{}

func Build(ctx iris.Context) {

	var req BuildPodRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 获取数据库中对应的user对象
	user := service.GetUserByAccount(req.Username)
	// 对应user的uid
	uid := user.UID.Hex()

	// 创建Pod，service包含了检查是否存在
	if !service.BuildPod(uid, 9999) {
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
		common.ParamErrorResponse(ctx, "Pod_IN_USED")
		return
	}

	// 插入数据库
	if err := service.InsertPod(&pynnpod); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	common.SuccessResponse(ctx, BuildPodResponseData{})

}
