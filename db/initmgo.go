package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

// Session -
var Session *mgo.Session

// SomeTables - Mgo Tables for Snns Service
var (
	UserCollection     *mgo.Collection
	NodeCollection     *mgo.Collection
	TaskCollection     *mgo.Collection
	TimelistCollection *mgo.Collection

	PyNNPodCollection   *mgo.Collection
	UserDirCollection   *mgo.Collection
	PortCountCollection *mgo.Collection
)

// ConnectMgo -
func ConnectMgo() error {
	// Mongodb url should be like:
	// mongodb://username:password@addr:port/dbname?authSource=admin
	url := os.Getenv("MGO_URL")
	dbname := os.Getenv("MGO_DB")
	fmt.Print(url)
	fmt.Print(dbname)

	Session, err := mgo.Dial(url)

	if err == nil {
		Session.SetMode(mgo.Monotonic, true)
		UserCollection = Session.DB(dbname).C("user")
		NodeCollection = Session.DB(dbname).C("node")
		TaskCollection = Session.DB(dbname).C("task")
		TimelistCollection = Session.DB(dbname).C("timelist")

		// Pod表
		PyNNPodCollection = Session.DB(dbname).C("pynnpod")
		// userdir表
		UserDirCollection = Session.DB(dbname).C("userdir")
		PortCountCollection = Session.DB(dbname).C("portcount")
	} else {
		return err
	}
	return nil
}
