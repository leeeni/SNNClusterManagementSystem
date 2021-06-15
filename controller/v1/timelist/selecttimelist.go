package timelist

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
)

// SelectTimelistRequest -
type SelectTimelistRequest struct {
	pageindex int `bson:"pageindex"`
	pagerows  int `bson:"pagerows"`
}

// SelectTimelistResponseData -
type SelectTimelistResponseData struct{}

// SelectTimelist -
func SelectTimelist(ctx iris.Context) {

	//var req SelectTimelistRequest
	var nodes []repository.Timelist
	// Get Request
	//if err := ctx.ReadJSON(&req); err != nil {
	//common.FormErrorResponse(ctx, err)
	//	return
	//}

	//Get Timelists
	// nodes, i, err := service.SelectTimelist(req.pageindex, req.pagerows)

	nodes, err := service.SelectTimelist()
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, nodes)
}

// SelectTimelist -
func SelectTimelistun(ctx iris.Context) {

	//var req SelectTimelistRequest
	var nodes []repository.Timelist
	// Get Request
	//if err := ctx.ReadJSON(&req); err != nil {
	//common.FormErrorResponse(ctx, err)
	//	return
	//}

	//Get Timelists
	// nodes, i, err := service.SelectTimelist(req.pageindex, req.pagerows)

	nodes, err := service.SelectTimelistun(ctx.Params().Get("username"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, nodes)
}
