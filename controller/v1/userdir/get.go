package userdir

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

type GetUserDirRequest struct {
	Username string `json:"username"   validate:"required"`
}

type GetUserDirResponseData struct {
	Uid         string `json:"uid"`
	UserName    string `json:"username"`
	DirName     string `json:"dirname"`
	DirPath     string `json:"dirpath"`
	CreatedTime int64  `json:"created_time"`
}

func Get(ctx iris.Context) {

	var req GetUserDirRequest

	// json格式检验
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	// 数据库中有无记录
	if exist, err := service.CheckUserDirExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if !exist {
		common.ParamErrorResponse(ctx, "USERDIR_EXIST_ERROR")
		return
	}

	// 存在就获取
	userdir := service.GetUserDirByUsername(req.Username)

	response := GetUserDirResponseData{
		Uid:         userdir.Uid.Hex(),
		UserName:    userdir.UserName,
		DirName:     userdir.DirName,
		DirPath:     userdir.DirPath,
		CreatedTime: userdir.CreatedTime,
	}

	common.SuccessResponse(ctx, response)
}
