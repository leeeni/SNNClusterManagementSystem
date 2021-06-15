package service

import (
	"snns_srv/repository"
)

//GetMemByHost :
func GetMemByHost( hostname string) (repository.Mem, error) {
	return repository.GetMemByHost(hostname)
}
