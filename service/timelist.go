package service

import (
	"snns_srv/repository"
)

//InsertTimelist :
func InsertTimelist(Timelist *repository.Timelist) error {
	return repository.InsertTimelist(Timelist)
}

//UpdateTimelist :
func UpdateTimelist(Timelist *repository.Timelist) error {
	return repository.UpdateTimelist(Timelist)
}

//SelectTimelist :
// func SelectTimelist(pageindex int ,pagerows int) ([]repository.Timelist,int, error) {
func SelectTimelist() ([]repository.Timelist, error) {
	print("Enter Service SelectTimelist!")
	return repository.GetTimelistAll()
}
func SelectTimelistun(username string) ([]repository.Timelist, error) {
	print("Enter Service SelectTimelist!")
	return repository.GetTimelistAllun(username)
}

// DelTimelist :
func DelTimelist(id string) error {
	print("Enter Service DelTimelist!")
	return repository.DelTimelist(id)
}
