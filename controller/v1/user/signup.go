package user

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// SignupRequest :
type SignupRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

// SignupResponseData :
type SignupResponseData struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	RoleSys   bool   `json:"rolesys"`
	RoleAdmin bool   `json:"roleadmin"`
}

// Signup :
func Signup(ctx iris.Context) {
	var req SignupRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	if req.Username == "root" || req.Username == "admin" {
		common.ParamErrorResponse(ctx, "该用户已存在")
		return
	}

	if exist, err := service.CheckUserExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist {
		common.ParamErrorResponse(ctx, "该用户已存在")
		return
	}

	// Prepare user
	user := repository.User{
		UID:         bson.NewObjectId(),
		Verified:    false,
		Username:    strings.ToLower(req.Username),
		CreatedTime: time.Now().Unix(),
		LastLogin:   time.Now().Unix(),
		RoleSys:     false,
		RoleAdmin:   false,
		State:       0,
		IsPclUser:   false,
	}

	// 设置hash密码
	if user.SetPassword(req.Password) != nil {
		common.EncryptErrorResponse(ctx)
		return
	}

	// 1、数据库插入用户信息
	if err := service.InsertUser(&user); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// 2、创建用户目录
	service.BuildUserDir(user.UID.Hex())
	// 保存userdir相关信息到数据库
	path := "/home/work/ClientDir/" + user.UID.Hex()
	userdir := repository.UserDir{
		Uid:         user.UID,
		UserName:    strings.ToLower(user.Username),
		DirName:     strings.ToLower(user.UID.Hex()),
		DirPath:     strings.ToLower(path),
		CreatedTime: time.Now().Unix(),
		LastUse:     time.Now().Unix(),
		Role:        repository.NormalUserDir,
		IsBanned:    false,
	}
	if err := service.InsertUserDir(&userdir); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// 3、创建pod
	// 获取现在使用的HostPort
	PortInfo := repository.GetPortInfo()
	NowUsePort := PortInfo.NowUsePort
	HostPort := NowUsePort + 1
	_ = service.BuildPod(user.UID.Hex(), HostPort)
	// 获取现在运行的pod
	RunningPod := service.GetPodByUsernameFromServer(user.UID.Hex())
	podcpu := RunningPod.Spec.Containers[0].Resources.Requests.Cpu().String()
	podmemory := RunningPod.Spec.Containers[0].Resources.Requests.Memory().String()
	// 修改数据库端口信息
	PortInfo.NowUsePort = HostPort
	_ = repository.UpdatePortInfo(PortInfo)
	pynnpod := repository.PynnPod{
		Uid:         user.UID,
		Username:    strings.ToLower(user.Username),
		PodName:     strings.ToLower("pynn-pod-" + user.UID.Hex()),
		PodIp:       strings.ToLower(RunningPod.Status.PodIP),
		PodStatus:   strings.ToLower(string(RunningPod.Status.Phase)),
		HostPort:    HostPort,
		PodCpu:      podcpu,
		PodMemory:   podmemory,
		CreatedTime: time.Now().Unix(),
		LastUse:     time.Now().Unix(),
		Role:        repository.NormalUserDir,
		IsBanned:    false,
	}

	// 插入数据库
	if err := service.InsertPod(&pynnpod); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	// user = service.GetUserByAccount(req.Username)
	// if !user.CheckPassword(req.Password) {
	// 	common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
	// 	return
	// }

	//service.SetLoginSession(ctx, user.UID)
	common.SuccessResponse(ctx, SignupResponseData{
		UID:       user.UID.Hex(),
		Username:  user.Username,
		RoleSys:   user.RoleSys,
		RoleAdmin: user.RoleAdmin,
	})
}
