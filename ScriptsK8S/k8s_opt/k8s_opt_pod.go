package k8s_opt

import (
	"SNNClusterManagementSystem/ScriptsK8S/common"
	"encoding/json"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"time"
)


// 从Yaml文件中解析出Pod2对象
func GetPodFromYaml(YamlPath string) (pod v1.Pod) {
	// 读取yaml文件
	PodYaml, err := ioutil.ReadFile(YamlPath)
	common.CheckError(err)

	// yaml转json
	PodJson, err := yaml2.ToJSON(PodYaml)
	common.CheckError(err)

	// yaml转struct
	_ = json.Unmarshal(PodJson, &pod)

	return pod
}

// 创建一个新的pod
func CreatePynnPod(clientSet *kubernetes.Clientset,  ClientName string, nameSpace string) (flag bool) {

	// 获取文件生成的deployment对象
	basePod := GetPodFromYaml("./ScriptsK8S/yaml_files/base_pynn_pod.yaml")

	// 修改名称
	// Pod
	baseName := basePod.Name
	basePod.Name = baseName + ClientName
	// Container
	baseName = basePod.Spec.Containers[0].Name
	basePod.Spec.Containers[0].Name = baseName + ClientName
	// 修改容器挂载卷目录
	var baseHostPathSrc v1.HostPathVolumeSource
	baseHostPath := "/home/work/ClientDir/" + ClientName
	baseHostPathSrc.Path = baseHostPath
	basePod.Spec.Volumes[0].HostPath = &baseHostPathSrc

	// 创建Pod
	if _, err := clientSet.CoreV1().Pods(nameSpace).Get(basePod.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			flag = false
			panic(err)
		}
		// 不存在则创建
		if _, err := clientSet.CoreV1().Pods(nameSpace).Create(&basePod); err != nil {
			flag = false
			panic(err)
		}
	}

	// 等待创建完成
	var k8sPod *v1.Pod
	var err error
	for {
		// 获取k8s中deployment的状态
		if k8sPod, err = clientSet.CoreV1().Pods(nameSpace).Get(basePod.Name, metav1.GetOptions{}); err != nil {
			goto RETRY
		}
		// 进行状态判定
		if string(k8sPod.Status.Phase) == "Running"{
			// 创建完成
			break
		}
	RETRY:
		time.Sleep(1 * time.Millisecond)

	}
	return true
}

// 创建一个新的pod
func CreatePynnPodUpdateCpuAndMemory(clientSet *kubernetes.Clientset,  ClientName string, nameSpace string, cpu int64,memory int64) (flag bool) {

	// 获取文件生成的deployment对象
	basePod := GetPodFromYaml("./ScriptsK8S/yaml_files/base_pynn_pod.yaml")

	// 修改名称
	// Pod
	baseName := basePod.Name
	basePod.Name = baseName + ClientName
	// Container
	baseName = basePod.Spec.Containers[0].Name
	basePod.Spec.Containers[0].Name = baseName + ClientName
	// 修改容器挂载卷目录
	var baseHostPathSrc v1.HostPathVolumeSource
	baseHostPath := "/home/work/ClientDir/" + ClientName
	baseHostPathSrc.Path = baseHostPath
	basePod.Spec.Volumes[0].HostPath = &baseHostPathSrc
	// 修改cpu和memory
	cpuData := basePod.Spec.Containers[0].Resources.Requests["cpu"]
	memoryData := basePod.Spec.Containers[0].Resources.Requests["memory"]
	cpuData.Set(cpu)
	memoryData.Set(memory)
	basePod.Spec.Containers[0].Resources.Requests["cpu"] = cpuData
	basePod.Spec.Containers[0].Resources.Requests["memory"] = memoryData

	// 创建Pod
	if _, err := clientSet.CoreV1().Pods(nameSpace).Get(basePod.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			flag = false
			panic(err)
		}
		// 不存在则创建
		if _, err := clientSet.CoreV1().Pods(nameSpace).Create(&basePod); err != nil {
			flag = false
			panic(err)
		}
	}

	// 等待创建完成
	var k8sPod *v1.Pod
	var err error
	for {
		// 获取k8s中deployment的状态
		if k8sPod, err = clientSet.CoreV1().Pods(nameSpace).Get(basePod.Name, metav1.GetOptions{}); err != nil {
			goto RETRY
		}
		// 进行状态判定
		if string(k8sPod.Status.Phase) == "Running"{
			// 创建完成
			break
		}
	RETRY:
		time.Sleep(1 * time.Millisecond)

	}
	return true
}

// 删除指定的Pod
func DeleteSpecPod(clientSet *kubernetes.Clientset, PodName string, nameSpace string) (flag bool) {
	if _, err := clientSet.CoreV1().Pods(nameSpace).Get(PodName, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			flag = false
			panic(err)
		}
		// 不存在则报错
		flag = false
		panic("Pods doesnt exist")
	} else {
		// 已存在则删除
		if err := clientSet.CoreV1().Pods(nameSpace).Delete(PodName,&metav1.DeleteOptions{}); err != nil {
			flag = false
			panic(err)
		}
	}

	// 保证完全删除，能获取说明还存在
	for {
		// 服务器上不存在，则退出
		if _, err := clientSet.CoreV1().Pods(nameSpace).Get(PodName, metav1.GetOptions{});err != nil{
			break
		}
		time.Sleep(1 * time.Millisecond)
	}

	return true
}

// 获取指定命名空间中的所有Pod
func GetSpecAllPod(clientSet *kubernetes.Clientset, nameSpace string) (pod * v1.PodList) {
	podList, err := clientSet.CoreV1().Pods(nameSpace).List(metav1.ListOptions{})
	common.CheckError(err)
	return podList
}

// 获取指定命名空间中的指定Pod
func GetSpecPod(clientSet *kubernetes.Clientset, nameSpace string, podName string) (pod * v1.Pod) {
	pod, err := clientSet.CoreV1().Pods(nameSpace).Get(podName,metav1.GetOptions{})
	common.CheckError(err)
	return pod
}

// 更新Pod
func UpdateSpecPod(clientSet *kubernetes.Clientset, nameSpace string, pod * v1.Pod) bool {
	_, err := clientSet.CoreV1().Pods(nameSpace).Update(pod)
	return err != nil
}