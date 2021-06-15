package service

import (
	"snns_srv/ScriptsK8S/k8s_opt"
	"snns_srv/repository"
)

// 在服务器上创建目录
func BuildUserDir(username string) bool {
	return k8s_opt.BuildClientDir(username)
}

// 通过用户名查看目录是否存在
func CheckUserDirExistByUsername(username string) (bool, error) {
	return repository.CheckDirExistByUsername(username)
}

// 通过目录名查看目录是否存在
func CheckUserDirExistByDirname(dirname string) (bool, error) {
	return repository.CheckDirExistByDirName(dirname)
}

// 插入用户目录记录
func InsertUserDir(userdir *repository.UserDir) error {
	return repository.InsertUserDir(userdir)
}

// 通过用户名获取用户目录信息
func GetUserDirByUsername(username string) (userdir repository.UserDir) {
	userdir = repository.GetUserDirByUsername(username)
	return
}

// 改
func UpdateUserDir(userdir repository.UserDir) error {
	return repository.UpdateUserDir(userdir)
}

// 数据库中删除记录
func DeleteUserDirFromDB(username string) error {
	return repository.DeleteUserDirByUserName(username)
}

// 服务器上删除目录
func DeleteUserDirFromServer(username string) bool {
	return k8s_opt.DeleteClientDir(username)
}
