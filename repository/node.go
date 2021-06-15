package repository

import (
	"snns_srv/db"

	"gopkg.in/mgo.v2/bson"
)

//Node :
type Node struct {
	// Uid: Primary key (_id)
	UID      bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	IP       string        `bson:"ip"`
	Describe string        `bson:"describe"`
	Cores    int           `bson:"cores"`

	// CreatedTime and LastLogin use timestamp.
	CreatedTime int64 `bson:"createdate"`
	LastRun     int64 `bson:"lastrun"`
	IsOnline    bool  `bson:"isonline"`
}

//SetLived :
func (node *Node) SetLived(lived bool) bool {
	node.IsOnline = lived
	return node.IsOnline
}

//Checklived :
func (node *Node) Checklived() bool {
	return node.IsOnline
}

//InsertNode :
func InsertNode(node *Node) error {
	return Insert(db.NodeCollection, node)
}

//DelNode :
func DelNode(id string) error {
	return DelByID(db.NodeCollection, id)
}

//UpdateNode :
func UpdateNode(node *Node) error {
	println("Enter UpdateNode@repository:")
	print(node.UID.Hex())
	selector := bson.M{"_id": node.UID}
	data := bson.M{"$set": bson.M{"describe": node.Describe}}
	return db.NodeCollection.Update(selector, data)
}

//GetNodeAll :
// func GetNodeAll(pageindex int ,pagerows int) ([]Node, int,error) {
func GetNodeAll() ([]Node, error) {
	var nodes []Node
	// err := NodeCollection.Find(nil).Limit(pagerows).All(&nodes)
	err := db.NodeCollection.Find(nil).All(&nodes)
	return nodes, err
}

//GetNodeByName :
func GetNodeByName(name string) Node {
	node := Node{}
	GetByQ(db.NodeCollection, bson.M{"name": name}, &node)
	return node
}

//GetNodeByUID :
func GetNodeByUID(uid bson.ObjectId) Node {
	node := Node{}
	GetByQ(db.NodeCollection, bson.M{"_id": uid}, &node)
	return node
}
