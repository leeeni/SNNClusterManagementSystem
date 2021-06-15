package repository

import (
	"snns_srv/db"

	"gopkg.in/mgo.v2/bson"
)

//Port :端口数据
type PortInfo struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Tag        string        `bson:"tag"`
	StartPort  int32         `bson:"start_port"`
	NowUsePort int32         `bson:"now_use_port"`
	endPort    int32         `bson:"end_port"`
}

// 查
func GetPortInfo() PortInfo {
	portInfo := PortInfo{}
	GetByQ(db.PortCountCollection, bson.M{"tag": "tag"}, &portInfo)
	return portInfo
}

// 改
func UpdatePortInfo(portInfo PortInfo) error {
	err := db.PortCountCollection.Update(bson.M{"tag": "tag"}, portInfo)
	return err
}
