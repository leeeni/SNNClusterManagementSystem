package neuron

import (
	"snns_srv/controller/v1/common"
	"snns_srv/repository"
	"snns_srv/service"
	"strconv"

	"github.com/kataras/iris/v12"
)

// GetNeuronSpikingResponseData :
type GetNeuronSpikingResponseData struct{}

// GetNeuronSpiking :
func GetNeuronSpiking(ctx iris.Context) {

	//var req SelectNodeRequest
	var neuronp1 repository.NeuronP1
	// service.GetNeuronSpiking(ctx.Params().Get("gid"))
	neuronp1, err := service.GetNeuronSpiking(ctx.Params().Get("gid"))
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(neuronp1)
	common.SuccessResponse(ctx, neuronp1)
}

// GetNeuronVResponseData :
type GetNeuronVResponseData struct{}

// GetNeuronSpiking :
func GetNeuronV(ctx iris.Context) {

	var neuronp2 repository.NeuronP2
	print(ctx.Params().Get("gid") + "--" + ctx.Params().Get("tt"))
	strT := ctx.Params().Get("tt")
	var left int
	var right int
	var index int = 0
	for _, str := range strT {
		if str == '=' {
			left = index + 1
		}
		if str == ')' {
			right = index
		}
		index = index + 1
	}

	tt, _ := strconv.ParseInt(strT[left:right], 10, 64)
	print(tt)
	neuronp2, err := service.GetNeuronV(ctx.Params().Get("gid"), tt)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, neuronp2)
}

// GetNeuronOneV :
func GetNeuronOneV(ctx iris.Context) {

	var neuronp3 repository.NeuronP3
	print(ctx.Params().Get("gid") + "--" + ctx.Params().Get("id"))
	strID := ctx.Params().Get("id")
	var left int
	var right int
	var index int = 0
	for _, str := range strID {
		if str == '=' {
			left = index + 1
		}
		if str == ')' {
			right = index
		}
		index = index + 1
	}

	id, _ := strconv.ParseInt(strID[left:right], 10, 64)
	print(id)
	neuronp3, err := service.GetNeuronOne(ctx.Params().Get("gid"), id)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, neuronp3)
}

// GetNeuronSpiking :
func GetNeuronOnesV(ctx iris.Context) {

	// id就是num，现在还没改

	var neuronp4 repository.NeuronP4
	print(ctx.Params().Get("gid") + "--" + ctx.Params().Get("id"))
	strID := ctx.Params().Get("id")
	var left int
	var right int
	var index int = 0
	for _, str := range strID {
		if str == '=' {
			left = index + 1
		}
		if str == ')' {
			right = index
		}
		index = index + 1
	}

	id, _ := strconv.ParseInt(strID[left:right], 10, 64)
	print(id)
	neuronp4, err := service.GetNeuronOnes(ctx.Params().Get("gid"), id)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	//fmt.Print(nodes)
	common.SuccessResponse(ctx, neuronp4)
}

// GetNeuronSpiking :
func GetNeuronSpiking3(ctx iris.Context) {

	// id就是max，现在还没改

	strID := ctx.Params().Get("id")
	var left int
	var right int
	var index int = 0
	for _, str := range strID {
		if str == '=' {
			left = index + 1
		}
		if str == ')' {
			right = index
		}
		index = index + 1
	}

	var neuronp1 repository.NeuronP1

	id, _ := strconv.ParseInt(strID[left:right], 10, 64)

	neuronp1, err := service.GetNeuronSpiking3(ctx.Params().Get("gid"), id)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	common.SuccessResponse(ctx, neuronp1)
}

// GetNeuronSpiking :
func GetNeuronLimitSpiking(ctx iris.Context) {

	// id就是max，现在还没改

	strID := ctx.Params().Get("id")
	var left int
	var right int
	var index int = 0
	for _, str := range strID {
		if str == '=' {
			left = index + 1
		}
		if str == ')' {
			right = index
		}
		index = index + 1
	}

	var neuronp1 repository.NeuronP1

	id, _ := strconv.ParseInt(strID[left:right], 10, 64)

	neuronp1, err := service.GetNeuronSpiking3(ctx.Params().Get("gid"), id)
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	common.SuccessResponse(ctx, neuronp1)
}

//////////////////////////////////////////   读取二进制文件   //////////////////////////////////////
// 热力图回复数据
type GetHeatMapNvlFromBinRespond struct {
	NeuNum   int
	Time     int
	NeuronHM repository.NeuronHM
}

// 热力图数据
func GetHeatMapNvlFromBin(ctx iris.Context) {

	gid := ctx.Params().Get("gid")
	tt, _ := strconv.Atoi(ctx.Params().Get("tt"))

	username := ctx.Params().Get("username")
	user := service.GetUserByAccount(username)
	uid := user.UID.Hex()
	dataPath := "/home/work/ClientDir/" + uid + "/" + gid

	var neuronHM repository.NeuronHM
	// neuronHM, err := service.GetHeatMapNvlFromBin("/home/work/TaskData/"+gid+"/index_sum", "/home/work/TaskData/"+gid+"/neu_rl", tt)
	// neuronNum, _, time := repository.ReadIndexSumFile("/home/work/TaskData/" + gid + "/index_sum")
	neuronHM, err := service.GetHeatMapNvlFromBin(dataPath+"/index_sum", dataPath+"/neu_rl", tt)
	neuronNum, _, time := repository.ReadIndexSumFile(dataPath + "/index_sum")
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, GetHeatMapNvlFromBinRespond{NeuNum: int(neuronNum), Time: int(time), NeuronHM: neuronHM})
}

// 范围脉冲序列回复数据
type GetSpkingFromBinRespond struct {
	NeuNum    int
	Time      int
	NeuronSTs repository.NeuronSTs
}

// 范围脉冲序列点图
func GetSpkingFromBin(ctx iris.Context) {

	gid := ctx.Params().Get("gid")
	minIdInt, _ := strconv.Atoi(ctx.Params().Get("minId"))
	minId := int32(minIdInt)
	maxIdInt, _ := strconv.Atoi(ctx.Params().Get("maxId"))
	maxId := int32(maxIdInt)

	username := ctx.Params().Get("username")
	user := service.GetUserByAccount(username)
	uid := user.UID.Hex()
	dataPath := "/home/work/ClientDir/" + uid + "/" + gid

	var neuronSTs repository.NeuronSTs

	// neuronSTs, err := service.GetSpkingFromBin("/home/work/TaskData/"+gid+"/index_sum", "/home/work/TaskData/"+gid+"/neu_st", minId, maxId)
	// neuronNum, _, time := repository.ReadIndexSumFile("/home/work/TaskData/" + gid + "/index_sum")
	neuronSTs, err := service.GetSpkingFromBin(dataPath+"/index_sum", dataPath+"/neu_st", minId, maxId)
	neuronNum, _, time := repository.ReadIndexSumFile(dataPath + "/index_sum")
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}

	var neuronSTsRespond repository.NeuronSTs
	for t := 0; t < int(time); t++ {
		for stIndex := 0; stIndex < len(neuronSTs.Nst); stIndex++ {
			if int(neuronSTs.Nst[stIndex].Time) == t {
				neuronSTsRespond.Nst = append(neuronSTsRespond.Nst, neuronSTs.Nst[stIndex])
			}
		}
	}
	common.SuccessResponse(ctx, GetSpkingFromBinRespond{NeuNum: int(neuronNum), Time: int(time), NeuronSTs: neuronSTsRespond})
}

// 范围脉冲序列回复数据
type GetMulNeuVlFromBinRespond struct {
	NeuNum    int
	NeuronVls repository.NeuronVls
}

// 范围脉冲电位
func GetMulNeuVlFromBin(ctx iris.Context) {
	gid := ctx.Params().Get("gid")
	minIdInt, _ := strconv.Atoi(ctx.Params().Get("minId"))
	minId := int32(minIdInt)
	maxIdInt, _ := strconv.Atoi(ctx.Params().Get("maxId"))
	maxId := int32(maxIdInt)

	username := ctx.Params().Get("username")
	user := service.GetUserByAccount(username)
	uid := user.UID.Hex()
	dataPath := "/home/work/ClientDir/" + uid + "/" + gid

	var neuronVls repository.NeuronVls

	// neuronVls, err := service.GetMulNeuVlFromBin("/home/work/TaskData/"+gid+"/index_sum", "/home/work/TaskData/"+gid+"/neu_vl", minId, maxId)
	// neuronNum, _, _ := repository.ReadIndexSumFile("/home/work/TaskData/" + gid + "/index_sum")
	neuronVls, err := service.GetMulNeuVlFromBin(dataPath+"/index_sum", dataPath+"/neu_vl", minId, maxId)
	neuronNum, _, _ := repository.ReadIndexSumFile(dataPath + "/index_sum")
	if err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, GetMulNeuVlFromBinRespond{NeuNum: int(neuronNum), NeuronVls: neuronVls})
}
