package service

import (
	"snns_srv/repository"
)

//InsertNode :
func InsertNode(Node *repository.Node) error {
	return repository.InsertNode(Node)
}

//UpdateNode :
func UpdateNode(Node *repository.Node) error {
	return repository.UpdateNode(Node)
}

//SelectNode :
// func SelectNode(pageindex int ,pagerows int) ([]repository.Node,int, error) {
func SelectNode() ([]repository.Node, error) {
	print("Enter Service SelectNode!")
	return repository.GetNodeAll()
}
// DelNode :
func DelNode(id string) error {
	print("Enter Service DelNode!")
	return repository.DelNode(id)
}


