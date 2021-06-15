package timelist

import (
	"snns_srv/controller/v1/common"
	"snns_srv/db"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// UpdateTimelistRequest -
type UpdateTimelistRequest struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Rdate    string `json:"rdate"`
	Rhourse  string `json:"rhourse"`
	Notes    int    `json:"notes"`
	Userid   string `json:"userid"`
	Username string `json:"username"`
	State    int    `json:"state"`
}

// UpdateTimelistResponseData -
type UpdateTimelistResponseData struct{}

// UpdateTimelist -
func UpdateTimelist(ctx iris.Context) {

	var req UpdateTimelistRequest
	// Get Request
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	objID := bson.ObjectIdHex(req.UID)
	print(objID.String())
	selector := bson.M{"_id": objID}
	data := bson.M{"$set": bson.M{"name": req.Name, "rdate": req.Rdate, "rhourse": req.Rhourse, "notes": req.Notes, "userid": req.Userid, "username": req.Username, "state": req.State}}

	if err := db.TimelistCollection.Update(selector, data); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, UpdateTimelistResponseData{})
}
