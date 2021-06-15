package service

import (
	"snns_srv/repository"
)

//GetCPUMEMs :
func GetCPUMEMs(hostname string, len int) (repository.MonitorDatas, error) {
	return repository.GetMonitorDatas(hostname, len)
}
