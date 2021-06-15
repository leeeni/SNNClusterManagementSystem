package node

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// InsertNodeRequest -
type InsertNodeRequest struct {
	UID      bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	IP       string        `bson:"ip"`
	Describe string        `bson:"describe"`
	Cores    int           `bson:"cores"`
}

// InsertNodeResponseData -
type InsertNodeResponseData struct{}

// InsertNode -
func InsertNode(ctx iris.Context) {

	var req InsertNodeRequest
	var Node repository.Node
	// Get Request
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	// Prepare Node
	req.UID = bson.NewObjectId()
	println("insert Node:" + req.UID.Hex())
	Node = repository.Node{
		UID:      req.UID,
		Name:     req.Name,
		IP:       req.IP,
		Describe: req.Describe,
		Cores:    req.Cores,
	}
	if err := service.InsertNode(&Node); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, Node)
}
