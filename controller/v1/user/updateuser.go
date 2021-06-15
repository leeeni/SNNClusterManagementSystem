package user

import (
	"snns_srv/controller/v1/common"
	"snns_srv/db"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// UpdateUserRequest -
type UpdateUserRequest struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	RoleSys   bool   `json:"rolesys"`
	RoleAdmin bool   `json:"roleadmin"`
	State     int    `json:"state"`
}

// UpdateUserResponseData -
type UpdateUserResponseData struct{}

// UpdateUser -
func UpdateUser(ctx iris.Context) {

	var req UpdateUserRequest
	// Get Request
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	objID := bson.ObjectIdHex(req.UID)
	selector := bson.M{"_id": objID}
	data := bson.M{"$set": bson.M{"username": req.Username, "rolesys": req.RoleSys, "roleadmin": req.RoleAdmin, "state": req.State}}

	if err := db.UserCollection.Update(selector, data); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, UpdateUserResponseData{})
}
