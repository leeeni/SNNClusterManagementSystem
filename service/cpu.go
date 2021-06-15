package service

import (
	"snns_srv/repository"
)

//GetCPUByHost :
func GetCPUByHost(hostname string) (repository.CPU, error) {
	return repository.GetCPUByHost(hostname)
}
