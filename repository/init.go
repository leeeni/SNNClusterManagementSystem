package repository

import (
	"os"

	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

var (
	UserCollection *mgo.Collection
	PyNNPodCollection *mgo.Collection
	UserDirCollection *mgo.Collection
)

func ConnectMgo() error {
	// Mongodb url should be like:
	// mongodb://username:password@addr:port/dbname?authSource=admin
	url := os.Getenv("MGO_URL")
	dbname := os.Getenv("MGO_DB")

	Session, err := mgo.Dial(url)

	if err != nil {
		return err
	} else {
		Session.SetMode(mgo.Monotonic, true)
		// user表
		UserCollection = Session.DB(dbname).C("user")
		// Pod表
		PyNNPodCollection = Session.DB(dbname).C("pynnpod")
		// userdir表
		UserDirCollection = Session.DB(dbname).C("userdir")
	}
	return nil
}
