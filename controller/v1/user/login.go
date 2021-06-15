package user

import (
	"snns_srv/controller/v1/common"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// LoginRequest -
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//LoginResponseData -
type LoginResponseData struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	RoleSys   bool   `json:"rolesys"`
	RoleAdmin bool   `json:"roleadmin"`
}

// Login -
func Login(ctx iris.Context) {
	var req LoginRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	user := service.GetUserByAccount(req.Username)
	if !user.CheckPassword(req.Password) {
		common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
		return
	}
	service.SetLoginSession(ctx, user.UID)
	common.SuccessResponse(ctx, LoginResponseData{
		UID:       user.UID.Hex(),
		Username:  user.Username,
		RoleSys:   user.RoleSys,
		RoleAdmin: user.RoleAdmin,
	})
}
