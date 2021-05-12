package service

import (
	"SNNClusterManagementSystem/ScriptsK8S/common"
	"SNNClusterManagementSystem/ScriptsK8S/k8s_opt"
	"SNNClusterManagementSystem/repository"

	v1 "k8s.io/api/core/v1"
)



func CheckPodExistByUsername(username string) (bool, error) {
	return repository.CheckPodExistByUsername(username)
}

func CheckPodExistByPodName(podname string) (bool, error) {
	return repository.CheckPodExistByPodName(podname)
}

func InsertPod(pod *repository.PynnPod) error {
	return repository.InsertPod(pod)
}

func GetPodByUsername(username string) (pod repository.PynnPod) {
	pod = repository.GetPodByUsername(username)
	return
}

// 在服务器上创建Pod
func BuildPod(username string) bool {
	clientset,err := common.InitClient()
	common.CheckError(err)
	return k8s_opt.CreatePynnPod(clientset,username,"pynn-clients")
}


// 创建服务器上的pod，带cpu和memory信息的
func BuildPodByCpuAndMemoryFromServer(username string, cpu int64, memory int64) bool {
	clientset,err := common.InitClient()
	common.CheckError(err)
	return k8s_opt.CreatePynnPodUpdateCpuAndMemory(clientset,username,"pynn-clients",cpu,memory)
}

// 在服务器上获取Pod
func GetPodByUsernameFromServer(username string)  * v1.Pod  {
	clientset,err := common.InitClient()
	common.CheckError(err)
	pod := k8s_opt.GetSpecPod(clientset,"pynn-clients","pynn-pod-"+username)
	return pod
}

// 数据库中删除记录
func DeletePodFromDB(username string) error {
	return repository.DeletePodByUserName(username)
}

// 服务器上删除目录
func DeletPodFromServer(username string) bool {
	podname := "pynn-pod-"+username
	clientset,err := common.InitClient()
	common.CheckError(err)
	return k8s_opt.DeleteSpecPod(clientset, podname, "pynn-clients")
}

// 更新服务器上的pod
func UpdatePodFromServer(pod * v1.Pod) bool {
	clientset,err := common.InitClient()
	common.CheckError(err)
	return k8s_opt.UpdateSpecPod(clientset,"pynn-clients", pod)
}

// 更新数据库上的pod
func UpdatePodFromDB(pod repository.PynnPod) bool {
	return repository.UpdatePod(pod) != nil
}


