package k8s_opt

import (
	"SNNClusterManagementSystem/ScriptsK8S/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"time"
)

// 获取指定命名空间中指定的deployment
func GetSpecDeployment(clientSet *kubernetes.Clientset,deploymentName string, nameSpace string) (deployment * appsv1.Deployment)  {
	deployment, err := clientSet.AppsV1().Deployments(nameSpace).Get(deploymentName, metav1.GetOptions{})
	common.CheckError(err)
	return
}

// 获取指定命名空间中的所有deployment
func GetSpecAllDeployment(clientSet *kubernetes.Clientset, nameSpace string) (deployments * appsv1.DeploymentList) {
	deployments, err := clientSet.AppsV1().Deployments(nameSpace).List(metav1.ListOptions{})
	common.CheckError(err)
	return
}

// 获取所有deployment
func GetAllDeployment(clientSet *kubernetes.Clientset) (deployments * appsv1.DeploymentList) {
	deployments, err := clientSet.AppsV1().Deployments("").List(metav1.ListOptions{})
	common.CheckError(err)
	return
}

// 根据用户名创建一个Pynn镜像的Deployment
func CreatePynnDeployment(clientSet *kubernetes.Clientset, ClientName string, nameSpace string) (flag bool) {
	// 获取文件生成的deployment对象
	baseDeployment := GetDeploymentFromYaml("./ScriptsK8S/yaml_files/base_pynn_deployment.yaml")
	// 修改名称
	// deploy
	baseName := baseDeployment.Name
	baseDeployment.Name = baseName + ClientName
	// container
	baseContainer := baseDeployment.Spec.Template.Spec.Containers
	baseName = baseContainer[0].Name
	baseDeployment.Spec.Template.Spec.Containers[0].Name = baseName + ClientName
	// 创建deployment
	if _, err := clientSet.AppsV1().Deployments(nameSpace).Get(baseDeployment.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			flag = false
			panic(err)
		}
		// 不存在则创建
		if _, err := clientSet.AppsV1().Deployments(nameSpace).Create(&baseDeployment); err != nil {
			flag = false
			panic(err)
		}
	}
	return true
}

// 更新deployment的replicas
func UpdateDeploymentReplicas(clientSet *kubernetes.Clientset, deployment * appsv1.Deployment, nameSpace string, replicas int32) (flag bool) {
	// 修改replicas数量
	deployment.Spec.Replicas = &replicas
	// 查询k8s是否有该deployment
	if _, err := clientSet.AppsV1().Deployments(nameSpace).Get(deployment.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			flag = false
			panic(err)
		}
		// 不存在则创建
		if _, err := clientSet.AppsV1().Deployments(nameSpace).Create(deployment); err != nil {
			flag = false
			panic(err)
		}
	} else {
		// 已存在则更新
		if _, err := clientSet.AppsV1().Deployments(nameSpace).Update(deployment); err != nil {
			flag = false
			panic(err)
		}
	}

	// 等待更新完成
	var k8sDeployment *appsv1.Deployment
	var err error
	for {
		// 获取k8s中deployment的状态
		if k8sDeployment, err = clientSet.AppsV1().Deployments(nameSpace).Get(deployment.Name, metav1.GetOptions{}); err != nil {
			goto RETRY
		}
		// 进行状态判定
		if k8sDeployment.Status.UpdatedReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.Replicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.AvailableReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.ObservedGeneration == k8sDeployment.Generation {
			// 滚动升级完成
			break
		}
	RETRY:
		time.Sleep(1 * time.Second)
	}
	fmt.Println("部署成功!")
	return true
}

// 更新deployment的pods
func UpdateDeploymentPods(clientSet *kubernetes.Clientset, deployment * appsv1.Deployment, nameSpace string, container v1.Container, option string) (flag bool) {
	var TmpContainers []v1.Container
	var AddContainers []v1.Container
	var UpdateContainers []v1.Container

	switch option {
	case "add":
		TmpContainers = deployment.Spec.Template.Spec.Containers
		AddContainers = append(TmpContainers, container)
		deployment.Spec.Template.Spec.Containers = AddContainers
		// 添加容器
		if _, err := clientSet.AppsV1().Deployments(nameSpace).Update(deployment); err != nil {
			flag = false
			panic(err)
		}
	case "update":
		UpdateContainers = append(TmpContainers, container)
		deployment.Spec.Template.Spec.Containers = UpdateContainers
		// 修改容器
		if _, err := clientSet.AppsV1().Deployments(nameSpace).Update(deployment); err != nil {
			flag = false
			panic(err)
		}
	default:
		panic("you need choose the option to the pods in deployment")
	}

	// 等待更新完成
	var k8sDeployment *appsv1.Deployment
	var err error
	for {
		// 获取k8s中deployment的状态
		if k8sDeployment, err = clientSet.AppsV1().Deployments(nameSpace).Get(deployment.Name, metav1.GetOptions{}); err != nil {
			goto RETRY
		}
		// 进行状态判定
		if k8sDeployment.Status.UpdatedReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.Replicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.AvailableReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.ObservedGeneration == k8sDeployment.Generation {
			// 滚动升级完成
			break
		}
	RETRY:
		time.Sleep(1 * time.Second)
	}
	fmt.Println("修改成功!")
	return true
}

// 删除指定的deployment
func DeleteSpecDeployment(clientSet *kubernetes.Clientset, deployName string, nameSpace string) (flag bool) {
	if _, err := clientSet.AppsV1().Deployments(nameSpace).Get(deployName, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			flag = false
			panic(err)
		}
		// 不存在则报错
		flag = false
		panic("deployment doesnt exist")
	} else {
		// 已存在则删除
		if err := clientSet.AppsV1().Deployments(nameSpace).Delete(deployName,&metav1.DeleteOptions{}); err != nil {
			flag = false
			panic(err)
		}
	}
	return true
}

// 从Yaml文件中解析出Deployment对象
func GetDeploymentFromYaml(YamlPath string) (deployment appsv1.Deployment) {

	// 读取yaml文件
	deployYaml, err := ioutil.ReadFile(YamlPath)
	common.CheckError(err)

	// yaml转json
	deployJson, err := yaml2.ToJSON(deployYaml)
	common.CheckError(err)

	// yaml转struct
	err = json.Unmarshal(deployJson, &deployment)
	common.CheckError(err)

	return deployment
}
