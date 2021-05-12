package repository

import (
	"SNNClusterManagementSystem/util/log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const (
	NormalUser = 0
	Admin      = 1
	SuperAdmin = 2
)

// 与数据库表对应的user结构体
type User struct {
	// Uid: Primary key (_id)
	Uid         bson.ObjectId `bson:"_id,omitempty"`
	Email       string        `bson:"email"`
	Verified    bool          `bson:"verified"`
	Username    string        `bson:"username"`
	Password    string        `bson:"password"`
	CreatedTime int64         `bson:"created_time"`
	LastLogin   int64         `bson:"last_login"`
	Role        int           `bson:"role"`
	IsBanned    bool          `bson:"is_banned"`
}

// user的成员函数
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Logger.Errorf("Encrypt password error: %v", err)
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// 检索
func CheckExistByUsername(username string) (bool, error) {
	return Has(UserCollection, bson.M{"username": username})
}
// 检索
func CheckExistByEmail(email string) (bool, error) {
	return Has(UserCollection, bson.M{"email": email})
}

// 增
func InsertUser(user *User) error {
	return Insert(UserCollection, user)
}

// 查
func GetUserByUsername(username string) User {
	user := User{}
	GetByQ(UserCollection, bson.M{"username": username}, &user)
	return user
}

// 查
func GetUserByEmail(email string) User {
	user := User{}
	GetByQ(UserCollection, bson.M{"email": email}, &user)
	return user
}

// 改
func UpdateUser(user User) error {
	err := UserCollection.Update(bson.M{"_id": user.Uid}, user)
	return err
}

// 删
func DeleteUser(user User) error {
	err := UserCollection.Remove( bson.M{"_id": user.Uid})
	return err
}