package node

import (
	"snns_srv/controller/v1/common"
	"snns_srv/db"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// UpdateNodeRequest -
type UpdateNodeRequest struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Describe string `json:"describe"`
	Cores    int    `json:"cores"`
}

// UpdateNodeResponseData -
type UpdateNodeResponseData struct{}

// UpdateNode -
func UpdateNode(ctx iris.Context) {

	var req UpdateNodeRequest
	// Get Request
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	objID := bson.ObjectIdHex(req.UID)
	print(objID.String())
	selector := bson.M{"_id": objID}
	data := bson.M{"$set": bson.M{"name": req.Name, "ip": req.IP, "describe": req.Describe, "cores": req.Cores}}

	if err := db.NodeCollection.Update(selector, data); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, UpdateNodeResponseData{})
}
