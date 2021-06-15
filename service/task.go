package service

import (
	"snns_srv/repository"
)

// InsertTask :
func InsertTask(task *repository.Task) error {
	return repository.InsertTask(task)
}

// SelectTask :
func SelectTask(username string) ([]repository.Task, error) {
	print("Enter Service SelectTask!")
	return repository.GetTaskAll(username)
}

// DelTask :
func DelTask(taskid string) error {
	print("Enter Service DelTask!")
	return repository.DelTask(taskid)
}
