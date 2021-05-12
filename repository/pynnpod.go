package repository

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	NormalUserDir = 0
)

type PynnPod struct {
	// Pid: Primary key (_id)
	Uid      bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	PodName  string        `bson:"podname"`
	PodIp    string        `bson:"podip"`
	PodStatus    string    `bson:"podstatus"`
	PodCpu		string		`bson:"podcpu"`
	PodMemory   string		`bson:"podmemory"`
	// CreatedTime and LastLogin use timestamp.
	CreatedTime int64 `bson:"created_time"`
	LastUse     int64 `bson:"last_use"`
	Role        int   `bson:"role"`
	IsBanned    bool  `bson:"is_banned"`
}

// 检索
func CheckPodExistByUsername(username string) (bool, error) {
	return Has(PyNNPodCollection, bson.M{"username": username})
}

// 检索
func CheckPodExistByPodName(podname string) (bool, error) {
	return Has(PyNNPodCollection, bson.M{"podname": podname})
}

// 增
func InsertPod(pod *PynnPod) error {
	return Insert(PyNNPodCollection, pod)
}

// 查
func GetPodByUsername(username string) PynnPod {
	pod := PynnPod{}
	GetByQ(PyNNPodCollection, bson.M{"username": username}, &pod)
	return pod
}

// 查
func GetPodByUid(uid bson.ObjectId) PynnPod {
	pod := PynnPod{}
	GetByQ(PyNNPodCollection, bson.M{"_id": uid}, &pod)
	return pod
}

// 改
func UpdatePod(pod PynnPod) error {
	err := PyNNPodCollection.Update(bson.M{"_id": pod.Uid}, pod)
	return err
}

// 删
func DeletePodByUserName(username string) error {
	err := PyNNPodCollection.Remove(bson.M{"username":username})
	return err
}