package node

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// SelectNodeRequest -
type SelectNodeRequest struct {
	pageindex int `bson:"pageindex"`
	pagerows  int `bson:"pagerows"`
}

// SelectNodeResponseData -
type SelectNodeResponseData struct{}

// SelectNode -
func SelectNode(ctx iris.Context) {

	//var req SelectNodeRequest
	var nodes []repository.Node
	// Get Request
	//if err := ctx.ReadJSON(&req); err != nil {
	//common.FormErrorResponse(ctx, err)
	//	return
	//}

	//Get Nodes
	// nodes, i, err := service.SelectNode(req.pageindex, req.pagerows)

	nodes, err := service.SelectNode()
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, nodes)
}
