package repository

import (
	"fmt"
	"log"
	"math"
	"time"

	"snns_srv/db"

	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
	"strconv"
)

// NeuronP1
type NeuronP1 struct {
	Sid   int        `bson:"sid"`
	Stime int64      `bson:"stime"`
	Nst   []NeuronST `bson:"nst"`
}

// NeuronST :
type NeuronST struct {
	// Uid: Primary key (_id)
	Time int64 `bson:"time"`
	Id   int   `bson:"id"`
}

// NeuronP2
type NeuronP2 struct {
	Sid int       `bson:"sid"`
	Mvl float64   `bson:"mvl"`
	Nvl []NeuronV `bson:"nvl"`
}

type NeuronOneV struct {
	Time int64 `bson:"time"`
	Vl   int   `bson:"vl"`
}

// NeuronP3
type NeuronP3 struct {
	Nvl []float64 `bson:"nvl"`
}

// NeuronV :
type NeuronV struct {
	// Uid: Primary key (_id)
	//time int     `bson:"time"`
	Id int `bson:"id"`
	//st   int     `bson:"st"`
	Vl float64 `bson:"vl"`
}

// NeuronP4
type NeuronP4 struct {
	Neu []NeuronP3
}

//  查询指定ID“范围”所有周期的数据 :
func GetNeuronOnes(gid string, num int64) (NeuronP4, error) {
	var err error
	// 神经元数组
	neuronp4 := NeuronP4{}
	var i int64
	for i = num - 99; i <= num; i++ {
		var tmpNeu = NeuronP3{}
		// 取数据
		qs := "select *  from \"" + gid + "\"  where id='" + strconv.FormatInt(i, 10) + "'"
		res, err := db.QueryInfluxDB(db.Conntsdb, "neuron_state_http", qs)
		if err != nil {
			fmt.Printf("err !")
		}
		if len(res[0].Series[0].Values) == 0 {
			continue
		}
		for _, row := range res[0].Series[0].Values {
			vs := fmt.Sprint(row[3])
			vv, _ := strconv.ParseFloat(vs, 64)
			tmpNeu.Nvl = append(tmpNeu.Nvl, vv)
		}
		neuronp4.Neu = append(neuronp4.Neu, tmpNeu)
	}
	return neuronp4, err
}

//  查询指定ID所有周期的数据 : :
func GetNeuronOne(gid string, id int64) (NeuronP3, error) {
	neuronp3 := NeuronP3{}
	qs := "select *  from \"" + gid + "\"  where id='" + strconv.FormatInt(id, 10) + "'"

	res, err := db.QueryInfluxDB(db.Conntsdb, "neuron_state_http", qs)
	if err != nil {
		fmt.Printf("err !")
	}
	for _, row := range res[0].Series[0].Values {
		vs := fmt.Sprint(row[3])
		vv, _ := strconv.ParseFloat(vs, 64)
		neuronp3.Nvl = append(neuronp3.Nvl, vv)
	}

	return neuronp3, err
}

// 查询指定周期的所有神经元的电位值 :
func GetNeuronV(gid string, tt int64) (NeuronP2, error) {
	neuronp2 := NeuronP2{}
	//neuronv := NeuronV{}
	var maxvl float64
	maxvl = 0.0
	maxid := 0
	qs := "select *  from \"" + gid + "\"  where time=" + strconv.FormatInt(tt, 10)

	res, err := db.QueryInfluxDB(db.Conntsdb, "neuron_state_http", qs)
	if err != nil {
		fmt.Printf("err !")
	}
	for _, row := range res[0].Series[0].Values {
		// t, err := time.Parse(time.RFC3339, row[0].(string))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//t1 := t.UnixNano()
		//fmt.Println(reflect.TypeOf(row[1]))
		id, _ := strconv.Atoi(row[1].(string))
		vs := fmt.Sprint(row[3])
		vv, _ := strconv.ParseFloat(vs, 64)
		// fmt.Println(i)
		//fmt.Println(t)

		//fmt.Println(t.UnixNano())
		neuronp2.Nvl = append(neuronp2.Nvl, NeuronV{Id: id, Vl: vv})
		if math.Abs(vv) > maxvl {
			maxvl = math.Abs(vv)
		}
		if id > maxid {
			maxid = id
		}
	}
	neuronp2.Sid = maxid
	neuronp2.Mvl = maxvl
	//var v float64
	// row := res[0].Series[0].Values[0]
	// vs := fmt.Sprint(row[1])
	// f, _ := strconv.ParseFloat(vs, 64)
	// v, _ := decimal.NewFromFloat(f).Round(2).Float64()
	// fmt.Println(v)
	return neuronp2, err
}

// 查询所有周期被激活的神经元ID :
func GetNeuronSpiking(gid string) (NeuronP1, error) {

	var neuronp1 NeuronP1
	var maxt int64
	maxt = 0
	maxid := 0
	qs := "select time,id,st  from \"" + gid + "\"  where st=1"

	res, err := db.QueryInfluxDB(db.Conntsdb, "neuron_state_http", qs)
	if err != nil {
		fmt.Printf("err !")
	}

	for _, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		t1 := t.UnixNano()

		id, _ := strconv.Atoi(row[1].(string))

		neuronp1.Nst = append(neuronp1.Nst, NeuronST{Time: t1, Id: id})
		if t1 > maxt {
			maxt = t1
		}
		if id > maxid {
			maxid = id
		}
	}
	neuronp1.Sid = maxid
	neuronp1.Stime = maxt
	return neuronp1, err
}

// NeuronP1
type NeuronP5 struct {
	Nst []NeuronST `bson:"nst"`
}

// NeuronST :
type NeuronST5 struct {
	// Uid: Primary key (_id)
	Time int64 `bson:"time"`
	Id   int   `bson:"id"`
}

// 取制定ID范围的神经元，所有周期的脉冲 :
func GetNeuronSpiking3(gid string, max int64) (NeuronP1, error) {

	var neuronp1 NeuronP1
	var maxt int64
	var min int64
	maxt = 0
	maxid := 0
	min = max - 9999

	qs := "select time,id,st  from \"" + gid + "\"  where st=1 and idi > " + strconv.FormatInt(min, 10) + " and idi < " + strconv.FormatInt(max, 10)
	print(qs)
	res, err := db.QueryInfluxDB(db.Conntsdb, "neuron_state_http", qs)
	if err != nil {
		fmt.Printf("err !")
	}

	for _, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		t1 := t.UnixNano()

		id, _ := strconv.Atoi(row[1].(string))

		neuronp1.Nst = append(neuronp1.Nst, NeuronST{Time: t1, Id: id})
		if t1 > maxt {
			maxt = t1
		}
		if id > maxid {
			maxid = id
		}
	}
	neuronp1.Sid = maxid
	neuronp1.Stime = maxt
	return neuronp1, err
}

// 取制定ID范围的神经元，所有周期的脉冲 :
func GetNeuronLimitSpiking(gid string, min int64, max int64) (NeuronP1, error) {

	var neuronp1 NeuronP1
	var maxt int64
	maxt = 0
	maxid := 0

	qs := "select time,id,st  from \"" + gid + "\"  where st=1 and idi > " + strconv.FormatInt(min, 10) + " and idi < " + strconv.FormatInt(max, 10)
	print(qs)
	res, err := db.QueryInfluxDB(db.Conntsdb, "neuron_state_http", qs)
	if err != nil {
		fmt.Printf("err !")
	}

	for _, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		t1 := t.UnixNano()

		id, _ := strconv.Atoi(row[1].(string))

		neuronp1.Nst = append(neuronp1.Nst, NeuronST{Time: t1, Id: id})
		if t1 > maxt {
			maxt = t1
		}
		if id > maxid {
			maxid = id
		}
	}
	neuronp1.Sid = maxid
	neuronp1.Stime = maxt
	return neuronp1, err
}

///////////////////////////////////////////////   读取二进制文件   //////////////////////////////////////

type FileIndex struct {
	StartID  int32
	FileSize int32
	FileName int32
}

// NeuronHM
type NeuronHM struct {
	Sid int        `bson:"sid"`
	Mvl float64    `bson:"mvl"`
	Nvl []NeuronV2 `bson:"nvl"`
}

// 多电压曲线图
type NeuronVls struct {
	Nvls []NeuronVl `bson:"nvls"`
}

// 脉冲点图
type NeuronSTs struct {
	Nst []NeuronST `bson:"nst"`
}

// NeuronST :
type NeuronST2 struct {
	Time int64 `bson:"time"`
	Id   int   `bson:"id"`
}

// NeuronVl
type NeuronVl struct {
	Nvl []float32 `bson:"nvl"`
}

// NeuronV :
type NeuronV2 struct {
	Id int     `bson:"id"`
	Vl float32 `bson:"vl"`
}

// 根据索引文件获取范围内的激活点图
func GetHeatMapNvlFromBin(IndexSumFileName string, NeuSTDir string, t int) (neuronHM NeuronHM, err error) {

	// 获取文件索引
	fileIndexs, _ := GetAllFileIndex(IndexSumFileName, NeuSTDir)
	totalFileNum := len(fileIndexs)
	// 循环每一个文件
	for fIndex := 0; fIndex < totalFileNum; fIndex++ {
		StFileName := NeuSTDir + "/" + strconv.Itoa(int(fileIndexs[fIndex].FileName))
		buf, err := ioutil.ReadFile(StFileName)
		if err != nil {
			fmt.Println(err)
		}
		thisFileNeuNum := fileIndexs[fIndex].FileSize // 该文件中包含了多少个神经元信息
		OneTimeByteNum := thisFileNeuNum * 4          // 该文件一个周期所有神经元所占的字节数
		thisTimeBuf := buf[t*int(thisFileNeuNum) : t*int(thisFileNeuNum)+int(OneTimeByteNum)]
		for neu := 0; neu < int(thisFileNeuNum); neu++ {
			var vdata float32
			vDataBuf := thisTimeBuf[neu*4 : neu*4+4]
			byteBuffer := bytes.NewBuffer(vDataBuf)                  //根据二进制写入二进制结合
			_ = binary.Read(byteBuffer, binary.LittleEndian, &vdata) //解码
			neuronHM.Nvl = append(neuronHM.Nvl, NeuronV2{Id: int(fileIndexs[fIndex].FileName) + neu, Vl: vdata})
		}
	}
	return neuronHM, err
}

// 根据索引文件获取范围内的激活点图
func GetSpkingFromBin(IndexSumFileName string, NeuSTDir string, minId int32, maxId int32) (neuronSTs NeuronSTs, err error) {

	var minIDFileIndex int32
	var maxIDFileIndex int32
	// 获取文件索引
	fileIndexs, time32 := GetAllFileIndex(IndexSumFileName, NeuSTDir)
	time := int(time32)
	for f := 0; f < len(fileIndexs); f++ { // 找到该范围Id在哪几个文件中
		// 最小的
		if fileIndexs[f].StartID <= minId && fileIndexs[f].StartID-1+fileIndexs[f].FileSize >= minId {
			minIDFileIndex = int32(f)
		}
		if fileIndexs[f].StartID <= maxId && fileIndexs[f].StartID-1+fileIndexs[f].FileSize >= maxId {
			maxIDFileIndex = int32(f)
			break
		}
	}

	for f := minIDFileIndex; f <= maxIDFileIndex; f++ {
		StFileName := NeuSTDir + "/" + strconv.Itoa(int(fileIndexs[f].FileName))
		buf, err := ioutil.ReadFile(StFileName)
		if err != nil {
			fmt.Println(err)
		}

		if minIDFileIndex == maxIDFileIndex { // 同一个文件
			for id := minId; id < maxId; id++ { // id
				for t := 0; t < time; t++ {
					if buf[time*int(id-fileIndexs[f].StartID)+t] == 1 {
						neuronSTs.Nst = append(neuronSTs.Nst, NeuronST{Time: int64(t), Id: int(id)})
					}
				}
			}
			break
		}

		if f != minIDFileIndex && f != maxIDFileIndex {
			for id := fileIndexs[f].StartID; id < fileIndexs[f].StartID+fileIndexs[f].FileSize; id++ { // id
				for t := 0; t < time; t++ { //时间
					// 判断是否激活id
					if buf[time*int(id-fileIndexs[f].StartID)+t] == 1 {
						neuronSTs.Nst = append(neuronSTs.Nst, NeuronST{Time: int64(t), Id: int(id)})
					}
				}
			}
		} else if f == minIDFileIndex {
			// 最小的文件夹
			for id := minId; id < fileIndexs[minIDFileIndex].StartID+fileIndexs[minIDFileIndex].FileSize; id++ { // id
				for t := 0; t < time; t++ { //时间
					// 判断是否激活id
					if buf[time*int(id-fileIndexs[minIDFileIndex].StartID)+t] == 1 {
						neuronSTs.Nst = append(neuronSTs.Nst, NeuronST{Time: int64(t), Id: int(id)})
					}
				}
			}
		} else if f == maxIDFileIndex {
			// 最大的文件
			for id := fileIndexs[maxIDFileIndex].StartID; id <= maxId; id++ { // id
				for t := 0; t < time; t++ { //时间
					// 判断是否激活id
					if buf[time*int(id-fileIndexs[maxIDFileIndex].StartID)+t] == 1 {
						neuronSTs.Nst = append(neuronSTs.Nst, NeuronST{Time: int64(t), Id: int(id)})
					}
				}
			}
		}
	}
	return neuronSTs, err
}

// 根据索引文件获取范围内的激活数量
func GetMulNeuVlFromBin(IndexSumFileName string, NeuNlDir string, minId int32, maxId int32) (neuronVls NeuronVls, err error) {

	var minIDFileIndex int32
	var maxIDFileIndex int32
	// 获取文件索引
	fileIndexs, time32 := GetAllFileIndex(IndexSumFileName, NeuNlDir)
	time := int(time32)
	for f := 0; f < len(fileIndexs); f++ { // 找到该范围Id在哪几个文件中
		// 最小的
		if fileIndexs[f].StartID <= minId && fileIndexs[f].StartID-1+fileIndexs[f].FileSize >= minId {
			minIDFileIndex = int32(f)
		}
		if fileIndexs[f].StartID <= maxId && fileIndexs[f].StartID-1+fileIndexs[f].FileSize >= maxId {
			maxIDFileIndex = int32(f)
			break
		}
	}

	for f := minIDFileIndex; f <= maxIDFileIndex; f++ {

		NvlFileName := NeuNlDir + "/" + strconv.Itoa(int(fileIndexs[f].FileName))

		buf, err := ioutil.ReadFile(NvlFileName)
		if err != nil {
			fmt.Println(err)
		}

		if minIDFileIndex == maxIDFileIndex { // 同一个文件
			for id := minId; id < maxId; id++ { // id
				TmpNeuronVl := NeuronVl{}
				for t := 0; t < time; t++ {
					//时间
					var vdata float32
					index := 4*time*int(id-fileIndexs[f].StartID) + t*4
					vDataBuf := buf[index : index+4]
					byteBuffer := bytes.NewBuffer(vDataBuf)                  //根据二进制写入二进制结合
					_ = binary.Read(byteBuffer, binary.LittleEndian, &vdata) //解码
					TmpNeuronVl.Nvl = append(TmpNeuronVl.Nvl, vdata)
				}
				neuronVls.Nvls = append(neuronVls.Nvls, TmpNeuronVl)
			}
			break
		}

		if f != minIDFileIndex && f != maxIDFileIndex {
			// 跨多个文件
			for id := fileIndexs[f].StartID; id < fileIndexs[f].StartID+fileIndexs[f].FileSize; id++ { // id
				TmpNeuronVl := NeuronVl{}
				for t := 0; t < time; t++ {
					//时间
					var vdata float32
					index := 4*time*int(id-fileIndexs[f].StartID) + t*4
					vDataBuf := buf[index : index+4]
					byteBuffer := bytes.NewBuffer(vDataBuf)                  //根据二进制写入二进制结合
					_ = binary.Read(byteBuffer, binary.LittleEndian, &vdata) //解码
					TmpNeuronVl.Nvl = append(TmpNeuronVl.Nvl, vdata)
				}
				neuronVls.Nvls = append(neuronVls.Nvls, TmpNeuronVl)
			}
		} else if f == minIDFileIndex {
			// 最小的文件夹
			for id := minId; id < fileIndexs[minIDFileIndex].StartID+fileIndexs[minIDFileIndex].FileSize; id++ { // id
				TmpNeuronVl := NeuronVl{}
				for t := 0; t < time; t++ {
					//时间
					var vdata float32
					index := 4*time*int(id-fileIndexs[f].StartID) + t*4
					vDataBuf := buf[index : index+4]
					byteBuffer := bytes.NewBuffer(vDataBuf)                  //根据二进制写入二进制结合
					_ = binary.Read(byteBuffer, binary.LittleEndian, &vdata) //解码
					TmpNeuronVl.Nvl = append(TmpNeuronVl.Nvl, vdata)
				}
				neuronVls.Nvls = append(neuronVls.Nvls, TmpNeuronVl)
			}
		} else if f == maxIDFileIndex {
			// 最大的文件
			for id := fileIndexs[maxIDFileIndex].StartID; id <= maxId; id++ { // id
				TmpNeuronVl := NeuronVl{}
				for t := 0; t < time; t++ { //时间
					//时间
					var vdata float32
					index := 4*time*int(id-fileIndexs[maxIDFileIndex].StartID) + t*4
					vDataBuf := buf[index : index+4]
					byteBuffer := bytes.NewBuffer(vDataBuf)                  //根据二进制写入二进制结合
					_ = binary.Read(byteBuffer, binary.LittleEndian, &vdata) //解码
					TmpNeuronVl.Nvl = append(TmpNeuronVl.Nvl, vdata)
				}
				neuronVls.Nvls = append(neuronVls.Nvls, TmpNeuronVl)
			}
		}
	}
	return neuronVls, err
}

// 读取索引文件
func ReadNeuIndexFile(IndexFileName string) (FileIndexs []FileIndex) {
	//将整个文件的内容读到一个字节切片中。
	buf, err := ioutil.ReadFile(IndexFileName)
	if err != nil {
		fmt.Println(err)
	}

	// 获取文件大小
	fi, err := os.Stat(IndexFileName)
	if err != nil {
		fmt.Println(err)
	}

	fileSize := fi.Size()
	fileNum := int(fileSize / 12)

	for f := 0; f < fileNum; f++ {

		index := 12 * f
		var fileIndexs FileIndex
		var UnitData int32
		bytebuffer := bytes.NewBuffer(buf[index : index+4])         //根据二进制写入二进制结合
		_ = binary.Read(bytebuffer, binary.LittleEndian, &UnitData) //解码
		fileIndexs.StartID = UnitData

		bytebuffer = bytes.NewBuffer(buf[index+4 : index+8])        //根据二进制写入二进制结合
		_ = binary.Read(bytebuffer, binary.LittleEndian, &UnitData) //解码
		fileIndexs.FileSize = UnitData

		bytebuffer = bytes.NewBuffer(buf[index+8 : index+12])       //根据二进制写入二进制结合
		_ = binary.Read(bytebuffer, binary.LittleEndian, &UnitData) //解码
		fileIndexs.FileName = UnitData

		FileIndexs = append(FileIndexs, fileIndexs)
	}

	return FileIndexs
}

// 读取总索引文件
func ReadIndexSumFile(IndexSumFileName string) (int32, int32, int32) {

	//将整个文件的内容读到一个字节切片中。
	buf, err := ioutil.ReadFile(IndexSumFileName)
	if err != nil {
		fmt.Println(err)
	}

	var neuNum int32
	var indexFileNum int32
	var time int32

	_ = binary.Read(bytes.NewBuffer(buf[0:4]), binary.LittleEndian, &neuNum)       // 神经元个数
	_ = binary.Read(bytes.NewBuffer(buf[4:8]), binary.LittleEndian, &indexFileNum) // 使用节点数
	_ = binary.Read(bytes.NewBuffer(buf[8:12]), binary.LittleEndian, &time)        // 周期数

	return neuNum, indexFileNum, time
}

func GetAllFileIndex(IndexSumFileName string, NeuDir string) (FileIndexs []FileIndex, time int32) {

	_, indexFileNum, time := ReadIndexSumFile(IndexSumFileName)

	// 获取并拼接所有的index_n中的映射
	for indexF := 0; indexF < int(indexFileNum); indexF++ {
		fileIndexs := ReadNeuIndexFile(NeuDir + "/index_" + strconv.Itoa(indexF))
		FileIndexs = append(FileIndexs, fileIndexs...)
	}

	return FileIndexs, time
}
