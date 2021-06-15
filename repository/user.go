package repository

import (
	"snns_srv/db"
	"snns_srv/util/log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//User :用户数据
type User struct {
	// Uid: Primary key (_id)
	UID      bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Verified bool          `bson:"verified"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	// CreatedTime and LastLogin use timestamp.
	CreatedTime int64 `bson:"created_time"`
	LastLogin   int64 `bson:"last_login"`
	RoleSys     bool  `bson:"rolesys"`
	RoleAdmin   bool  `bson:"roleadmin"`
	State       int   `bson:"state"`
	IsPclUser   bool  `bson:"is_pcluser"`
	// State => 0:NewReg , 1:Active , 2: Inactive
}

//SetPassword :
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Logger.Errorf("Encrypt password error: %v", err)
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword :
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

//CheckExistByUsername :
func CheckExistByUsername(username string) (bool, error) {
	return Has(db.UserCollection, bson.M{"username": username})
}

//CheckExistByEmail :
func CheckExistByEmail(email string) (bool, error) {
	return Has(db.UserCollection, bson.M{"email": email})
}

//InsertUser :
func InsertUser(user *User) error {
	return Insert(db.UserCollection, user)
}

//GetUserAll :
func GetUserAll() ([]User, error) {
	var users []User
	// err := NodeCollection.Find(nil).Limit(pagerows).All(&nodes)
	err := db.UserCollection.Find(nil).All(&users)
	return users, err
}

//GetUserByUsername :
func GetUserByUsername(username string) User {
	user := User{}
	GetByQ(db.UserCollection, bson.M{"username": username}, &user)
	return user
}

//GetUserByEmail :
func GetUserByEmail(email string) User {
	user := User{}
	GetByQ(db.UserCollection, bson.M{"email": email}, &user)
	return user
}

//DelUser :
func DelUser(id string) error {
	return DelByID(db.UserCollection, id)
}
