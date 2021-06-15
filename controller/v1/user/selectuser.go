package user

import (
	"fmt"
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// SelectUserRequest :
// type SelectUserRequest struct {
// 	pageindex int `bson:"pageindex"`
// 	pagerows  int `bson:"pagerows"`
// }

// SelectUserResponseData :
type SelectUserResponseData struct{}

// SelectUser :
func SelectUser(ctx iris.Context) {

	var users []repository.User

	users, err := service.SelectUser()
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	fmt.Print(users)
	common.SuccessResponse(ctx, users)
}
