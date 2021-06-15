package timelist

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// InsertTimelistRequest -
type InsertTimelistRequest struct {
	UID      bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Rdate    string        `bson:"rdate"`
	Rhourse  string        `bson:"rhourse"`
	Notes    int           `bson:"notes"`
	Userid   string        `bson:"userid"`
	Username string        `bson:"username"`
	State    int           `bson:"state"`
}

// InsertTimelistResponseData -
type InsertTimelistResponseData struct{}

// InsertTimelist -
func InsertTimelist(ctx iris.Context) {

	var req InsertTimelistRequest
	var Timelist repository.Timelist
	// Get Request
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
	// Prepare Timelist
	req.UID = bson.NewObjectId()
	println("insert Timelist:" + req.UID.Hex())
	Timelist = repository.Timelist{
		UID:      req.UID,
		Name:     req.Name,
		Rdate:    req.Rdate,
		Rhourse:  req.Rhourse,
		Notes:    req.Notes,
		Userid:   req.Userid,
		Username: req.Username,
		State:    req.State,
	}
	if err := service.InsertTimelist(&Timelist); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, Timelist)
}
