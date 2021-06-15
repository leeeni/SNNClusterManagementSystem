package repository

import (
	"snns_srv/db"

	"gopkg.in/mgo.v2/bson"
)

type UserDir struct {
	// Uid: Primary key (_id)
	Uid      bson.ObjectId `bson:"_id,omitempty"`
	UserName string        `bson:"username"`
	DirName  string        `bson:"dirname"`
	DirPath  string        `bson:"dirpath"`
	// CreatedTime and LastLogin use timestamp.
	CreatedTime int64 `bson:"created_time"`
	LastUse     int64 `bson:"last_use"`
	Role        int   `bson:"role"`
	IsBanned    bool  `bson:"is_banned"`
}

// 检索
func CheckDirExistByUsername(username string) (bool, error) {
	return Has(db.UserDirCollection, bson.M{"username": username})
}

// 检索
func CheckDirExistByDirName(dirname string) (bool, error) {
	return Has(db.UserDirCollection, bson.M{"dirname": dirname})
}

// 增
func InsertUserDir(userdir *UserDir) error {
	return Insert(db.UserDirCollection, userdir)
}

// 查
func GetUserDirByUsername(username string) UserDir {
	userdir := UserDir{}
	GetByQ(db.UserDirCollection, bson.M{"username": username}, &userdir)
	return userdir
}

// 查
func GetUserDirByDirname(dirname string) UserDir {
	userdir := UserDir{}
	GetByQ(db.UserDirCollection, bson.M{"dirname": dirname}, &userdir)
	return userdir
}

// 改
func UpdateUserDir(userdir UserDir) error {
	err := db.UserCollection.Update(bson.M{"_id": userdir.Uid}, userdir)
	return err
}

// 删
func DeleteUserDirByUserName(username string) error {
	err := db.UserDirCollection.Remove(bson.M{"username": username})
	return err
}
