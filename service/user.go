package service

import (
	"SNNClusterManagementSystem/repository"

	"strings"
)

// 检索
func CheckUserExistByUsername(username string) (bool, error) {
	return repository.CheckExistByUsername(username)
}

// 检索
func CheckUserExistByEmail(email string) (bool, error) {
	return repository.CheckExistByEmail(email)
}

// 增
func InsertUser(user *repository.User) error {
	return repository.InsertUser(user)
}

// 查
func GetUserByAccount(account string) (user repository.User) {
	if strings.Contains(account, "@") {
		user = repository.GetUserByEmail(account)
	} else {
		user = repository.GetUserByUsername(account)
	}
	return
}

// 改
func UpdateUser(user repository.User) error {
	return repository.UpdateUser(user)
}

// 删
func DeleteUser(user repository.User) error {
	return repository.DeleteUser(user)
}