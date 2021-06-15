package service

import (
	"snns_srv/repository"
)

//GetNeuronSpiking :
func GetNeuronSpiking(gid string) (repository.NeuronP1, error) {
	return repository.GetNeuronSpiking(gid)
}

//GetNeuronV :
func GetNeuronV(gid string, tt int64) (repository.NeuronP2, error) {
	return repository.GetNeuronV(gid, tt)
}

//GetNeuronV2 :
func GetNeuronOne(gid string, id int64) (repository.NeuronP3, error) {
	return repository.GetNeuronOne(gid, id)
}

//GetNeuronV2 :
func GetNeuronOnes(gid string, num int64) (repository.NeuronP4, error) {
	return repository.GetNeuronOnes(gid, num)
}

//GetNeuronV23 :
func GetNeuronSpiking3(gid string, max int64) (repository.NeuronP1, error) {
	return repository.GetNeuronSpiking3(gid, max)
}

//GetNeuronLimitSpiking :
func GetNeuronLimitSpiking(gid string, min int64, max int64) (repository.NeuronP1, error) {
	return repository.GetNeuronLimitSpiking(gid, min, max)
}

//////////////////////////////////////////   读取二进制文件   //////////////////////////////////////
// 热力图数据
func GetHeatMapNvlFromBin(IndexSumFileName string, NeuSTDir string, t int) (repository.NeuronHM, error) {
	return repository.GetHeatMapNvlFromBin(IndexSumFileName, NeuSTDir, t)
}

// 范围脉冲序列点图
func GetSpkingFromBin(IndexSumFileName string, NeuSTDir string, minId int32, maxId int32) (repository.NeuronSTs, error) {
	return repository.GetSpkingFromBin(IndexSumFileName, NeuSTDir, minId, maxId)
}

// 范围脉冲电位
func GetMulNeuVlFromBin(IndexSumFileName string, NeuNlDir string, minId int32, maxId int32) (repository.NeuronVls, error) {
	return repository.GetMulNeuVlFromBin(IndexSumFileName, NeuNlDir, minId, maxId)
}
