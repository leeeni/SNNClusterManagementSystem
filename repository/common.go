package repository

import (
	"snns_srv/util/log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ErrHandle -
func ErrHandle(err error) error {
	if err == nil || err.Error() == "not found" {
		return nil
	}
	log.Logger.Errorf("Database error: %v", err)
	return err
}

// Insert -
// ---------------------------------------------------------------------
func Insert(collection *mgo.Collection, i interface{}) error {
	err := collection.Insert(i)
	return ErrHandle(err)
}

// Del -
// ---------------------------------------------------------------------
func DelByID(collection *mgo.Collection, uid string) error {
	err := collection.Remove(bson.M{"_id": bson.ObjectIdHex(uid)})
	return ErrHandle(err)
}

// QueryALL -
// ---------------------------------------------------------------------
func QueryALL(collection *mgo.Collection, q interface{}, i interface{}) {
	cnt, _ := collection.Find(q).Count()
	print("\nCount%d", cnt)
	_ = collection.Find(q).All(&i)
	print("Enter QueryALL ! \n")
	// for index, element := range i {
	// 	print("---------------")
	// 	print(index)
	// 	print(element.(Node).Describe)
	// }
}

// GetByQ - Get one document by query
func GetByQ(collection *mgo.Collection, q interface{}, i interface{}) {
	_ = collection.Find(q).One(i)
}

// Count -
func Count(collection *mgo.Collection, q interface{}) (int, error) {
	cnt, err := collection.Find(q).Count()
	err = ErrHandle(err)
	return cnt, err
}

// Has -
func Has(collection *mgo.Collection, q interface{}) (bool, error) {
	cnt, err := Count(collection, q)
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}
