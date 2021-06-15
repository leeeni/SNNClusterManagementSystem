package repository

import (
	"snns_srv/db"

	"gopkg.in/mgo.v2/bson"
)

//Timelist :
type Timelist struct {
	// Uid: Primary key (_id)
	UID      bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Rdate    string        `bson:"rdate"`
	Rhourse  string        `bson:"rhourse"`
	Notes    int           `bson:"notes"`
	Userid   string        `bson:"userid"`
	Username string        `bson:"username"`
	State    int           `bson:"state"`
}

// //SetLived :
// func (timelist *Timelist) SetLived(lived bool) bool {
// 	timelist.IsOnline = lived
// 	return timelist.IsOnline
// }

// //Checklived :
// func (timelist *Timelist) Checklived() bool {
// 	return timelist.IsOnline
// }

//InsertTimelist :
func InsertTimelist(timelist *Timelist) error {
	return Insert(db.TimelistCollection, timelist)
}

//DelTimelist :
func DelTimelist(id string) error {
	return DelByID(db.TimelistCollection, id)
}

//UpdateTimelist :
func UpdateTimelist(timelist *Timelist) error {
	println("Enter UpdateTimelist@repository:")
	print(timelist.UID.Hex())
	selector := bson.M{"_id": timelist.UID}
	data := bson.M{"$set": bson.M{"name": timelist.Name}}
	return db.TimelistCollection.Update(selector, data)
}

//GetTimelistAll :
// func GetTimelistAll(pageindex int ,pagerows int) ([]Timelist, int,error) {
func GetTimelistAll() ([]Timelist, error) {
	var nodes []Timelist
	// err := TimelistCollection.Find(nil).Limit(pagerows).All(&nodes)
	err := db.TimelistCollection.Find(nil).All(&nodes)
	return nodes, err
}

func GetTimelistAllun(username string) ([]Timelist, error) {
	var nodes []Timelist
	// err := TimelistCollection.Find(nil).Limit(pagerows).All(&nodes)
	err := db.TimelistCollection.Find(bson.M{"$or": []bson.M{bson.M{"username": username}, bson.M{"state": 0}}}).All(&nodes)
	return nodes, err
}

//GetTimelistByName :
func GetTimelistByName(name string) Timelist {
	timelist := Timelist{}
	GetByQ(db.TimelistCollection, bson.M{"name": name}, &timelist)
	return timelist
}

//GetTimelistByUID :
func GetTimelistByUID(uid bson.ObjectId) Timelist {
	timelist := Timelist{}
	GetByQ(db.TimelistCollection, bson.M{"_id": uid}, &timelist)
	return timelist
}
