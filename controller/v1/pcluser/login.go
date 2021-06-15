package pcluser

import (
	"fmt"
	"snns_srv/repository"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"

	"snns_srv/controller/v1/common"
	"snns_srv/service"
)

type LoginRequest struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseData struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	RoleSys   bool   `json:"rolesys"`
	RoleAdmin bool   `json:"roleadmin"`
}

func Login(ctx iris.Context) {

	var req LoginRequest

	// 查看请求格式
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 用户中心认证获取token
	tokenRespond, err := service.GetToken(req.Username, req.Password)
	// 不通过直接返回相应错误
	if err != nil || tokenRespond.AccessToken == "" {
		common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
		return
	}

	// 到这一步了说明已经通过
	// 查看系统服务器有无用户信息（1、判断是否是第一次登陆；2、判断是否需要插入用户信息至用户数据库）
	// 根据账号获取数据库中用户对象
	ifExist, _ := service.CheckUserExistByUsername(req.Username)

	if !ifExist {
		// 如果不存，则是第一次登陆，需要添加用户信息-创建用户目录-启动pod-创建session
		// 从用户中心获取用户信息
		userInfo, _ := service.GetUserInfo(tokenRespond.TokenType + tokenRespond.AccessToken)
		fmt.Println("userInfo:", userInfo)
		// 1、数据库用户信息
		pcluser := repository.User{
			UID:         bson.NewObjectId(),
			Email:       strings.ToLower(userInfo.Email),
			Verified:    false,
			Username:    strings.ToLower(userInfo.Username),
			CreatedTime: time.Now().Unix(),
			LastLogin:   time.Now().Unix(),
			RoleSys:     true,
			RoleAdmin:   false,
			State:       0,
			IsPclUser:   true,
		}
		fmt.Println(pcluser)
		// 插入数据库
		if err := service.InsertUser(&pcluser); err != nil {
			common.DatabaseErrorResponse(ctx)
			return
		}
		// 2、创建用户目录
		service.BuildUserDir(pcluser.UID.Hex())
		// 保存userdir相关信息到数据库
		path := "/home/work/ClientDir/" + pcluser.UID.Hex()
		userdir := repository.UserDir{
			Uid:         pcluser.UID,
			UserName:    strings.ToLower(pcluser.Username),
			DirName:     strings.ToLower(pcluser.UID.Hex()),
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
		_ = service.BuildPod(pcluser.UID.Hex(), HostPort)
		// 获取现在运行的pod
		RunningPod := service.GetPodByUsernameFromServer(pcluser.UID.Hex())
		podcpu := RunningPod.Spec.Containers[0].Resources.Requests.Cpu().String()
		podmemory := RunningPod.Spec.Containers[0].Resources.Requests.Memory().String()
		// 修改数据库端口信息
		PortInfo.NowUsePort = HostPort
		_ = repository.UpdatePortInfo(PortInfo)
		pynnpod := repository.PynnPod{
			Uid:         pcluser.UID,
			Username:    strings.ToLower(pcluser.Username),
			PodName:     strings.ToLower("pynn-pod-" + pcluser.UID.Hex()),
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

		// 设置session
		service.SetLoginSession(ctx, pcluser.UID)
		// 回复响应
		common.SuccessResponse(ctx, LoginResponseData{
			UID:       pcluser.UID.Hex(),
			Username:  pcluser.Username,
			RoleSys:   pcluser.RoleSys,
			RoleAdmin: pcluser.RoleAdmin,
		})
	} else {
		// 如果存在，则非第一次登陆，直接创建session
		// 根据账号获取数据库中用户对象
		pcluser := service.GetUserByAccount(req.Username)

		// 一样的话创建登陆会话
		service.SetLoginSession(ctx, pcluser.UID)
		common.SuccessResponse(ctx, LoginResponseData{
			UID:       pcluser.UID.Hex(),
			Username:  pcluser.Username,
			RoleSys:   pcluser.RoleSys,
			RoleAdmin: pcluser.RoleAdmin,
		})
	}

}
