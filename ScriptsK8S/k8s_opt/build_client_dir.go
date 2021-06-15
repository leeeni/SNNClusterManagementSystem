package k8s_opt

import (
	"os"
)

func BuildClientDir(ClientName string) (flag bool) {
	ClientDir := "/home/work/ClientDir/"
	clientPath := ClientDir + ClientName

	// 查看路径是否存在
	IsNotExist := CheckDirExit(clientPath)

	// 创建文件夹
	if IsNotExist == false {
		// 如果不存在，则创建
		_ = os.MkdirAll(clientPath, 0777)
	} else {
		// 如果存在，则报错
		flag = false
		return flag
	}
	// 再判断是否创建成功
	IfExist := CheckDirExit(clientPath)
	if IfExist == false {
		// 如果不存在，则创建失败
		flag = false
		return flag
	} else {
		flag = true
	}
	return flag
}

func DeleteClientDir(ClientName string) (flag bool) {
	// 本质不是删除是移动到垃圾回收站

	ClientDir := "/home/work/ClientDir/"
	clientPath := ClientDir + ClientName

	// 查看路径是否存在
	IsNotExist := CheckDirExit(clientPath)

	// 创建文件夹
	if IsNotExist == false {
		// 如果不存在，则报错
		flag = false
		return flag
	} else {
		// 如果存在，则移动
		_ = os.RemoveAll(clientPath)
	}
	// 再判断是否移动成功
	IfExist := CheckDirExit(clientPath)
	if IfExist == false {
		// 如果不存在，则移动成功
		flag = true

	} else {
		// 如果存在，则返回错误
		flag = false
		return flag
	}
	return flag
}

func CheckDirExit(path string) bool {
	/*
	   判断文件或文件夹是否存在
	   如果返回的错误为nil,说明文件或文件夹存在
	   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	   如果返回的错误为其它类型,则不确定是否在存在
	*/
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
