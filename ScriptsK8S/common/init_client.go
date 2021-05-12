package common

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)



// 初始化k8s客户端
func InitClient() (clientSet *kubernetes.Clientset, err error){

	// 获取config
	config, err := getConfig()

	// 根据指定的 config 创建一个新的 clientSet
	clientSet, err = kubernetes.NewForConfig(config)

	return
}

// 获取config配置
func getConfig() (config *rest.Config, err error)  {

	//var kubeconfig * string
	// 配置 k8s 集群外 kubeconfig 配置文件
	//kubeconfig = flag.String("kubeconfig", "/root/.kube/config", "absolute path to the kubeconfig file")
	//flag.Parse()
	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err = clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	CheckError(err)
	return
}

func CheckError(err error){
	if err != nil{
		panic(err)
	}
}