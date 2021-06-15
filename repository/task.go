package repository

import (
	"snns_srv/db"

	"gopkg.in/mgo.v2/bson"
)

//Task :
type Task struct {
	UID      bson.ObjectId `bson:"_id,omitempty"`
	TaskID   string `bson:"taskid"`
	Name     string `bson:"name"`
	Describe string `bson:"describe"`
	User     string `bson:"user"`

	// CreatedTime and LastLogin use timestamp.
	//	CreatedTime int64 `bson:"created_time"`
	//	LastRun     int64 `bson:"last_run"`
	//	IsBanned    bool  `bson:"is_banned"`
}

//SetBanned :
// func (task *Task) SetBanned(banned bool) bool {
// 	task.IsBanned = banned
// 	return banned
// }

//CheckBanned :
// func (task *Task) CheckBanned() bool {
// 	return task.IsBanned
// }

//InsertTask :
func InsertTask(task *Task) error {
	return Insert(db.TaskCollection, task)
}

//DelTask :
func DelTask(taskid string) error {
	return DelByID(db.TaskCollection, taskid)
}

//GetTaskAll :
func GetTaskAll(username string) ([]Task, error) {
	var tasks []Task
	// err := NodeCollection.Find(nil).Limit(pagerows).All(&nodes)
	print(username)
	err := db.TaskCollection.Find(bson.M{"user": username}).All(&tasks)
	print(tasks)
	return tasks, err
}

//GetTaskByName :
func GetTaskByName(name string) Task {
	task := Task{}
	GetByQ(db.TaskCollection, bson.M{"name": name}, &task)
	return task
}

//GetTaskByUID :
func GetTaskByUID(uid bson.ObjectId) Task {
	task := Task{}
	GetByQ(db.TaskCollection, bson.M{"_id": uid}, &task)
	return task
}
