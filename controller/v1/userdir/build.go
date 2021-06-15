package userdir

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

type BuildUserDirRequest struct {
	Username string `json:"username"   validate:"required"`
}

type BuildUserDirResponseData struct{}

func Build(ctx iris.Context) {

	var req BuildUserDirRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 数据库中是否已经存在目录
	if exist, err := service.CheckUserDirExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist {
		common.ParamErrorResponse(ctx, "USERDIR_IN_USED")
		return
	}

	// 不存在就创建目录,失败则返回
	// 根据名称获取对应的user
	user := service.GetUserByAccount(req.Username)
	// 用user的ID当作文件夹的名字
	if !service.BuildUserDir(user.UID.Hex()) {
		common.ParamErrorResponse(ctx, "USERDIR_IN_USED")
		return
	}

	// 保存userdir相关信息到数据库
	userdir := repository.UserDir{
		Uid:         user.UID,
		UserName:    user.Username,
		DirName:     strings.ToLower(user.UID.Hex()),
		DirPath:     strings.ToLower("/home/work/ClientDir/" + user.UID.Hex()),
		CreatedTime: time.Now().Unix(),
		LastUse:     time.Now().Unix(),
		Role:        repository.NormalUserDir,
		IsBanned:    false,
	}

	if err := service.InsertUserDir(&userdir); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, BuildUserDirResponseData{})
}
