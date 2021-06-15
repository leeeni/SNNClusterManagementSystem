package service

import (
	"snns_srv/repository"
	//	"strings"
)

//CheckUserExistByUsername ：
func CheckUserExistByUsername(username string) (bool, error) {
	return repository.CheckExistByUsername(username)
}

//CheckUserExistByEmail :
func CheckUserExistByEmail(email string) (bool, error) {
	return repository.CheckExistByEmail(email)
}

//InsertUser :
func InsertUser(user *repository.User) error {
	return repository.InsertUser(user)
}

//SelectUser :
// func SelectNode(pageindex int ,pagerows int) ([]repository.Node,int, error) {
func SelectUser() ([]repository.User, error) {
	print("Enter Service SelectUser!")
	return repository.GetUserAll()
}

//GetUserByAccount :
func GetUserByAccount(account string) (user repository.User) {
	// if strings.Contains(account, "@") {
	// 	user = repository.GetUserByEmail(account)
	// } else {
	user = repository.GetUserByUsername(account)
	// }
	return user
}

// DelUser :
func DelUser(id string) error {
	print("Enter Service DelUser!")
	return repository.DelUser(id)
}